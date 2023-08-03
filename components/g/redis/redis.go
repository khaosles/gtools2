package redis

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/khaosles/gtools2/core/config"
	"go.uber.org/zap"

	glog "github.com/khaosles/gtools2/core/log"
	gerr "github.com/khaosles/gtools2/utils/err"
)

/*
   @File: redis.go
   @Author: khaosles
   @Time: 2023/3/3 22:34
   @Desc:
*/

var (
	RDB  *redis.Client
	once sync.Once
)

func Init(cfg *config.Redis) {
	once.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:               cfg.Addr,
			Password:           cfg.Password,
			DB:                 cfg.DB,
			MaxRetries:         cfg.MaxRetries,
			DialTimeout:        cfg.DialTimeout,
			ReadTimeout:        cfg.ReadTimeout,
			WriteTimeout:       cfg.WriteTimeout,
			PoolSize:           cfg.PoolSize,
			MinIdleConns:       cfg.MinIdleConns,
			MaxConnAge:         cfg.MaxConnAge,
			PoolTimeout:        cfg.PoolTimeout,
			IdleTimeout:        cfg.IdleTimeout,
			IdleCheckFrequency: cfg.IdleCheckFrequency,
		})
		pong, err := RDB.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln("redis connect ping failed, err:", zap.Error(err))
		} else {
			glog.Debug("redis connect ping response:", zap.String("pong", pong))
		}
	})
}

//  /////////////////////////// 字符串 ///////////////////////////

// Set 设置string
func Set(key string, value interface{}) error {
	err := RDB.Set(context.Background(), key, value, 0).Err()
	return err
}

// SetExpire 设置string
func SetExpire(key string, value interface{}, expireTime time.Duration) error {
	err := RDB.Set(context.Background(), key, value, expireTime).Err()
	return err
}

// Get 获取string
func Get(key string) (string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return "", gerr.RedisKeyNotFoundException.New(key)
	}
	// 获取key
	val, err := RDB.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

//  /////////////////////////// hash表 ///////////////////////////

// HSet 设置单个哈希表字段
func HSet(key, field string, value any) error {
	err := RDB.HSet(context.Background(), key, field, value).Err()
	return err
}

// HSetMap 设置map对象位hash表
func HSetMap(key string, fields map[string]any) error {
	// 遍历字段
	for field, value := range fields {
		err := HSet(key, field, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// HSetMapExpire 设置map对象位hash表
func HSetMapExpire(key string, fields map[string]any, expireTime time.Duration) error {
	// 遍历字段
	for field, value := range fields {
		err := HSet(key, field, value)
		if err != nil {
			return err
		}
	}
	return ExpireTime(key, expireTime)
}

// HGet 获取hash字段值
func HGet(key, field string) (string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return "", gerr.RedisKeyNotFoundException.New(key)
	}
	// 获取数据
	result, err := RDB.HGet(context.Background(), key, field).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// HGetAll 获取所有key 转成map
func HGetAll(key string) (map[string]string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return nil, gerr.RedisKeyNotFoundException.New(key)
	}
	// 获取数据
	fields, err := RDB.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

// HExists 检查哈希表中给定字段是否存在
func HExists(key, field string) (bool, error) {
	return RDB.HExists(context.Background(), key, field).Result()
}

// HDel 删除哈希表中的一个或多个字段
func HDel(key string, fields ...string) (int64, error) {
	return RDB.HDel(context.Background(), key, fields...).Result()
}

// HKeys 获取哈希表中所有的字段
func HKeys(key string) ([]string, error) {
	return RDB.HKeys(context.Background(), key).Result()
}

// HVals 获取哈希表中所有字段的值
func HVals(key string) ([]string, error) {
	return RDB.HVals(context.Background(), key).Result()
}

// HLen 获取哈希表中字段的数量
func HLen(key string) (int64, error) {
	return RDB.HLen(context.Background(), key).Result()
}

//  /////////////////////////// 列表 ///////////////////////////

// LPush 将元素插入 Redis 的 list 头部
func LPush(key string, values ...interface{}) error {
	return RDB.LPush(context.Background(), key, values...).Err()
}

// RPush 将元素插入 Redis 的 list 尾部
func RPush(key string, values ...interface{}) error {
	return RDB.RPush(context.Background(), key, values...).Err()
}

// LRange 获取 Redis 的 list 中指定范围的元素
func LRange(key string, start, stop int64) ([]string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return []string{}, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.LRange(context.Background(), key, start, stop).Result()
}

// LIndex 获取 Redis 的 list 中指定值的元素
func LIndex(key string, index int64) (string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return "", gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.LIndex(context.Background(), key, index).Result()
}

// LPop 获取 Redis 的 list 删除左边的元素
func LPop(key string) (string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return "", gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.LPop(context.Background(), key).Result()
}

// RPop 获取 Redis 的 list 删除右边的元素
func RPop(key string) (string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return "", gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.RPop(context.Background(), key).Result()
}

//  /////////////////////////// 无序集合 ///////////////////////////

// SAdd 无序集合添加元素
func SAdd(key string, members ...interface{}) error {
	return RDB.SAdd(context.Background(), key, members...).Err()
}

// SMembers 无序集合获取所有成员
func SMembers(key string) ([]string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return []string{}, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.SMembers(context.Background(), key).Result()
}

// SIsMember 无序集合判断是否属于该集合
func SIsMember(key string, member interface{}) (bool, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return false, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.SIsMember(context.Background(), key, member).Result()
}

// SRem 无序集合删除成员
func SRem(key string, members ...interface{}) error {
	// 判断key是否存在
	if !ExistKey(key) {
		return gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.SRem(context.Background(), key, members...).Err()
}

//  /////////////////////////// 有序集合 ///////////////////////////

// ZAdd 添加元素
func ZAdd(key string, members ...*redis.Z) error {
	return RDB.ZAdd(context.Background(), key, members...).Err()
}

// ZRange 获取范围内元素
func ZRange(key string, start, stop int64) ([]string, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return []string{}, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.ZRange(context.Background(), key, start, stop).Result()
}

// ZRank 有序集合中指定成员的序号
func ZRank(key string, member string) (int64, error) {
	// 判断key是否存在
	if !ExistKey(key) {
		return -1, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.ZRank(context.Background(), key, member).Result()
}

// ZRem 删除
func ZRem(key string, members ...interface{}) error {
	// 判断key是否存在
	if !ExistKey(key) {
		return gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.ZRem(context.Background(), key, members...).Err()
}

// Rename 更改key的名字
func Rename(oldKey, newKey string) (bool, error) {
	// 判断key是否存在
	if !ExistKey(oldKey) {
		return false, gerr.RedisKeyNotFoundException.New(oldKey)
	}
	return RDB.RenameNX(context.Background(), oldKey, newKey).Result()
}

// ExistKey key是否存在
func ExistKey(key string) bool {
	ok, _ := RDB.Exists(context.Background(), key).Result()
	return ok == 1
}

// Delete 删除key
func Delete(keys ...string) error {
	return RDB.Del(context.Background(), keys...).Err()
}

// ExpireTime 设置key过期时间
func ExpireTime(key string, t time.Duration) error {
	// 判断key是否存在
	if !ExistKey(key) {
		return gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.Expire(context.Background(), key, t).Err()
}

// GetExpire 获得 key 的过期时间
func GetExpire(key string) (time.Duration, error) {
	if !ExistKey(key) {
		return -1, gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.TTL(context.Background(), key).Result()
}

// RemoveExpire 删除key过期时间
func RemoveExpire(key string) error {
	// 判断key是否存在
	if !ExistKey(key) {
		return gerr.RedisKeyNotFoundException.New(key)
	}
	return RDB.Persist(context.Background(), key).Err()
}

// Subscribe 订阅消息
func Subscribe(channel string, cb func(string)) {
	pubsub := RDB.Subscribe(context.Background(), channel)
	defer pubsub.Close()
	// 处理接收到的消息
	for msg := range pubsub.Channel() {
		cb(msg.Payload)
	}
}

// Publish 发布消息
func Publish(channel, message string) {
	RDB.Publish(context.Background(), channel, message)
}
