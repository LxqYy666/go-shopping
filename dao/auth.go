package dao

import (
	"errors"
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"

	"gorm.io/gorm"
)

func HasAUser(loginInfo net.LoginReq) (models.User, error) {
	var user models.User
	err := utils.DB.Where("username = ? and password = ?", loginInfo.Username, loginInfo.Password).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateAUser(registerInfo net.RegisterReq) error {
	var user models.User

	err := utils.DB.Take(&user, "username = ?", registerInfo.Username).Error
	if err == nil {
		return errors.New("该用户名已经被使用")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = utils.DB.Take(&user, "email = ?", registerInfo.Email).Error
	if err == nil {
		return errors.New("该邮箱已经被使用")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	newUser := models.User{Username: registerInfo.Username, Email: registerInfo.Email, Password: registerInfo.Password}
	err = utils.DB.Create(&newUser).Error
	return err
}
