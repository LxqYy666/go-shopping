package dao

import (
	"go-shopping/config"
	"fmt"
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"
	"log"
	"time"
)

func GetCartItems(userID uint) ([]net.CartItemData, error) {
	cacheKey := cartCacheKey(userID)
	cached, found, err := utils.GetCache[[]net.CartItemData](cacheKey)
	if err != nil {
		log.Printf("get cart cache failed, user_id=%d err=%v", userID, err)
	} else if found {
		return cached, nil
	}

	var cartItems []models.Cart
	err = utils.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error
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

	if err := utils.SetCache(cacheKey, result, time.Duration(config.CacheTTLSeconds)*time.Second); err != nil {
		log.Printf("set cart cache failed, user_id=%d err=%v", userID, err)
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
		if err := utils.DB.Model(&existingCart).Update("quantity", newQuantity).Error; err != nil {
			return err
		}
		if err := invalidateCartCache(userID); err != nil {
			log.Printf("invalidate cart cache failed, user_id=%d err=%v", userID, err)
		}
		return nil
	}

	// 新增购物车项
	cart := models.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	if err := utils.DB.Create(&cart).Error; err != nil {
		return err
	}
	if err := invalidateCartCache(userID); err != nil {
		log.Printf("invalidate cart cache failed, user_id=%d err=%v", userID, err)
	}
	return nil
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

	if err := utils.DB.Model(&cart).Update("quantity", req.Quantity).Error; err != nil {
		return err
	}
	if err := invalidateCartCache(userID); err != nil {
		log.Printf("invalidate cart cache failed, user_id=%d err=%v", userID, err)
	}
	return nil
}

func DeleteCartItem(userID uint, cartID uint) error {
	if err := utils.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&models.Cart{}).Error; err != nil {
		return err
	}
	if err := invalidateCartCache(userID); err != nil {
		log.Printf("invalidate cart cache failed, user_id=%d err=%v", userID, err)
	}
	return nil
}
