package handler

import (
	"entry_task/errno"
	"entry_task/model"
)

type createAccountReq struct {
	UserName   string         `json:"user_name"`
	Password   string         `json:"password"`
	UserType   model.UserType `json:"user_type"`
	Email      string         `json:"email"`
	ProfileUri string         `json:"profile_uri"`
}

func (r createAccountReq) isValidate() errno.Payload {
	if len(r.UserName) == 0 || len(r.UserName) > 20 {
		return errno.ERR_USER_NAME_LEN
	}
	if len(r.Password) == 0 || len(r.Password) > 20 {
		return errno.ERR_PASSWORD_LEN
	}
	return errno.OK(nil)
}

type loginReq struct {
	UserName string         `form:"user_name"  binding:"required"`
	Password string         `form:"password" binding:"required"`
	UserType model.UserType `form:"user_type" binding:"required"`
}

type GetUserInfoResp struct {
	UserName   string         `json:"user_name"`
	UserType   model.UserType `json:"user_type"`
	Email      string         `json:"email"`
	ProfileUri string         `json:"profile_uri"`
}

type createShopReq struct {
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type createShopResp struct {
	ShopID string `json:"shop_id" `
}

type sellerShopListResp struct {
	ShopBaseInfos []ShopBaseInfo `json:"shop_base_infos"`
}

type sellerGetShopInfoReq struct {
	ShopID string `form:"shop_id" binding:"required"`
}

type sellerGetShopInfoResp struct {
	ShopInfo
}

type createProductReq struct {
	ProductInfo
}

type createProductResp struct {
	ProductID string `json:"product_id"`
}

type updateProductReq struct {
	ProductInfo
}

func (r updateProductReq) isValidate() errno.Payload {
	return errno.OK(nil)
}

type SellerGetShopProductsReq struct {
	ShopID   string `form:"shop_id" binding:"required"`
	Status   int    `form:"status"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size" `
}

type sellerGetShopProductsResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum          int               `json:"page_number"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type UpdateProductStatusReq struct {
	ShopID     string `json:"shop_id"`
	ProductID  string `json:"product_id"`
	Status     uint8  `json:"status"`
	ProviderID string `json:"provider_id"`
}

type customerShopListReq struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type getProductInfoReq struct {
	ShopID    string `form:"shop_id" binding:"required"`
	ProductID string `form:"product_id" binding:"required"`
}

type productInfoReqResp struct {
	ProductInfo
}

type customerShopListResp struct {
	ShopBaseInfos []ShopBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum       int            `json:"page_number"`
	PageSize      int            `json:"page_size"`
	TotalCount    int64          `json:"total_count"`
}

type customerGetShopProductsReq struct {
	ShopID   string `form:"shop_id" binding:"required"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}

func (c customerGetShopProductsReq) isValidate() errno.Payload {
	if c.PageSize < 0 || c.PageSize > 100 {
		return errno.ERR_PAGESIZE
	}
	if c.PageNum < 0 || c.PageNum > 100 {
		return errno.ERR_PAGESIZE
	}
	if len(c.ShopID) == 0 || len(c.ShopID) > 36 {
		return errno.ERR_INVALID_PARAM
	}

	return errno.OK(nil)
}

type customerGetShopProductsResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum          int               `json:"page_num"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type customerProductSearchReq struct {
	KeyWord  string `form:"keyword" binding:"required"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}

type customerSearchResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"product_base_infos,omitempty"`
	PageNum          int               `json:"page_number"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type feedReq struct {
	Num int `form:"num"`
}
type feedResp struct {
	ProductBaseInfos []ProductBaseInfo
}

type ShopBaseInfo struct {
	ShopID       string `json:"shop_id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction,omitempty"`
}

type ShopInfo struct {
	ShopID       string `json:"shop_id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction,omitempty"`
	Level        int8   `json:"level,omitempty"`
	location     string `json:"location,omitempty"`
}

type AttrInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductInfo struct {
	ShopID     string     `json:"shop_id"`
	ProductID  string     `json:"product_id"`
	Title      string     `json:"title"`
	CoverUri   string     `json:"cover_uri"`
	ShowUris   []string   `json:"show_uris"`
	Price      uint32     `json:"price"`
	Stock      uint32     `json:"stock"`
	BrandID    int64      `json:"brand_id"`
	CategoryID int64      `json:"category_id"`
	Details    []string   `json:"details"`
	AttrInfo   []AttrInfo `json:"attr_info"`
}

type ProductBaseInfo struct {
	ProductID string `json:"product_id"`
	ShopID    string `json:"shop_id"`
	Title     string `json:"title"`
	CoverUri  string `json:"cover_uri"`
	Price     uint32 `json:"price"`
}
