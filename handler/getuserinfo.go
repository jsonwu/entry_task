package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) GetUserInfo(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("begin get user in db")
	userInfo, err := h.DB.GetUser(base.UserName, base.UserType)
	if err != nil {
		clog.Errorf("DB.GetUser err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if userInfo == nil {
		clog.Errorf("user login but get userinfo  empty from db check")
		return errno.ERR_INTERNAL
	}
	resp := model.GetUserInfoResp{UserName: base.UserName, UserType: base.UserType, Email: userInfo.Email, ProfileUri: userInfo.ProfileUri}
	clog.Infof("get user info success")
	return errno.OK(resp)
}
