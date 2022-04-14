package handler

import (
	"crypto/md5"
	"entry_task/errno"
	"entry_task/model"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"strings"
	"time"
)

var l = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(l)
	for i := range b {
		b[i] = l[rand.Intn(length)]
	}
	return string(b)
}

func hashSalt(str, salt string) string {
	m := md5.New()
	io.WriteString(m, str)
	io.WriteString(m, salt)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func (h *Handler) CreateAccount(c *gin.Context, base model.UserBase) errno.Payload {
	var req createAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}

	clog := logs.WithFields(logs.Fields{"user_name": req.UserName, "user_type": req.UserType})
	errPaylod := checkCreateAccountReqParam(req)
	if errPaylod.Code != errno.CODE_SUCCESS {
		return errPaylod
	}

	clog.Infof("CreateAccount req %v", req)
	salt := randomStr(10)
	encodePassWord := hashSalt(req.Password, salt)
	err = h.DB.CreateUser(model.User{UserName: req.UserName, Password: encodePassWord, Salt: salt, UserType: req.UserType, Email: req.Email, ProfileUri: req.Email})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			clog.Infof("user name exist")
			return errno.ERR_USERNAME_EXIST
		}
		clog.Errorf("DB.CreateUser err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("create account success")
	return errno.OK(nil)
}

func checkCreateAccountReqParam(req createAccountReq) errno.Payload {
	return errno.OK(nil)
}
