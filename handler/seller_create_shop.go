package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/id"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) SellerCreateShop(c *gin.Context, base model.UserBase) errno.Payload {
	var req createShopReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("SellerCreateShop req %v", req)

	userName := base.UserName
	if base.UserType != 1 {
		return errno.ERR_NO_PERMISSION
	}
	shopID := id.NewShopID()
	shop := model.SellerShop{
		ShopID:       shopID,
		Name:         req.Name,
		UserName:     userName,
		Introduction: req.Introduction,
	}
	err = h.DB.CreateShop(shop)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			clog.Infof("shop name exist")
			return errno.ERR_SHOP_NAME_EXIST
		}
		clog.Errorf("DB.CreateShop err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	return errno.OK(createShopResp{ShopID: shopID})
}
func checkSellerCreateShopParm(req createShopReq) errno.Payload {
	if len(req.Name) == 0 || len(req.Name) > 20 {
		return errno.ERR_USER_NAME_LEN
	}
	if len(req.Introduction) == 0 || len(req.Introduction) > 20 {
		return errno.ERR_PASSWORD_LEN
	}
	return errno.OK(nil)
	return errno.OK(nil)
}
