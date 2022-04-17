package user_center

import (
	"context"
	"entry_task/database"
	"entry_task/errno"
	"entry_task/model"
)

type UserCenter struct {
	db *database.MyDB
}

func NewUserCenter(db *database.MyDB) *UserCenter {
	return &UserCenter{db: db}
}

func (u *UserCenter) Login(ctx context.Context, userName string, p string, userType model.UserType) model.Payload {
	return errno.OK(nil)
}

func (u *UserCenter) CreateAccout(ctx context.Context, user *model.User) model.Payload {
	return errno.OK(nil)
}
