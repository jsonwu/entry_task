package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) SellerUpdateProduct(c *gin.Context, base model.UserBase) errno.Payload {
	var req updateProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("SellerUpdateProduct req %v", req)
	
	dbProduct := model.Product{
		ShopID:     req.ShopID,
		ProductID:  req.ProductID,
		Title:      req.Title,
		CoverUri:   req.CoverUri,
		ShowUris:   strings.Join(req.ShowUris, ","),
		Price:      req.Price,
		Stock:      req.Stock,
		BrandID:    req.BrandID,
		CategoryID: req.CategoryID,
		Details:    strings.Join(req.Details, ","),
	}

	productAttrs := make([]model.ProductAttr, 0, len(req.AttrInfo))
	for _, v := range req.AttrInfo {
		productAttrs = append(productAttrs, model.ProductAttr{ProductID: req.ProductID, Name: v.Name, Value: v.Value})
	}

	//to check shop owner
	shopInfo, err := h.DB.GetShopInfo(req.ShopID)
	if err != nil {
		clog.Errorf("get shop info err %s", err.Error())
		return errno.ERR_INTERNAL
	}

	if shopInfo == nil {
		return errno.ERR_SHOP_NOT_EXIST
	}

	if shopInfo.UserName != base.UserName {
		clog.Errorf("not shop owner")
		return errno.ERR_NO_PERMISSION
	}
	fmt.Println("to save", productAttrs)
	err, effectRows := h.DB.UpdateProductWithAttr(dbProduct, productAttrs)
	if err != nil {
		clog.Errorf("DB.UpdateProductWithAttr err %s", err.Error(), effectRows)
		return errno.ERR_INTERNAL
	}
	if effectRows == 0 {
		clog.Infof("effect db num %d", effectRows)
		return errno.ERR_PRODUCT_NO_EXIST
	}
	return errno.OK(nil)
}
