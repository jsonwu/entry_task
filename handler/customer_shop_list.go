package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) CustomerShopList(c *gin.Context, base model.UserBase) errno.Payload {
	var req customerShopListReq
	err := c.ShouldBind(&req)
	if err != nil {
		logs.Errorf("c.BindJson err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("CustomerShopList %v", req)
	errPaylaod := checkCustomerShopListReqParm(req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageNum == 0 || req.PageSize == 0 {
		req.PageNum = 1
		req.PageSize = 20
	}

	offset := (req.PageNum - 1) * req.PageSize
	clog.Infof("serarch db offset %d limit %d", offset, req.PageSize)
	shopInfos, totalNum, err := h.DB.GetShopList(offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopList err", err.Error())
		return errno.ERR_INTERNAL
	}
	resp := customerShopListResp{PageNum: req.PageNum, PageSize: req.PageSize, TotalCount: totalNum}
	resp.ShopBaseInfos = make([]ShopBaseInfo, 0, len(shopInfos))
	for _, v := range shopInfos {
		s := ShopBaseInfo{
			ShopID:       v.ShopID,
			Name:         v.Name,
			Introduction: v.Introduction,
		}
		resp.ShopBaseInfos = append(resp.ShopBaseInfos, s)
	}
	return errno.OK(resp)
}

func checkCustomerShopListReqParm(req customerShopListReq) errno.Payload {
	return errno.OK(nil)
}
