package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"go-shopping/config"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const OrderTimeoutQueueKey = "orders:timeout:zset"

var (
	RedisClient *redis.Client
	redisCtx    = context.Background()
)

func InitRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := client.Ping(redisCtx).Err(); err != nil {
		return err
	}
	RedisClient = client
	return nil
}

func GetCache[T any](key string) (T, bool, error) {
	var zero T
	if RedisClient == nil {
		return zero, false, nil
	}

	raw, err := RedisClient.Get(redisCtx, key).Result()
	if err == redis.Nil {
		return zero, false, nil
	}
	if err != nil {
		return zero, false, err
	}

	var value T
	if err := json.Unmarshal([]byte(raw), &value); err != nil {
		return zero, false, err
	}
	return value, true, nil
}

func SetCache(key string, value any, ttl time.Duration) error {
	if RedisClient == nil {
		return nil
	}

	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisClient.Set(redisCtx, key, body, ttl).Err()
}

func DeleteCache(keys ...string) error {
	if RedisClient == nil || len(keys) == 0 {
		return nil
	}
	return RedisClient.Del(redisCtx, keys...).Err()
}

func EnqueueOrderTimeout(orderID uint, deadline time.Time) error {
	if RedisClient == nil {
		return fmt.Errorf("redis is not initialized")
	}
	return RedisClient.ZAdd(redisCtx, OrderTimeoutQueueKey, redis.Z{
		Score:  float64(deadline.Unix()),
		Member: strconv.FormatUint(uint64(orderID), 10),
	}).Err()
}

func FetchExpiredOrderIDs(now time.Time, limit int64) ([]uint, error) {
	if RedisClient == nil {
		return nil, fmt.Errorf("redis is not initialized")
	}

	values, err := RedisClient.ZRangeByScore(redisCtx, OrderTimeoutQueueKey, &redis.ZRangeBy{
		Min:   "-inf",
		Max:   strconv.FormatInt(now.Unix(), 10),
		Count: limit,
	}).Result()
	if err != nil {
		return nil, err
	}

	orderIDs := make([]uint, 0, len(values))
	for _, value := range values {
		id, convErr := strconv.ParseUint(value, 10, 64)
		if convErr != nil {
			continue
		}
		orderIDs = append(orderIDs, uint(id))
	}
	return orderIDs, nil
}

func RemoveOrderTimeout(orderID uint) error {
	if RedisClient == nil {
		return fmt.Errorf("redis is not initialized")
	}
	return RedisClient.ZRem(redisCtx, OrderTimeoutQueueKey, strconv.FormatUint(uint64(orderID), 10)).Err()
}
