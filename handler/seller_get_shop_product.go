package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) SellerGetShopProducts(c *gin.Context, base model.UserBase) errno.Payload {
	var req SellerGetShopProductsReq
	err := c.ShouldBind(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	errPayload := checkSellerGetShopProductsReq(req)
	if errPayload.Code != errno.CODE_SUCCESS {
		return errPayload
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = 20
		req.PageNum = 1
	}

	clog.Infof("SellerGetShopProducts req %v", req)
	offset := req.PageSize * (req.PageNum - 1)
	clog.Infof("offset %d limit %d", offset, req.PageSize)
	productInfos, totalCount, err := h.DB.GetShopProductsWithStatus(req.ShopID, model.PStatus(req.Status), offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopProductsWithStatus err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	logs.Infof("product num %d", len(productInfos), productInfos, err)
	resp := sellerGetShopProductsResp{TotalCount: totalCount, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}
	return errno.OK(resp)
}
func checkSellerGetShopProductsReq(req SellerGetShopProductsReq) errno.Payload {
	return errno.OK(nil)
}
