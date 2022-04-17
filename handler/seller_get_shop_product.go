package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

const DefaultSellerShopProductsPageSize = 20

func (h *Handler) SellerGetShopProducts(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.SellerGetShopProductsReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("ShouldBind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	errPayload := checkSellerGetShopProductsReq(&req)
	if errPayload.Code != errno.CODE_SUCCESS {
		return errPayload
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = DefaultSellerShopProductsPageSize
		req.PageNum = 1
	}

	clog.Infof("SellerGetShopProducts req %v", req)
	offset := req.PageSize * (req.PageNum - 1)
	clog.Infof("begin load in db offset %d limit %d", offset, req.PageSize)
	productInfos, totalCount, err := h.DB.GetShopProductsWithStatus(req.ShopID, model.PStatus(req.Status), offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopProductsWithStatus err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	logs.Infof("load product in db success total num: %d shop num: %d", totalCount, len(productInfos))

	resp := model.SellerGetShopProductsResp{TotalCount: totalCount, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]model.ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, model.ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}
	clog.Infof("seller get shop product success")
	return errno.OK(resp)
}
func checkSellerGetShopProductsReq(req *model.SellerGetShopProductsReq) model.Payload {
	return errno.OK(nil)
}
