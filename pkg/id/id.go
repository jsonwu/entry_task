package id

import "github.com/satori/go.uuid"

func NewSessionID() string {
	return uuid.NewV4().String()
}
func NewProductID() string {
	return uuid.NewV4().String()
}
func NewShopID() string {
	return uuid.NewV4().String()
}
