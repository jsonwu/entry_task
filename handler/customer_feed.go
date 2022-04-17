package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/feed"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	_ "net/http/pprof"
)

//logrus 为同步写日志，性能较低，需要修改为异步日志库，日志暂时时debug
const DefaultFeedNum = 10

func (h *Handler) CustomerFeed(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.FeedReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("ShouldBind err %s", err.Error())
		return errno.ERR_INVALID_PARAM
	}
	errPaylod := checkfeedReqParam(&req)
	if err != nil {
		return errPaylod
	}
	clog.Debugf("CustomerFeed %+v", req)

	if req.Num == 0 {
		req.Num = DefaultFeedNum
	}
	clog.Debugf("begin get product from feed system")
	product, err := feed.GetFeed(base.UserName, req.Num)
	if err != nil {
		clog.Info("feed.GetFeed err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	//clog.Infof("get product from feed product num: %d", len(product))
	resp := model.FeedResp{ProductBaseInfos: make([]model.ProductBaseInfo, 0, len(product))}
	for _, v := range product {
		p := model.ProductBaseInfo{
			ShopID:    v.ShopID,
			ProductID: v.ProductID,
			CoverUri:  v.CoverUri,
			Price:     v.Price,
			Title:     v.Title,
		}
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, p)
	}
	clog.Debugf("feed success")
	return errno.OK(resp)
}

const FeedNumMax = 50
const FeedNumMin = 1

func checkfeedReqParam(req *model.FeedReq) model.Payload {
	if req.Num < FeedNumMin || req.Num > FeedNumMax {
		return errno.ERR_INVALID_PARAM
	}
	return errno.OK(nil)
}
