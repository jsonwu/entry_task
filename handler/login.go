package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) Login(c *gin.Context, base model.UserBase) errno.Payload {
	var req loginReq
	err := c.ShouldBind(&req)
	if err != nil {
		logs.Errorf("gin bind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog := logs.WithFields(logs.Fields{"user_name": req.UserName, "user_type": req.UserType})
	clog.Infof("Login request: %v", req)

	errPaylod := checkLoginParam(req)
	if errPaylod.Code != 0 {
		clog.Debugf("param err %v", errPaylod.Msg)
		return errPaylod
	}

	user, err := h.DB.GetUser(req.UserName, req.UserType)
	if err != nil {
		clog.Errorf("get user info from db eff %s", err.Error())
		return errno.ERR_INTERNAL
	}

	if user == nil {
		return errno.ERR_USER_NOT_EXIST
	}

	hs := hashSalt(req.Password, user.Salt)
	if hs != user.Password {
		clog.Infof("user input password err")
		return errno.ERR_PASSWORD
	}
	session, err := h.SessionStore.Get(c.Request, "session_id")
	if err != nil {
		clog.Errorf("sesssion get err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	session.Values["user_name"] = req.UserName
	session.Values["user_type"] = uint8(req.UserType)
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		clog.Infof("session save error %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("login success")
	return errno.OK(nil)
}

func checkLoginParam(req loginReq) errno.Payload {
	if len(req.UserName) < 1 || len(req.UserName) > 20 {
		return errno.ERR_USER_NAME_LEN
	}
	if len(req.Password) < 1 || len(req.Password) > 20 {
		return errno.ERR_PASSWORD_LEN
	}

	if !(req.UserType == model.UserTypeSeller || req.UserType == model.UserTypeCustomer) {
		return errno.ERR_USER_TYPE
	}
	return errno.OK(nil)
}
