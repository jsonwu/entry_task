package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/feed"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func (h *Handler) CustomerFeed(c *gin.Context, base model.UserBase) errno.Payload {
	var req feedReq
	err := c.ShouldBind(&req)
	if err != nil {
		return errno.ERR_INVALID_PARAM
	}
	errPaylod := checkfeedReqParam(req)
	if err != nil {
		return errPaylod
	}
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	clog.Infof("CustomerFeed %v", req)

	if req.Num == 0 {
		req.Num = 10
	}
	product, err := feed.GetFeed(base.UserName, req.Num)
	if err != nil {
		clog.Info("feed.GetFeed err %s", err.Error())
	}
	resp := feedResp{ProductBaseInfos: make([]ProductBaseInfo, 0, len(product))}
	for _, v := range product {
		p := ProductBaseInfo{
			ShopID:    v.ShopID,
			ProductID: v.ProductID,
			CoverUri:  v.CoverUri,
			Price:     v.Price,
		}
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, p)
	}
	return errno.OK(resp)
}

func checkfeedReqParam(req feedReq) errno.Payload {
	if req.Num < 0 || req.Num > 50 {
		return errno.ERR_INVALID_PARAM
	}
	return errno.OK(nil)
}
