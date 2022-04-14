package main

import (
	"entry_task/database"
	"entry_task/handler"
	"entry_task/middleware"
	"entry_task/pkg/feed"
	"entry_task/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

func main() {
	log.Init(logrus.DebugLevel)
	db, err := database.NewMyDB("test")
	if err != nil {
		panic(fmt.Errorf("database.NewMyDB err %s", err.Error()))
	}
	logrus.WithFields(logrus.Fields{"aa":1}).Infof("test")
	err = feed.Init(db)
	if err != nil {
		panic(fmt.Errorf("feed init err %s", err.Error()))
	}

	/*
	err = data_manager.Init(db)
	if err != nil {
		panic(fmt.Errorf("data_manager init err %s", err.Error()))
	}
	
	 */

	//sessionStore := sessions.NewCookieStore([]byte("test"))
	//sessionStore := sessions.NewFilesystemStore("./", securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
	sessionStore := sessions.NewFilesystemStore("./output/session", []byte("session_id"))
	h := handler.Handler{DB: db, SessionStore: sessionStore}

	router := gin.Default()
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.StampMilli),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		fmt.Println("stack trace from panic ", string(debug.Stack()))
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	v1 := router.Group("/v1")
	{
		v1.GET("/user/login", middleware.Response("user.login", sessionStore, h.Login))
		v1.POST("/user/create", middleware.Response("user.create", sessionStore, h.CreateAccount))
		v1.GET("/user/info", middleware.UserAccess("user.info", sessionStore, h.GetUserInfo))

		v1.GET("/seller/shop_list", middleware.SellerAccess("seller.shop_list", sessionStore, h.SellerShopList))
		v1.GET("/seller/shop_info", middleware.SellerAccess("seller.shop_info", sessionStore, h.SellerGetShopInfo))
		v1.GET("/seller/shop_products", middleware.SellerAccess("seller.shop_products", sessionStore, h.SellerGetShopProducts))

		v1.POST("/seller/create_shop", middleware.SellerAccess("seller.create_shop", sessionStore, h.SellerCreateShop))
		v1.POST("/seller/create_product", middleware.SellerAccess("seller.create_product", sessionStore, h.SellerCreateProduct))
		v1.POST("/seller/update_product", middleware.SellerAccess("seller.update_product", sessionStore, h.SellerUpdateProduct))

		v1.GET("/customer/shop_list", middleware.Response("customer.shop_list", sessionStore, h.CustomerShopList))
		v1.GET("/customer/shop_products", middleware.Response("customer.shop", sessionStore, h.CumtomerGetShopProducts))
		v1.GET("/customer/feed", middleware.Response("customer.feed", sessionStore, h.CustomerFeed))
		v1.GET("/customer/product_search", middleware.Response("customer.search", sessionStore, h.CustomerProductSearch))

		//sell accesss with worker accesss
		v1.POST("/product/update_status", middleware.Response("product.update_status", sessionStore, h.UpdateProductStatus))
		v1.GET("/product/info", middleware.Response("product.info", sessionStore, h.PorductDetails))
	}

	fmt.Println("begin to listen")
	router.Run(":8080")
}
