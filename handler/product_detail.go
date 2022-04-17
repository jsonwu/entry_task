package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) PorductDetails(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.GetProductInfoReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("DB.GetUser err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}

	errPaylaod := checkGetProductInfoReq(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	clog.Infof("begin load product in db")
	productInfo, attrs, err := h.DB.GetProductAndAttr(req.ShopID, req.ProductID)
	if err != nil {
		clog.Errorf("DB.GetProductWithAttr err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if productInfo == nil {
		return errno.ERR_PRODUCT_NO_EXIST
	}
	clog.Infof("load product in db success")

	resp := model.ProductInfoReqResp{model.ProductInfo{
		ShopID:     productInfo.ShopID,
		ProductID:  productInfo.ProductID,
		Title:      productInfo.Title,
		CoverUri:   productInfo.CoverUri,
		Price:      productInfo.Price,
		Stock:      productInfo.Stock,
		BrandID:    productInfo.BrandID,
		CategoryID: productInfo.CategoryID,
	}}
	if base.UserType == model.UserTypeSeller {
		resp.Status = int(productInfo.Status)
	}
	resp.ShowUris = []string{}
	if len(productInfo.ShowUris) != 0 {
		resp.ShowUris = strings.Split(productInfo.ShowUris, ",")
	}
	resp.Details = []string{}
	if len(productInfo.Details) != 0 {
		resp.Details = strings.Split(productInfo.Details, ",")
	}
	resp.AttrInfos = make([]model.AttrInfo, 0, len(attrs))
	for _, v := range attrs {
		resp.AttrInfos = append(resp.AttrInfos, model.AttrInfo{v.Name, v.Value})
	}
	clog.Infof("get proudct detail success")
	return errno.OK(resp)
}

func checkGetProductInfoReq(req *model.GetProductInfoReq) model.Payload {
	if !isProductIDValid(req.ProductID) {
		return errno.ERR_PARAM_PRODUCT_ID
	}
	if !isShopIDValid(req.ShopID) {
		return errno.ERR_PARAM_SHOP_ID
	}
	return errno.OK(nil)
}
