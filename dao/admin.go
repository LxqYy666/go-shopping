package dao

import (
	"fmt"
	"go-shopping/models"
	"go-shopping/net"
	"go-shopping/utils"

	"gorm.io/gorm"
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

func AddProduct(req net.AddProductReq) error {
	product := models.Product{
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Desc:       req.Desc,
		Price:      req.Price,
		Stock:      req.Stock,
		ImageUrl:   req.ImageUrl,
		Status:     "on",
	}
	return utils.DB.Create(&product).Error
}

func UpdateProduct(id uint, req net.UpdateProductReq) error {
	updates := make(map[string]interface{})
	if req.CategoryID != nil {
		updates["category_id"] = *req.CategoryID
	}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Desc != nil {
		updates["desc"] = *req.Desc
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.Stock != nil {
		updates["stock"] = *req.Stock
	}
	if req.ImageUrl != nil {
		updates["image_url"] = *req.ImageUrl
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	
	return utils.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}

func DeleteProduct(id uint) error {
	return utils.DB.Delete(&models.Product{}, id).Error
}

func GetOrderList() ([]net.OrderData, error) {
	var orders []models.Order
	err := utils.DB.Preload("User").Preload("Items.Product").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	orderList := make([]net.OrderData, len(orders))
	for i, order := range orders {
		orderList[i] = net.OrderData{
			ID:            order.ID,
			UserID:        order.UserID,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			ReceiverAddr:  order.ReceiverAddr,
			ReceiverName:  order.ReceiverName,
			ReceiverPhone: order.ReceiverPhone,
			Remark:        order.Remark,
			CreatedAt:     order.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		if order.User != nil {
			orderList[i].User = &net.UserInfoReqData{
				ID:       order.User.ID,
				Username: order.User.Username,
				Avatar:   order.User.Avatar,
				Role:     order.User.Role,
				Status:   order.User.Status,
			}
		}

		orderList[i].Items = make([]net.OrderItemData, len(order.Items))
		for j, item := range order.Items {
			orderList[i].Items[j] = net.OrderItemData{
				ID:         item.ID,
				ProductID:  item.ProductID,
				Quantity:   item.Quantity,
				TotalPrice: item.TotalPrice,
			}
			if item.Product != nil {
				orderList[i].Items[j].Product = &net.ProductInfoReqData{
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
	}

	return orderList, nil
}

func GetUserOrders(userID uint) ([]net.OrderData, error) {
	var orders []models.Order
	err := utils.DB.Where("user_id = ?", userID).Preload("Items.Product").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	orderList := make([]net.OrderData, len(orders))
	for i, order := range orders {
		orderList[i] = net.OrderData{
			ID:            order.ID,
			UserID:        order.UserID,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			ReceiverAddr:  order.ReceiverAddr,
			ReceiverName:  order.ReceiverName,
			ReceiverPhone: order.ReceiverPhone,
			Remark:        order.Remark,
			CreatedAt:     order.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		orderList[i].Items = make([]net.OrderItemData, len(order.Items))
		for j, item := range order.Items {
			orderList[i].Items[j] = net.OrderItemData{
				ID:         item.ID,
				ProductID:  item.ProductID,
				Quantity:   item.Quantity,
				TotalPrice: item.TotalPrice,
			}
			if item.Product != nil {
				orderList[i].Items[j].Product = &net.ProductInfoReqData{
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
	}

	return orderList, nil
}

func CreateOrder(userID uint, req net.CreateOrderReq) error {
	// 获取用户购物车
	var cartItems []models.Cart
	err := utils.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return fmt.Errorf("购物车为空")
	}

	// 计算总金额
	var totalAmount float32 = 0
	for _, item := range cartItems {
		if item.Product == nil {
			return fmt.Errorf("商品不存在")
		}
		if item.Product.Stock < item.Quantity {
			return fmt.Errorf("商品 %s 库存不足", item.Product.Name)
		}
		totalAmount += item.Product.Price * float32(item.Quantity)
	}

	// 开始事务
	tx := utils.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	order := models.Order{
		UserID:        userID,
		TotalAmount:   totalAmount,
		Status:        "pending",
		ReceiverAddr:  req.ReceiverAddr,
		ReceiverName:  req.ReceiverName,
		ReceiverPhone: req.ReceiverPhone,
		Remark:        req.Remark,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建订单明细并更新库存
	for _, cartItem := range cartItems {
		orderItem := models.OrderItem{
			OrderID:    order.ID,
			ProductID:  cartItem.ProductID,
			Quantity:   cartItem.Quantity,
			TotalPrice: cartItem.Product.Price * float32(cartItem.Quantity),
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 更新库存和销量
		if err := tx.Model(&models.Product{}).Where("id = ?", cartItem.ProductID).Updates(map[string]interface{}{
			"stock":      gorm.Expr("stock - ?", cartItem.Quantity),
			"sold_count": gorm.Expr("sold_count + ?", cartItem.Quantity),
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 清空购物车
	if err := tx.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateOrder(id uint, req net.UpdateOrderReq) error {
	updates := make(map[string]interface{})
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	return utils.DB.Model(&models.Order{}).Where("id = ?", id).Updates(updates).Error
}
