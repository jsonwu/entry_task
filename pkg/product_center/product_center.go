package product_center

import (
	"context"
	"entry_task/database"
	"entry_task/errno"
	"entry_task/model"
)

type ProductCenter struct {
	db *database.MyDB
}

func NewProductCenter(db *database.MyDB) *ProductCenter {
	return &ProductCenter{db: db}
}

func (p *ProductCenter) GetProduct(ctx context.Context, shopID string, productID string) (*model.Product, model.Payload) {
	return nil, errno.OK(nil)
}

func (p *ProductCenter) UpdateProductStatus(ctx context.Context, shopID string, productID string, status model.PStatus) (*model.Product, model.Payload) {
	return nil, errno.OK(nil)
}
