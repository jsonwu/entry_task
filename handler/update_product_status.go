package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) UpdateProductStatus(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.UpdateProductStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		clog.Errorf("ShouldBindJSON err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	clog.Infof("UpdateProductStatus req %+v", req)

	errPaylod := checkUpdateProductStatusReqParam(&req)
	if errPaylod.Code != errno.CODE_SUCCESS {
		return errPaylod
	}

	//gorm  rowaffected  have bug so use select
	p, err := h.DB.GetProductInfo(req.ShopID, req.ProductID)
	if err != nil {
		clog.Errorf("DB.GetProductInfo err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	if p == nil {
		clog.Infof("DB.GetProduct nil")
		return errno.ERR_PRODUCT_NO_EXIST
	}

	clog.Infof("begin update product status in db")
	err, effected := h.DB.UpdateProductStatus(req.ShopID, req.ProductID, model.PStatus(req.Status))
	if err != nil {
		return errno.ERR_INTERNAL
	}
	clog.Infof("effected %v", effected)
	/*
		if effected == 0 {
			return errno.ERR_PRODUCT_NO_EXIST
		}
	*/
	clog.Infof("update product status success")
	return errno.OK(nil)
}

func checkUpdateProductStatusReqParam(req *model.UpdateProductStatusReq) model.Payload {
	if !isProductStatusValid(req.Status) {
		return errno.ERR_PRODUCT_STATUS
	}
	return errno.OK(nil)
}
