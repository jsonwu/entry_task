package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) CumtomerGetShopProducts(c *gin.Context, base model.UserBase) errno.Payload {
	var req customerGetShopProductsReq
	err := c.ShouldBind(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}

	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})

	clog.Infof("CumtomerGetShopProducts req %v", req)

	errPaylaod := checkCustomerGetShopProductsParam(req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = 20
		req.PageNum = 1
	}

	offset := req.PageSize * (req.PageNum - 1)
	clog.Infof("db serarch offset %d limit %d", offset, req.PageSize)
	productInfos, totalCount, err := h.DB.GetShopProducts(req.ShopID, offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopProducts err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Debugf("db get ", totalCount, productInfos)
	resp := customerGetShopProductsResp{TotalCount: totalCount, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}
	return errno.OK(resp)
}
func checkCustomerGetShopProductsParam(req customerGetShopProductsReq) errno.Payload {
	return errno.OK(nil)
}
