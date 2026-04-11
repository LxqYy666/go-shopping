package services

import (
	"go-shopping/config"
	"go-shopping/dao"
	"log"
	"time"
)

func StartOrderTimeoutWorker() {
	go func() {
		if err := dao.ProcessExpiredOrders(100); err != nil {
			log.Printf("process expired orders failed: %v", err)
		}

		ticker := time.NewTicker(time.Duration(config.OrderCancelScanSeconds) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := dao.ProcessExpiredOrders(100); err != nil {
				log.Printf("process expired orders failed: %v", err)
			}
		}
	}()
}
