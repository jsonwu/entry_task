package model

type SellerShop struct {
	ShopID       string `gorm:column:shop_id`
	Name         string `gorm:column:name`
	UserName     string
	Introduction string
	Level        int8
	Location     string
	CreateTime   uint32 `gorm:"autoCreateTime"`
	UpdateTime   uint32 `gorm:"autoUpdateTime"`
}

func (SellerShop) TableName() string {
	return "seller_shop_tab"
}

type ShopProduct struct {
	ShopID        string
	ProductId     string
	ProductStatus string
	CreateTime    uint32 `gorm:"autoCreateTime"`
	UpdateTime    uint32 `gorm:"autoUpdateTime"`
}
