package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/id"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

const DefaultShopLevel = 1

func (h *Handler) SellerCreateShop(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.CreateShopReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logs.Errorf("ShouldBindJson err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("SellerCreateShop req %+v", req)
	errPaylaod := checkSellerCreateShopParm(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	userName := base.UserName
	shopID := id.NewShopID()
	shop := model.SellerShop{
		ShopID:       shopID,
		Name:         req.Name,
		UserName:     userName,
		Level:        DefaultShopLevel,
		Introduction: req.Introduction,
	}

	clog.Infof("begin  create shop in db")
	err = h.DB.CreateShop(&shop)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			clog.Infof("shop name exist")
			return errno.ERR_SHOP_NAME_EXIST
		}
		clog.Errorf("DB.CreateShop err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("seller create shop success")
	return errno.OK(model.CreateShopResp{ShopID: shopID})
}
func checkSellerCreateShopParm(req *model.CreateShopReq) model.Payload {
	if !isShopNameValid(req.Name) {
		return errno.ERR_PARAM_SHOP_NAME
	}
	//Introduction
	return errno.OK(nil)
}
