package database

import (
	"entry_task/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *MyDB) CreateProduct(productInfo *model.Product) error {
	result := db.db.Create(productInfo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MyDB) CreateProductWithAttr(productInfo *model.Product, attr []*model.ProductAttr) error {
	err := db.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(productInfo)
		if result.Error != nil {
			return result.Error
		}
		if len(attr) == 0 {
			return nil
		}
		result = tx.Create(&attr)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err
}

func (db *MyDB) GetProductInfo(shopID string, productID string) (*model.Product, error) {
	var products []*model.Product
	result := db.db.Where(&model.Product{ShopID: shopID, ProductID: productID}).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(products) == 0 {
		return nil, nil
	}
	return products[0], nil
}

func (db *MyDB) GetProductAndAttr(shopID string, productID string) (*model.Product, []*model.ProductAttr, error) {
	var products []*model.Product
	var attrs []*model.ProductAttr
	err := db.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where(&model.Product{ShopID: shopID, ProductID: productID}).Find(&products)
		if result.Error != nil {
			return result.Error
		}
		if len(products) == 0 {
			return nil
		}

		result = tx.Where(&model.ProductAttr{ProductID: productID}).Find(&attrs)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	if len(products) == 0 {
		return nil, nil, nil
	}
	return products[0], attrs, nil
}

func (db *MyDB) UpdateProduct(productInfo *model.Product) (error, int64) {
	result := db.db.Model(&model.Product{}).
		Where(&model.Product{ShopID: productInfo.ShopID, ProductID: productInfo.ProductID}).Updates(productInfo)
	return result.Error, result.RowsAffected
}

func (db *MyDB) UpdateProductWithAttr(productInfo *model.Product, attr []*model.ProductAttr) (error, int64) {
	var effected int64
	err := db.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Product{}).
			Where(&model.Product{ShopID: productInfo.ShopID, ProductID: productInfo.ProductID}).Updates(productInfo)
		if result.Error != nil {
			return result.Error
		}
		//fmt.Println(productInfo.ShopID, productInfo.ProductID,  result.RowsAffected, tx.RowsAffected, db.db.RowsAffected)
		//fmt.Println(*productInfo)
		effected = result.RowsAffected
		if len(attr) == 0 {
			return nil
		}
		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "product_id"}, {Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).Create(&attr)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err, effected
}

func (db *MyDB) UpdateProductStatus(shopID string, productID string, status model.PStatus) (error, int64) {
	result := db.db.Model(&model.Product{}).Where(model.Product{ShopID: shopID, ProductID: productID}).Update("status", status)
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, result.RowsAffected
}

func (db *MyDB) GetShopProducts(shopID string, offset int, limit int) ([]*model.Product, int64, error) {
	var products []*model.Product
	var count int64
	result := db.db.Where(model.Product{ShopID: shopID}).Offset(offset).Limit(limit).Find(&products).Limit(-1).Offset(-1).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return products, count, nil
}

func (db *MyDB) GetShopProductsWithStatus(shopID string, status model.PStatus, offset int, limit int) ([]*model.Product, int64, error) {
	var products []*model.Product
	var count int64
	result := db.db.Where(model.Product{ShopID: shopID, Status: status}).Offset(offset).Limit(limit).Find(&products).Limit(-1).Offset(-1).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return products, count, nil
}

func (db *MyDB) SearchProductInTitle(keyword string, offset int, limit int) ([]*model.Product, int64, error) {
	var products []*model.Product
	var count int64
	search := "%" + keyword + "%"
	result := db.db.Model(&model.Product{}).Where("title LIKE ?", search).Offset(offset).Limit(limit).Find(&products).Limit(-1).Offset(-1).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return products, count, nil
}

func (db *MyDB) LoadProducts(offset int, limit int) ([]*model.Product, error) {
	var products []*model.Product
	result := db.db.Offset(offset).Limit(limit).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
