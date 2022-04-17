package handler

import (
	"entry_task/database"
	"entry_task/pkg/product_center"
	"entry_task/pkg/shop_center"
	"entry_task/pkg/user_center"
	"github.com/gorilla/sessions"
)

type Handler struct {
	//should not use db
	DB            *database.MyDB
	SessionStore  sessions.Store
	UserCeneter   *user_center.UserCenter
	ShopCenter    *shop_center.ShopCenter
	ProductCenter *product_center.ProductCenter
}
