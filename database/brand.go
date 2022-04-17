package database

import (
	"entry_task/model"
)

func (db *MyDB) CreateBrand(brand *model.Brand) error {
	result := db.db.Create(brand)
	return result.Error
}

func (db *MyDB) LoadBrand() ([]*model.Brand, error) {
	var brands []*model.Brand
	result := db.db.Model(&model.Brand{}).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	return brands, nil
}
