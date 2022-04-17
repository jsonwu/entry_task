package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

const DefaultCustomerShopProductsPageSize = 20

func (h *Handler) CumtomerGetShopProducts(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.CustomerGetShopProductsReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("ShouldBind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("CumtomerGetShopProducts req %+v", req)

	errPaylaod := checkCustomerGetShopProductsParam(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = DefaultCustomerShopProductsPageSize
		req.PageNum = 1
	}

	offset := req.PageSize * (req.PageNum - 1)
	clog.Infof("begin get shop products in db offset %d limit %d", offset, req.PageSize)
	productInfos, totalCount, err := h.DB.GetShopProductsWithStatus(req.ShopID, model.ProudctStatusNormal, offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopProducts err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Debugf("get shop product in db success totalcount: %d  product num: %d ", totalCount, len(productInfos))
	resp := model.CustomerGetShopProductsResp{TotalCount: totalCount, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]model.ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, model.ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}
	clog.Infof("get shop products success")
	return errno.OK(resp)
}

const CustomerShopListPageSizeMin = 0
const CustomerShopListPageSizeMax = 100

func checkCustomerGetShopProductsParam(req *model.CustomerGetShopProductsReq) model.Payload {
	if !isShopIDValid(req.ShopID) {
		return errno.ERR_PARAM_SHOP_ID
	}
	if req.PageSize < CustomerShopListPageSizeMin || req.PageSize > CustomerShopListPageSizeMax {
		return errno.ERR_PAGESIZE
	}
	return errno.OK(nil)
}
