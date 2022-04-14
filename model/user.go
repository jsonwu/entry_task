package model

type UserType uint8

const (
	UserTypeCustomer UserType = iota
	UserTypeSeller
)

type User struct {
	ID         uint
	UserName   string
	Password   string
	Email      string
	Salt       string
	ProfileUri string
	UserType   UserType
	CreateTime uint
	UpdateTime uint
}

func (User) TableName() string {
	return "user_tab"
}
