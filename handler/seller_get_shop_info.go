package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) SellerGetShopInfo(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.SellerGetShopInfoReq
	err := c.ShouldBind(&req)
	if err != nil {
		logs.Errorf("ShouldBind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("SellerGetShopInfo req %+v", req)

	errPaylaod := checksellerGetShopInfoReqParam(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	userName := base.UserName
	clog.Infof("begin load shop in db")
	shopInfo, err := h.DB.GetShopInfo(req.ShopID)
	if err != nil {
		clog.Errorf("DB.GetShopInfo err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if shopInfo == nil {
		clog.Infof("shop not exist")
		return errno.ERR_SHOP_NOT_EXIST
	}
	clog.Infof("load shop in db succes")
	//or search db with user name
	if shopInfo.UserName != userName {
		return errno.ERR_SHOP_ID_ILLEGAL
	}
	resp := model.SellerGetShopInfoResp{model.ShopInfo{ShopID: shopInfo.ShopID, Name: shopInfo.Name,
		Level: shopInfo.Level, Introduction: shopInfo.Introduction, Location: shopInfo.Location}}

	clog.Infof("seller get shop succes")
	return errno.OK(resp)
}

func checksellerGetShopInfoReqParam(req *model.SellerGetShopInfoReq) model.Payload {
	return errno.OK(nil)
}
