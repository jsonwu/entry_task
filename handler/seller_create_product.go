package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/id"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) SellerCreateProduct(c *gin.Context, base model.UserBase) errno.Payload {
	var req createProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	errPaylod := checkCreatProductParam(req)
	if errPaylod.Code != 0 {
		return errPaylod
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})

	productID := id.NewProductID()
	dbProduct := model.Product{
		ShopID:     req.ShopID,
		ProductID:  productID,
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

	errPaylod = checkProduct(dbProduct)
	if errPaylod.Code != 0 {
		return errPaylod
	}
	// checker shop owner
	shop, err := h.DB.GetShopInfo(req.ShopID)
	if err != nil {
		clog.Errorf("DB.GetShopInfo %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if shop == nil {
		return errno.ERR_SHOP_NOT_EXIST
	}
	if shop.UserName != base.UserName {
		clog.Errorf("create shop user name err")
		return errno.ERR_NO_PERMISSION
	}
	fmt.Println(productAttrs)

	err = h.DB.CreateProductWithAttr(dbProduct, productAttrs)
	if err != nil {
		clog.Errorf("DB.CreateProduct err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("create shop success")
	return errno.OK(createProductResp{ProductID: productID})
}

func checkCreatProductParam(product createProductReq) errno.Payload {
	return errno.OK(nil)
}

func checkProduct(product model.Product) errno.Payload {
	if len(product.Title) < 1 || len(product.Title) > 20 {
		return errno.ERR_PRODUCT_TITLE_LEN
	}
	//todo
	return errno.OK(nil)
}
