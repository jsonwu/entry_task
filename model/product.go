package model

type PStatus uint8

const (
	ProudctStatusNone PStatus = iota
	ProudctStatusNormal
	ProudctStatusAuditing
	ProudctStatusOff
	ProudctStatusForbiden
	ProudctStatusEdit
	ProudctStatusDeleted
)

type Product struct {
	ProductID  string
	Title      string
	ShopID     string
	CoverUri   string
	ShowUris   string
	Price      uint32
	Stock      uint32
	BrandID    int64
	CategoryID int64
	Details    string
	Status     PStatus
	CreateTime uint32 `gorm:"autoCreateTime"`
	UpdateTime uint32 `gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "product_tab"
}

type Category struct {
	ID         int64
	Name       string
	ParentID   string
	CreateTime uint32 `gorm:"autoCreateTime"`
	UpdateTime uint32 `gorm:"autoUpdateTime"`
}

func (Category) TableName() string {
	return "category_tab"
}

type ProductAttr struct {
	ProductID  string
	Name       string
	Value      string
	CreateTime uint32 `gorm:"autoCreateTime"`
	UpdateTime uint32 `gorm:"autoUpdateTime"`
}

func (ProductAttr) TableName() string {
	return "product_attr_tab"
}

type Attr struct {
	Name       string
	Desc       string
	CreateTime uint32 `gorm:"autoCreateTime"`
	UpdateTime uint32 `gorm:"autoUpdateTime"`
}

func (Attr) TableName() string {
	return "attr_tab"
}

type Brand struct {
	ID         int64
	Name       string
	Desc       string
	LogUri     string
	CreateTime uint32 `gorm:"autoCreateTime"`
	UpdateTime uint32 `gorm:"autoUpdateTime"`
}

func (Brand) TableName() string {
	return "brand_tab"
}
