package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/id"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) SellerCreateProduct(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.CreateProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logs.Errorf("ShouldBindJson err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("SellerCreateProduct req %+v", req)

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
		Status:     model.ProudctStatusAuditing,
		Details:    strings.Join(req.Details, ","),
	}
	productAttrs := make([]*model.ProductAttr, 0, len(req.AttrInfos))
	for _, v := range req.AttrInfos {
		productAttrs = append(productAttrs, &model.ProductAttr{ProductID: productID, Name: v.Name, Value: v.Value})
	}

	errPaylod := checkCreatProductParam(&req, &dbProduct)
	if errPaylod.Code != 0 {
		return errPaylod
	}
	shop, err := h.DB.GetShopInfo(req.ShopID)
	if err != nil {
		clog.Errorf("DB.GetShopInfo %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if shop == nil {
		return errno.ERR_SHOP_NOT_EXIST
	}
	if shop.UserName != base.UserName {
		clog.Errorf("create shop shop owner: %s  user: %s", shop.UserName, base.UserName)
		return errno.ERR_NOT_SHOP_OWNER
	}
	clog.Infof("beiin create product in db")
	err = h.DB.CreateProductWithAttr(&dbProduct, productAttrs)
	if err != nil {
		clog.Errorf("DB.CreateProduct err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("create product success")
	return errno.OK(model.CreateProductResp{ProductID: productID})
}

func checkCreatProductParam(req *model.CreateProductReq, product *model.Product) model.Payload {
	if !(isAttrValid(req.AttrInfos)) {
		return errno.ERR_ATTR_NAME
	}
	return checkProduct(product)
}
