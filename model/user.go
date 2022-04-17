package model

type UserType uint8

const (
	UserTypeNone UserType = iota
	UserTypeSeller
	UserTypeCustomer
)

type User struct {
	ID         uint
	UserName   string
	Password   string
	Email      string
	Salt       string
	ProfileUri string
	UserType   UserType
	CreateTime uint `gorm:"autoCreateTime"`
	UpdateTime uint `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user_tab"
}
