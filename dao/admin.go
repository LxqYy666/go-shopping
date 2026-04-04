package dao

import (
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"
)

func AddCategory(addCategoryReq net.AddCategoryReq) error {

	var category models.Category
	category.Name = addCategoryReq.Name

	if err := utils.DB.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func GetCategoryInfo() ([]net.CategoryInfoReqData, error) {
	var categoryInfoReqDataList []net.CategoryInfoReqData

	var categoryList []models.Category
	err := utils.DB.Preload("Products").Find(&categoryList).Error
	if err != nil {
		return nil, err
	}

	categoryInfoReqDataList = make([]net.CategoryInfoReqData, len(categoryList))
	for i := range categoryInfoReqDataList {
		categoryInfoReqDataList[i].ID = categoryList[i].ID
		categoryInfoReqDataList[i].Name = categoryList[i].Name
		categoryInfoReqDataList[i].ProductCount = len(categoryList[i].Products)
	}

	return categoryInfoReqDataList, nil

}

func GetProductList() ([]net.ProductInfoReqData, error) {
	var productList []net.ProductInfoReqData

	err := utils.DB.Raw("select id,name,'desc',category_id,price,stock,image_url,sold_count,status from products").Scan(&productList).Error
	if err != nil {
		return nil, err
	}
	return productList, nil
}

func GetUserList() ([]net.UserInfoReqData, error) {
	var userList []net.UserInfoReqData
	err := utils.DB.Raw("select id,username,avatar,role,status,created_at from users").Scan(&userList).Error
	if err != nil {
		return nil, err
	}
	for i := range userList {
		err := utils.DB.Raw("select count(*) from orders where user_id = ?", userList[i].ID).Scan(&userList[i].OrdersCount).Error
		if err != nil {
			return nil, err
		}
	}
	return userList, nil
}
