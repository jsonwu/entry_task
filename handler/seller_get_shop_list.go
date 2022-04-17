package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) SellerShopList(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("SellerShopList req")
	userName := base.UserName

	clog.Infof("begin load seller shop in db")
	shopInfos, err := h.DB.GetSellerShop(userName)
	if err != nil {
		clog.Errorf("DB.GetSellerShop err", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("load seller shop nums %d", len(shopInfos))
	resp := model.SellerShopListResp{}
	resp.ShopBaseInfos = make([]model.ShopBaseInfo, 0, len(shopInfos))
	for _, v := range shopInfos {
		s := model.ShopBaseInfo{
			ShopID:       v.ShopID,
			Name:         v.Name,
			Introduction: v.Introduction,
		}
		resp.ShopBaseInfos = append(resp.ShopBaseInfos, s)
	}
	clog.Infof("seller get shoplist success")
	return errno.OK(resp)
}
