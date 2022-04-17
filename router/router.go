package router

import (
	"entry_task/handler"
	"entry_task/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"
)

func Create(sessionStore sessions.Store, h *handler.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	gin.DefaultWriter = ioutil.Discard

	//f, _ := os.Create("./log/gin.log")
	//gin.DefaultWriter = f
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
		//fmt.Println("stack trace from panic ", string(debug.Stack()))
		logrus.Panicf("stack trace from panic ", string(debug.Stack()))
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"request": "hello"})
	})

	v1 := router.Group("/v1")
	{
		u := v1.Group("/user")
		u.GET("/login", middleware.Response("user.login", sessionStore, h.Login))
		u.POST("/create", middleware.Response("user.create", sessionStore, h.CreateAccount))
		u.GET("/info", middleware.UserAccess("user.info", sessionStore, h.GetUserInfo))

		s := v1.Group("/seller")
		s.GET("/shops", middleware.SellerAccess("seller.shop_list", sessionStore, h.SellerShopList))
		s.GET("/shop_info", middleware.SellerAccess("seller.shop_info", sessionStore, h.SellerGetShopInfo))
		s.GET("/shop_products", middleware.SellerAccess("seller.shop_products", sessionStore, h.SellerGetShopProducts))

		s.POST("/create_shop", middleware.SellerAccess("seller.create_shop", sessionStore, h.SellerCreateShop))
		s.POST("/create_product", middleware.SellerAccess("seller.create_product", sessionStore, h.SellerCreateProduct))
		s.POST("/update_product", middleware.SellerAccess("seller.update_product", sessionStore, h.SellerUpdateProduct))

		c := v1.Group("/customer")
		c.GET("/shops", middleware.Response("customer.shop_list", sessionStore, h.CustomerShopList))
		c.GET("/shop_products", middleware.Response("customer.shop", sessionStore, h.CumtomerGetShopProducts))
		c.GET("/feed", middleware.Response("customer.feed", sessionStore, h.CustomerFeed))
		c.GET("/product_search", middleware.Response("customer.search", sessionStore, h.CustomerProductSearch))

		//sell accesss with worker accesss
		p := v1.Group("/product")
		p.POST("/update_status", middleware.Response("product.update_status", sessionStore, h.UpdateProductStatus))
		p.GET("/info", middleware.Response("product.info", sessionStore, h.PorductDetails))
	}
	return router
}
