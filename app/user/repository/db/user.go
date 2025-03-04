package db

import (
	"github.com/littlestar2023/common_pkg/global"
	"micro_play/pkg/model"
)

func FindUserByUserName(userName string) (user *model.User, err error) {

	user = &model.User{}

	err = global.CMP_DB.Model(&model.User{}).Where("user_name = ?", userName).First(user).Error

	return
}

func CreateUser(user *model.User) (err error) {

	return global.CMP_DB.Model(&model.User{}).Create(user).Error
}
