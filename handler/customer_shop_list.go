package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

const DefaultCustomerShopPageSize = 20

func (h *Handler) CustomerShopList(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.CustomerShopListReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("c.BindJson err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("CustomerShopList %+v", req)
	errPaylaod := checkCustomerShopListReqParm(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageNum == 0 || req.PageSize == 0 {
		req.PageNum = 1
		req.PageSize = DefaultCustomerShopPageSize
	}

	offset := (req.PageNum - 1) * req.PageSize
	clog.Infof("begin get shop in db offset %d limit %d", offset, req.PageSize)
	shopInfos, totalNum, err := h.DB.GetShopList(offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.GetShopList err", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("begin get shop in db success total num: %d shop num: %d", totalNum, len(shopInfos))

	resp := model.CustomerShopListResp{PageNum: req.PageNum, PageSize: req.PageSize, TotalCount: totalNum}
	resp.ShopBaseInfos = make([]model.ShopBaseInfo, 0, len(shopInfos))
	for _, v := range shopInfos {
		s := model.ShopBaseInfo{
			ShopID:       v.ShopID,
			Name:         v.Name,
			Introduction: v.Introduction,
		}
		resp.ShopBaseInfos = append(resp.ShopBaseInfos, s)
	}
	clog.Infof("customer shop list success")
	return errno.OK(resp)
}

const ConstomerShopListPageSizeMax = 100

func checkCustomerShopListReqParm(req *model.CustomerShopListReq) model.Payload {
	if req.PageSize < 0 || req.PageSize > ConstomerShopListPageSizeMax {
		return errno.ERR_PAGESIZE
	}
	return errno.OK(nil)
}
