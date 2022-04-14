package database

import "entry_task/model"

func (db *MyDB) CreateAttr(attr model.Attr) error {
	result := db.db.Create(attr)
	return result.Error
}

func (db *MyDB) LoadAttr() ([]model.Attr, error) {
	var attrs []model.Attr
	result := db.db.Model(&model.Attr{}).Find(&attrs)
	if result.Error != nil {
		return nil, result.Error
	}
	return attrs, nil
}
