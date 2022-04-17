package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

const DefaultSearchProductPageSize = 20

func (h *Handler) CustomerProductSearch(c *gin.Context, base model.UserBase) model.Payload {
	clog := logs.WithFields(logs.Fields{"user_name": base.UserName, "user_type": base.UserType})
	var req model.CustomerProductSearchReq
	err := c.ShouldBind(&req)
	if err != nil {
		clog.Errorf("ShouldBind err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("CustomerProductSearch %+v", req)

	errPaylaod := checkProductSearchReqParam(&req)
	if errPaylaod.Code != errno.CODE_SUCCESS {
		return errPaylaod
	}

	if req.PageNum == 0 || req.PageSize == 0 {
		req.PageNum = 1
		req.PageSize = DefaultSearchProductPageSize
	}

	offset := (req.PageNum - 1) * req.PageSize
	clog.Infof("begin search product in db offset %d limit %d", offset, req.PageSize)
	productInfos, total, err := h.DB.SearchProductInTitle(req.KeyWord, offset, req.PageSize)
	if err != nil {
		clog.Errorf("DB.SearchProductInTitle err %s", err.Error())
		return errno.ERR_INTERNAL
	}
	clog.Infof("search product in db success  shop  total num:%d  product num: %d", total, len(productInfos))

	resp := model.CustomerSearchResp{TotalCount: total, PageNum: req.PageNum, PageSize: req.PageSize}
	resp.ProductBaseInfos = make([]model.ProductBaseInfo, 0, len(productInfos))
	for _, v := range productInfos {
		resp.ProductBaseInfos = append(resp.ProductBaseInfos, model.ProductBaseInfo{ProductID: v.ProductID, ShopID: v.ShopID, Price: v.Price, Title: v.Title, CoverUri: v.CoverUri})
	}
	clog.Infof("search products success")
	return errno.OK(resp)
}

const SearchKeyWordLenMax = 20
const SearchKeyWordLenMin = 1

const SearchPageSizeMax = 100

func checkProductSearchReqParam(req *model.CustomerProductSearchReq) model.Payload {
	if len(req.KeyWord) < SearchKeyWordLenMin || len(req.KeyWord) > SearchKeyWordLenMax {
		return errno.ERR_SEARCH_KEY_LEN
	}
	if req.PageSize > SearchPageSizeMax {
		return errno.ERR_PAGESIZE
	}
	return errno.OK(nil)
}
