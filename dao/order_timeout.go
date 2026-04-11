package dao

import (
	"errors"
	"go-shopping/config"
	"go-shopping/models"
	"go-shopping/utils"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CancelOrderIfPending(orderID uint) (bool, error) {
	tx := utils.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var order models.Order
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Items").First(&order, orderID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if order.Status != "pending" {
		tx.Rollback()
		return false, nil
	}

	for _, item := range order.Items {
		if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).Updates(map[string]interface{}{
			"stock":      gorm.Expr("stock + ?", item.Quantity),
			"sold_count": gorm.Expr("GREATEST(sold_count - ?, 0)", item.Quantity),
		}).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if err := tx.Model(&models.Order{}).Where("id = ?", order.ID).Update("status", "cancelled").Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}

	if err := invalidateOrderCaches(order.UserID); err != nil {
		log.Printf("invalidate order cache failed, user_id=%d err=%v", order.UserID, err)
	}
	return true, nil
}

func ProcessExpiredOrders(limit int64) error {
	if limit <= 0 {
		limit = 100
	}

	if utils.RedisClient != nil {
		if ids, err := utils.FetchExpiredOrderIDs(time.Now(), limit); err != nil {
			log.Printf("fetch expired orders from redis failed: %v", err)
		} else {
			for _, orderID := range ids {
				_, cancelErr := CancelOrderIfPending(orderID)
				if cancelErr != nil {
					log.Printf("cancel expired order failed, order_id=%d err=%v", orderID, cancelErr)
					continue
				}
				if err := utils.RemoveOrderTimeout(orderID); err != nil {
					log.Printf("remove timeout queue failed, order_id=%d err=%v", orderID, err)
				}
			}
		}
	}

	expiredBefore := time.Now().Add(-time.Duration(config.OrderAutoCancelMinutes) * time.Minute)
	var pendingOrders []models.Order
	if err := utils.DB.Where("status = ? AND created_at <= ?", "pending", expiredBefore).Limit(int(limit)).Find(&pendingOrders).Error; err != nil {
		return err
	}

	for _, order := range pendingOrders {
		if _, err := CancelOrderIfPending(order.ID); err != nil {
			log.Printf("cancel timed out order failed, order_id=%d err=%v", order.ID, err)
			continue
		}
		if utils.RedisClient != nil {
			if err := utils.RemoveOrderTimeout(order.ID); err != nil {
				log.Printf("remove timeout queue failed, order_id=%d err=%v", order.ID, err)
			}
		}
	}

	return nil
}
