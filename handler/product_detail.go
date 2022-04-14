package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) PorductDetails(c *gin.Context, base model.UserBase) errno.Payload {
	var req getProductInfoReq
	err := c.ShouldBind(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}

	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	errPaylaod := checkGetProductInfoReq(req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	productInfo, attrs, err := h.DB.GetProductWithAttr(req.ShopID, req.ProductID)
	if err != nil {
		clog.Errorf("DB.GetProductWithAttr err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if productInfo == nil {
		return errno.ERR_PRODUCT_NO_EXIST
	}
	resp := productInfoReqResp{ProductInfo{
		ShopID:     productInfo.ShopID,
		ProductID:  productInfo.ProductID,
		Title:      productInfo.Title,
		CoverUri:   productInfo.CoverUri,
		Price:      productInfo.Price,
		Stock:      productInfo.Stock,
		BrandID:    productInfo.BrandID,
		CategoryID: productInfo.CategoryID,
	}}
	resp.ShowUris = []string{}
	if len(productInfo.ShowUris) != 0 {
		resp.ShowUris = strings.Split(productInfo.ShowUris, ",")
	}
	resp.Details = []string{}
	if len(productInfo.Details) != 0 {
		resp.Details = strings.Split(productInfo.Details, ",")
	}
	resp.AttrInfo = make([]AttrInfo, 0, len(attrs))
	for _, v := range attrs {
		resp.AttrInfo = append(resp.AttrInfo, AttrInfo{v.Name, v.Value})
	}
	return errno.OK(resp)
}

func checkGetProductInfoReq(req getProductInfoReq) errno.Payload {
	return errno.OK(nil)
}
