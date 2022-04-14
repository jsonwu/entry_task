package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) UpdateProductStatus(c *gin.Context, base model.UserBase) errno.Payload {
	var req UpdateProductStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})

	clog.Infof("UpdateProductStatus req %+v", req)
	//check stauts
	err, effected := h.DB.UpdateProductStatus(req.ShopID, req.ProductID, model.PStatus(req.Status))
	if err != nil {
		return errno.ERR_INTERNAL
	}
	logs.Infof("effected %v", effected)
	return errno.OK(nil)

}
