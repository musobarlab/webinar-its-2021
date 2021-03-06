package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// RedisCache model
type RedisCache struct {
	readPool, writePool *redis.Pool
	expired             time.Duration
}

// NewRedisCache constructor
func NewRedisCache(read, write *redis.Pool, expired time.Duration) *RedisCache {
	return &RedisCache{
		readPool: read, writePool: write, expired: expired,
	}
}

// Find method
func (r *RedisCache) Find(ctx context.Context, target interface{}) error {
	conn := r.readPool.Get()
	defer conn.Close()

	key, _ := shared.GetValueFromContext(ctx, shared.ContextKey("cacheKey")).(string)
	key = helper.GenerateHMAC(key)

	res, err := conn.Do("GET", key)
	if err != nil {
		return err
	}

	if res == nil {
		err = errors.New(helper.DataNotFound)
		return err
	}

	data, _ := res.([]byte)
	return json.Unmarshal(data, target)
}

// Save method
func (r *RedisCache) Save(ctx context.Context, value interface{}) error {
	conn := r.writePool.Get()
	defer conn.Close()

	b, _ := json.Marshal(value)

	key, _ := shared.GetValueFromContext(ctx, shared.ContextKey("cacheKey")).(string)
	key = helper.GenerateHMAC(key)

	exp, ok := shared.GetValueFromContext(ctx, shared.ContextKey("cacheExp")).(int)
	if !ok {
		exp = int(r.expired.Seconds())
	}

	// _, err := conn.Do("SETEX", key, int(domain.LoginAttemptExpired.Seconds()), string(b))
	_, err := conn.Do("SETEX", key, exp, string(b))
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// Delete method
func (r *RedisCache) Delete(ctx context.Context) error {
	conn := r.writePool.Get()
	defer conn.Close()

	key, _ := shared.GetValueFromContext(ctx, shared.ContextKey("cacheKey")).(string)
	key = helper.GenerateHMAC(key)
	_, err := conn.Do("DEL", key)

	return err
}

// GetTTL method
func (r *RedisCache) GetTTL(ctx context.Context, key string) (i int) {
	conn := r.writePool.Get()
	defer conn.Close()

	key = helper.GenerateHMAC(key)
	reply, err := conn.Do("TTL", key)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := reply.(int64)
	return int(data)
}
