package dao

import (
	"ranbao/global"
	"ranbao/service/model"
)

func InitTable() error {
	err := global.Db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user model.User) error {
	return global.Db.Model(&model.User{}).Create(&user).Error
}

func IsUserExistByEmail(email string) bool {
	var user model.User
	if tx := global.Db.Model(&model.User{}).Where("name = ?", email).First(&user); tx.RowsAffected == 0 {
		return false
	}
	return true
}

func IsUserExistByName(name string) bool {
	var user model.User
	if tx := global.Db.Model(&model.User{}).Where("name = ?", name).First(&user); tx.RowsAffected == 0 {
		return false
	}
	return true
}

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := global.Db.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func FindUserById(userId int64) (model.User, error) {
	var user model.User
	err := global.Db.Model(&model.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func FindUserByName(name string) (model.User, error) {
	var user model.User
	err := global.Db.Model(&model.User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user model.User) error {
	return global.Db.Model(&model.User{}).Where("id = ?", user.ID).Updates(&user).Error
}
