package database

import (
	"entry_task/model"
)

func (db *MyDB) GetSellerShop(userName string) ([]*model.SellerShop, error) {
	var shops []*model.SellerShop
	result := db.db.Where(&model.SellerShop{UserName: userName}).Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}
	return shops, nil
}

func (db *MyDB) CreateShop(shopInfo *model.SellerShop) error {
	result := db.db.Create(shopInfo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MyDB) GetShopInfo(shopID string) (*model.SellerShop, error) {
	var shops []*model.SellerShop
	result := db.db.Where(&model.SellerShop{ShopID: shopID}).Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(shops) == 0 {
		return nil, nil
	}
	return shops[0], nil
}

func (db *MyDB) GetShopInfoByName(shopName string) (*model.SellerShop, error) {
	var shops []*model.SellerShop
	result := db.db.Where(&model.SellerShop{Name: shopName}).Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(shops) == 0 {
		return nil, nil
	}

	return shops[0], nil
}

func (db *MyDB) GetShopList(offset int, limit int) ([]*model.SellerShop, int64, error) {
	var shops []*model.SellerShop
	var count int64
	result := db.db.Offset(offset).Limit(limit).Find(&shops).Limit(-1).Offset(-1).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return shops, count, nil
}
