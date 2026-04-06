package dao

import (
	"fmt"
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"
)

func GetCartItems(userID uint) ([]net.CartItemData, error) {
	var cartItems []models.Cart
	err := utils.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	result := make([]net.CartItemData, len(cartItems))
	for i, item := range cartItems {
		result[i] = net.CartItemData{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if item.Product != nil {
			result[i].Product = net.ProductInfoReqData{
				ID:         item.Product.ID,
				Name:       item.Product.Name,
				Desc:       item.Product.Desc,
				CategoryId: item.Product.CategoryID,
				Price:      item.Product.Price,
				Stock:      item.Product.Stock,
				ImageUrl:   item.Product.ImageUrl,
				SoldCount:  item.Product.SoldCount,
				Status:     item.Product.Status,
			}
		}
	}

	return result, nil
}

func AddToCart(userID uint, req net.AddToCartReq) error {
	// 检查商品是否存在
	var product models.Product
	err := utils.DB.First(&product, req.ProductID).Error
	if err != nil {
		return fmt.Errorf("商品不存在")
	}

	// 检查库存
	if product.Stock < req.Quantity {
		return fmt.Errorf("库存不足")
	}

	// 检查是否已存在
	var existingCart models.Cart
	err = utils.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&existingCart).Error
	if err == nil {
		// 已存在，更新数量
		newQuantity := existingCart.Quantity + req.Quantity
		if product.Stock < newQuantity {
			return fmt.Errorf("库存不足")
		}
		return utils.DB.Model(&existingCart).Update("quantity", newQuantity).Error
	}

	// 新增购物车项
	cart := models.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	return utils.DB.Create(&cart).Error
}

func UpdateCartItem(userID uint, cartID uint, req net.UpdateCartReq) error {
	var cart models.Cart
	err := utils.DB.Where("id = ? AND user_id = ?", cartID, userID).Preload("Product").First(&cart).Error
	if err != nil {
		return fmt.Errorf("购物车项不存在")
	}

	// 检查库存
	if cart.Product != nil && cart.Product.Stock < req.Quantity {
		return fmt.Errorf("库存不足")
	}

	return utils.DB.Model(&cart).Update("quantity", req.Quantity).Error
}

func DeleteCartItem(userID uint, cartID uint) error {
	return utils.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&models.Cart{}).Error
}
