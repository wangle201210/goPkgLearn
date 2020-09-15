package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)
var ClientRedis *redis.Client

var c context.Context

const (
	REDIS_NETWORK   = "tcp"
	REDIS_HOST      = ""
	REDIS_PORT      = "6379"
	REDIS_PASSWORD  = ""
	REDIS_DB        = 0
)

func main() {
	defer ClientRedis.Close()
	c = context.Context(context.Background())
	String()
	Hash()
	Set()
	Zset()
}

func init() {
	options := redis.Options{
		Network:            REDIS_NETWORK,
		Addr:               fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           REDIS_PASSWORD,
		DB:                 REDIS_DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	}
	// 新建一个client
	ClientRedis = redis.NewClient(&options)
}

func String() {
	// 添加string
	ClientRedis.Set(c,"golang_redis", "golang", 0)
	ClientRedis.Set(c,"golang_string", "golang", 0)
	// 获取string
	stringCmd := ClientRedis.Get(c,"golang_redis")
	fmt.Println(stringCmd.String(), stringCmd.Args(), stringCmd.Val())
	// 删除string
	ClientRedis.Del(c,"golang_redis")
}

func Hash() {
	// hash - 添加field
	ClientRedis.HSet(c,"golang_hash", "key_1", "val_1", "key_2", "val_2")
	ClientRedis.HSet(c,"golang_hash", []string{"key_3", "val_3", "key_4", "val_4"})
	// hash - 获取一个field
	hCmd := ClientRedis.HGet(c,"golang_hash", "user")
	fmt.Println(hCmd.String(), hCmd.Err(), hCmd.Val())
	// hash - 获取长度
	cmd := ClientRedis.HLen(c,"golang_hash")
	fmt.Println(cmd.String(), cmd.Args(), cmd.Val())
	// hash - 获取全部
	cmdAll := ClientRedis.HGetAll(c,"golang_hash")
	fmt.Println(cmdAll.String(), cmdAll.Args(), cmdAll.Val())
	// hash - 获取多个key值
	hmCmd := ClientRedis.HMGet(c,"golang_hash", "key_1", "key_2")
	fmt.Println(hmCmd.String(), hmCmd.Args(), hmCmd.Val())
	// hash - 添加field，没有发现和HSet有什么区别
	ClientRedis.HMSet(c,"golang_hash", "key_5", "val_5", "key_6", "val_6")
	ClientRedis.HMSet(c,"golang_hash", []string{"key_7", "val_7", "key_8", "val_8"})
	// hash - 删除field
	ClientRedis.HDel(c,"golang_hash", "key_1", "key_2", "key_3")
}

func List() {
	// list - 从左追加 index 0 val val_3,index 1 val val_2
	ClientRedis.LPush(c,"golang_list", "val_2", "val_3")
	// list - 从左追加 index 0 val val_5,index 1 val val_4
	ClientRedis.LPushX(c,"golang_list", "val_4", "val_5")
	// list - 从右追加 index -1 val val_10,index -2 val val_9
	ClientRedis.RPush(c,"golang_list", "val_9", "val_10")
	// list - 通过index设置val
	ClientRedis.LSet(c,"golang_list", 0, "val_1")
	// list - 通过index获取val
	stringCmd := ClientRedis.LIndex(c,"golang_list", 0)
	fmt.Println(stringCmd.String(), stringCmd.Args(), stringCmd.Val())
	// list - 获取长度
	lenCmd := ClientRedis.LLen(c,"golang_list")
	fmt.Println(lenCmd.String(), lenCmd.Args(), lenCmd.Val())
	// list - 从左删除
	ClientRedis.LPop(c,"golang_list")
	// list - 取全部
	listCmd := ClientRedis.LRange(c,"golang_list", 0, -1)
	fmt.Println(listCmd.String(), listCmd.Args(), listCmd.Val())
	// list - 截取，当start为负数时从右向左截取
	ltrimCmd := ClientRedis.LTrim(c,"golang_list", 2, 3)
	fmt.Println(ltrimCmd.String(), ltrimCmd.Args(), ltrimCmd.Val())
}

func Set() {
	// 无序集合 ("马超", "关羽", "赵云") 后面的赵云会覆盖前面的赵云
	ClientRedis.SAdd(c,"golang_set", "马超", "赵云", "关羽", "张飞", "曹植", "司马懿")
	// 从右删除，并返回
	sCmd := ClientRedis.SPop(c,"golang_set")
	fmt.Println(sCmd.String(), sCmd.Args(), sCmd.Val())
	// 指定删除
	ClientRedis.SRem(c,"golang_set", "赵云")
	// 获取集合成员
	sMembers := ClientRedis.SMembers(c,"golang_set")
	fmt.Println(sMembers.String(), sMembers.Args(), sMembers.Val())
	// 返回集合成员数
	cCmd := ClientRedis.SCard(c,"golang_set")
	fmt.Println(cCmd.String(), cCmd.Args(), cCmd.Val())
}

func Zset() {
	// 新增
	ClientRedis.ZAdd(c,"golang_zset", &redis.Z{
		Score:  0,
		Member: "张飞",
	}, &redis.Z{
		Score:  1,
		Member: "关羽",
	}, &redis.Z{
		Score:  6,
		Member: "刘备",
	})
	// 取区间统计 分数值在 0 和 7 之间的成员的数量
	cCmd := ClientRedis.ZCount(c,"golang_zset", "0", "7")
	fmt.Println(cCmd.String(), cCmd.Args(), cCmd.Val())
	// 取成员数
	zCard := ClientRedis.ZCard(c,"golang_zset")
	fmt.Println(zCard.String(), zCard.Args(), zCard.Val())
	// 通过索引区间返回有序集合指定区间内的成员
	zRange := ClientRedis.ZRange(c,"golang_zset", 0, 1)
	fmt.Println(zRange.String(), zRange.Args(), zRange.Val())
	// 移除有序集合中给定的分数区间的所有成员
	ClientRedis.ZRemRangeByScore(c,"golang_zset", "0", "1")
}
