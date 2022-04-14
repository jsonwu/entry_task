package middleware

import (
	"entry_task/errno"
	"entry_task/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

type MyHandler func(c *gin.Context, base model.UserBase) errno.Payload


func UserAccess(mkey string, store sessions.Store, f MyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session_id")
		if err != nil {
			c.JSON(http.StatusOK, errno.ERR_INTERNAL)
			return
		}
		if (session.IsNew){
			c.JSON(http.StatusOK, errno.ERR_NO_LOGIN)
			return
		}
		userInfo := getUserBase(session)
		data := f(c, userInfo)
		c.JSON(http.StatusOK, data)
	}
}


func SellerAccess(mkey string, store sessions.Store, f MyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session_id")
		if err != nil {
			fmt.Println("1", err)
			c.JSON(http.StatusOK, errno.ERR_INTERNAL)
			return
		}
		if (session.IsNew){
			c.JSON(http.StatusOK, errno.ERR_NO_LOGIN)
			return
		}
		userInfo := getUserBase(session)
		if userInfo.UserType != model.UserTypeSeller{
			c.JSON(http.StatusOK, errno.ERR_MUST_USER_TYPE_SELLER)
			return
		}
		data := f(c, userInfo)
		c.JSON(http.StatusOK, data)
	}
}

func Response(mkey string, store sessions.Store, f MyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session_id")
		if err != nil {
			fmt.Println("1", err)
			c.JSON(http.StatusOK, errno.ERR_INTERNAL)
			return
		}
		userInfo := getUserBase(session)
		fmt.Println("session info", session.Values)
		data := f(c, userInfo)
		c.JSON(http.StatusOK, data)
	}
}

func getUserBase(session *sessions.Session) model.UserBase{
	userInfo := model.UserBase{}
	if v, ok := session.Values["user_name"]; ok {
		userInfo.UserName = v.(string)
	}
	if v, ok := session.Values["user_type"]; ok {
		userInfo.UserType = model.UserType(v.(uint8))
	}
	userInfo.DeviceID = session.ID //temp
	return userInfo
}
