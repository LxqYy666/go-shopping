package dao

import (
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"
)

func HasAUser(loginInfo net.LoginReq) (models.User, error) {
	var user models.User
	err := utils.DB.Where("username = ? and password = ?", loginInfo.Username, loginInfo.Password).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
