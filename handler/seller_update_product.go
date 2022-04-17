package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) SellerUpdateProduct(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.UpdateProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		clog.Errorf("ShouldBindJson err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("SellerUpdateProduct req %+v", req)

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
	productAttrs := make([]*model.ProductAttr, 0, len(req.AttrInfos))
	for _, v := range req.AttrInfos {
		productAttrs = append(productAttrs, &model.ProductAttr{ProductID: req.ProductID, Name: v.Name, Value: v.Value})
	}

	errPaylod := checkSellerUpdateProductReq(&req, &dbProduct)
	if errPaylod.Code != errno.CODE_SUCCESS {
		return errPaylod
	}
	shopInfo, err := h.DB.GetShopInfo(req.ShopID)
	if err != nil {
		clog.Errorf("get shop info err %s", err.Error())
		return errno.ERR_INTERNAL
	}

	if shopInfo == nil {
		return errno.ERR_SHOP_NOT_EXIST
	}

	if shopInfo.UserName != base.UserName {
		clog.Errorf("update product shop owner %s user %s", shopInfo.UserName, base.UserName)
		return errno.ERR_NOT_SHOP_OWNER
	}

	//gorm rowaffected  have bug
	p, err := h.DB.GetProductInfo(dbProduct.ShopID, dbProduct.ProductID)
	if err != nil {
		clog.Errorf("DB.GetProductInfo err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if p == nil {
		clog.Infof("DB.GetProduct nil")
		return errno.ERR_PRODUCT_NO_EXIST
	}

	clog.Infof("begin update product in db")
	err, effectRows := h.DB.UpdateProductWithAttr(&dbProduct, productAttrs)
	if err != nil {
		clog.Errorf("DB.UpdateProductWithAttr err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	/*
		if effectRows == 0 {
			clog.Errorf("effect db num %d %+v", effectRows, dbProduct)
			return errno.ERR_PRODUCT_NO_EXIST
		}
	*/
	clog.Infof("effect %d", effectRows)
	clog.Infof("update product success")
	return errno.OK(nil)
}
func checkSellerUpdateProductReq(req *model.UpdateProductReq, product *model.Product) model.Payload {
	if !(isAttrValid(req.AttrInfos)) {
		return errno.ERR_ATTR_NAME
	}
	return checkProduct(product)
}
