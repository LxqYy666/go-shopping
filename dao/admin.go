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
	err := utils.DB.Preload("products").Find(&categoryList).Error
	if err != nil {
		return nil, err
	}

	categoryInfoReqDataList = make([]net.CategoryInfoReqData, len(categoryList), len(categoryInfoReqDataList))
	for i, v := range categoryInfoReqDataList {
		v.ID = categoryList[i].ID
		v.Name = categoryList[i].Name
		v.ProductCount = len(categoryList[i].Products)
	}

	return categoryInfoReqDataList, nil

}
