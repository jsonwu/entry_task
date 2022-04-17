package id

import "github.com/google/uuid"

func NewSessionID() string {
	return uuid.New().String()
}
func NewProductID() string {
	return uuid.New().String()
}
func NewShopID() string {
	return uuid.New().String()
}
