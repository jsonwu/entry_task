package shop_center

import (
	"context"
	"entry_task/database"
	"entry_task/model"
)

type ShopCenter struct {
	db *database.MyDB
}

func NewShopCenter(db *database.MyDB) *ShopCenter {
	return &ShopCenter{db: db}
}

func (s *ShopCenter) GetUserShops(ctx context.Context, userName string, status model.PStatus) ([]*model.ShopInfo, error) {
	return nil, nil
}

func (p *ShopCenter) Shops(ctx context.Context, offset int, limit int) ([]*model.ShopInfo, error) {
	return nil, nil
}
