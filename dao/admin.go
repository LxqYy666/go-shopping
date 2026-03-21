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
