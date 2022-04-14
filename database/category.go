package database

import "entry_task/model"

func (db *MyDB) CreateCategory(c model.Category) error {
	result := db.db.Create(&c)
	return result.Error
}

func (db *MyDB) LoadCategory() ([]model.Category, error) {
	var cs []model.Category
	result := db.db.Model(&model.Brand{}).Find(&cs)
	if result.Error != nil {
		return nil, result.Error
	}
	return cs, nil
}
