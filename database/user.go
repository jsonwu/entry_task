package database

import (
	"entry_task/model"
)

func (db *MyDB) CreateUser(user model.User) error {
	result := db.db.Create(&user)
	return result.Error
}

func (db *MyDB) GetUser(userName string, userType model.UserType) (*model.User, error) {
	var users []model.User
	result := db.db.Where(&model.User{UserName: userName, UserType: userType}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &users[0], nil
}
