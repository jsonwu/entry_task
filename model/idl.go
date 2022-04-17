package model

type Payload struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type CreateAccountReq struct {
	UserName   string   `json:"user_name"`
	Password   string   `json:"password"`
	UserType   UserType `json:"user_type"`
	Email      string   `json:"email"`
	ProfileUri string   `json:"profile_uri"`
}

type LoginReq struct {
	UserName string   `form:"user_name"  binding:"required"`
	Password string   `form:"password" binding:"required"`
	UserType UserType `form:"user_type" binding:"required"`
}

type GetUserInfoResp struct {
	UserName   string   `json:"user_name"`
	UserType   UserType `json:"user_type"`
	Email      string   `json:"email"`
	ProfileUri string   `json:"profile_uri"`
}

type CreateShopReq struct {
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type CreateShopResp struct {
	ShopID string `json:"shop_id" `
}

type SellerShopListResp struct {
	ShopBaseInfos []ShopBaseInfo `json:"shop_base_infos"`
}

type SellerGetShopInfoReq struct {
	ShopID string `form:"shop_id" binding:"required"`
}

type SellerGetShopInfoResp struct {
	ShopInfo
}

type CreateProductReq struct {
	ProductInfo
}

type CreateProductResp struct {
	ProductID string `json:"product_id"`
}

type UpdateProductReq struct {
	ProductInfo
}

type SellerGetShopProductsReq struct {
	ShopID   string `form:"shop_id" binding:"required"`
	Status   int    `form:"status"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size" `
}

type SellerGetShopProductsResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum          int               `json:"page_number"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type UpdateProductStatusReq struct {
	ShopID     string  `json:"shop_id"`
	ProductID  string  `json:"product_id"`
	Status     PStatus `json:"status"`
	ProviderID string  `json:"provider_id"`
}

type CustomerShopListReq struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type GetProductInfoReq struct {
	ShopID    string `form:"shop_id" binding:"required"`
	ProductID string `form:"product_id" binding:"required"`
}

type ProductInfoReqResp struct {
	ProductInfo
}

type CustomerShopListResp struct {
	ShopBaseInfos []ShopBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum       int            `json:"page_number"`
	PageSize      int            `json:"page_size"`
	TotalCount    int64          `json:"total_count"`
}

type CustomerGetShopProductsReq struct {
	ShopID   string `form:"shop_id" binding:"required"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}

type CustomerGetShopProductsResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"shop_base_infos,omitempty"`
	PageNum          int               `json:"page_num"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type CustomerProductSearchReq struct {
	KeyWord  string `form:"keyword" binding:"required"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}

type CustomerSearchResp struct {
	ProductBaseInfos []ProductBaseInfo `json:"product_base_infos,omitempty"`
	PageNum          int               `json:"page_number"`
	PageSize         int               `json:"page_size"`
	TotalCount       int64             `json:"total_count"`
}

type FeedReq struct {
	Num int `form:"num"`
}
type FeedResp struct {
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
	Location     string `json:"location,omitempty"`
}

type AttrInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductInfo struct {
	ShopID     string     `json:"shop_id"`
	ProductID  string     `json:"product_id"`
	Title      string     `json:"title"`
	CoverUri   string     `json:"cover_uri,omitempty"`
	ShowUris   []string   `json:"show_uris"`
	Price      uint32     `json:"price"`
	Stock      uint32     `json:"stock"`
	BrandID    int64      `json:"brand_id,omitempty"`
	CategoryID int64      `json:"category_id"`
	Status     int        `json:"status,omitempty"`
	Details    []string   `json:"details,omitempty"`
	AttrInfos  []AttrInfo `json:"attr_infos,omitempty"`
}

type ProductBaseInfo struct {
	ProductID string `json:"product_id"`
	ShopID    string `json:"shop_id"`
	Title     string `json:"title"`
	CoverUri  string `json:"cover_uri"`
	Price     uint32 `json:"price"`
}
