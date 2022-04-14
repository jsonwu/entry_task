package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) SellerGetShopInfo(c *gin.Context, base model.UserBase) errno.Payload {
	var req sellerGetShopInfoReq
	err := c.ShouldBind(&req)
	if err != nil {
		logs.Errorf("gin bind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}

	errPaylaod := checksellerGetShopInfoReqParam(req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	userName := base.UserName
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})

	shopID := c.Param("shop_id")
	shopInfo, err := h.DB.GetShopInfo(shopID)
	if err != nil {
		clog.Errorf("DB.GetShopInfo err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if shopInfo == nil {
		return errno.ERR_SHOP_NOT_EXIST
	}
	//or search db with user name
	if shopInfo.UserName != userName {
		return errno.ERR_SHOP_ID_ILLEGAL
	}
	resp := sellerGetShopInfoResp{ShopInfo{ShopID: shopInfo.ShopID, Name: shopInfo.Name,
		Level: shopInfo.Level, Introduction: shopInfo.Introduction, location: shopInfo.Location}}
	return errno.OK(resp)
}

func checksellerGetShopInfoReqParam(req sellerGetShopInfoReq) errno.Payload {
	return errno.OK(nil)
}
