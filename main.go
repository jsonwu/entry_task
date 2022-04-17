package main

import (
	"entry_task/config"
	"entry_task/database"
	"entry_task/handler"
	"entry_task/pkg/data_manager"
	"entry_task/pkg/feed"
	"entry_task/pkg/log"
	"entry_task/pkg/product_center"
	"entry_task/pkg/shop_center"
	"entry_task/pkg/user_center"
	"entry_task/router"
	"fmt"
	"github.com/gorilla/sessions"
	logs "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Init(logs.InfoLevel)
	config, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("load config err %s", err.Error()))
	}

	db, err := database.NewMyDB(&config.MasterDB)
	if err != nil {
		//logrus 内部会直接调用 panic
		logs.Panicf("database.NewMyDB err %s", err.Error())
		panic(fmt.Errorf("database.NewMyDB err %s", err.Error()))
	}

	userCenter := user_center.NewUserCenter(db)
	shopCenter := shop_center.NewShopCenter(db)
	productCenter := product_center.NewProductCenter(db)
	err = feed.Init(db)
	if err != nil {
		logs.Panicf("feed init err %s", err.Error())
		panic(fmt.Errorf("feed init err %s", err.Error()))
	}

	err = data_manager.Init(db)
	if err != nil {
		logs.Panicf("data_manager init err %s", err.Error())
		panic(fmt.Errorf("data_manager init err %s", err.Error()))
	}
	rand.Seed(time.Now().Unix())

	sessionStore := sessions.NewFilesystemStore("./output/session", []byte("session_id"))
	h := handler.Handler{
		DB:            db,
		SessionStore:  sessionStore,
		UserCeneter:   userCenter,
		ShopCenter:    shopCenter,
		ProductCenter: productCenter,
	}
	//pprof
	go http.ListenAndServe(":6060", nil)
	
	router := router.Create(sessionStore, &h)
	logs.Infof("Listen")
	err = router.Run(":8080")
	if err != nil {
		logs.Panicf("router run err %s", err.Error())
		panic(fmt.Errorf("router run err %s", err.Error()))
	}
}
