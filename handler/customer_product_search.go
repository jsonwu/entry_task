package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) CustomerProductSearch(c *gin.Context, base model.UserBase) errno.Payload {
	var req customerProductSearchReq
	err := c.ShouldBind(&req)
	if err != nil {
		logs.Errorf("c.BindJson err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("CustomerProductSearch %v", req)

	errPaylaod := checkProductSearchReqParam(req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageNum == 0 || req.PageSize == 0 {
		req.PageNum = 1
		req.PageSize = 20
	}

	offset := (req.PageNum - 1) * req.PageSize
	productInfos, total, err := h.DB.SearchProductInTitle(req.KeyWord, offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.SearchProductInTitle err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	resp := customerSearchResp{TotalCount: total, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}

	return errno.OK(resp)
}
func checkProductSearchReqParam(req customerProductSearchReq) errno.Payload {
	//todo
	return errno.OK(nil)
}
