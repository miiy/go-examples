package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for i := 0; i <= 20; i++ {
		rst := isActionAllowed(rdb, ctx, "test", "reply", 60, 5)
		fmt.Println(rst)
	}

}

// 简单限流
// 指定用户 userId 的某个行为 actionKey 在特定的时间 period 只允许发生最多的次数 maxCount
// Redis 深度历险 1.8.1 如何使用 Redis 来实现简单限流策略
func isActionAllowed(rdb *redis.Client, ctx context.Context, userID, actionKey string, period, maxCount int) bool {
	key := fmt.Sprintf("hist:%s:%s", userID, actionKey)
	// 毫秒时间戳 如：1631516417827
	nowTs := time.Now().UnixNano() / 1e6

	pipe := rdb.Pipeline()
	// 记录行为：value 和 score 都使用毫秒时间戳
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(nowTs),
		Member: float64(nowTs),
	})
	// 移除窗口之前的行为记录， 剩下的都是时间窗口的
	pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(nowTs-int64(period*1000), 10))
	// 获取窗口内的行为数量
	count := pipe.ZCard(ctx, key)
	// 设置 zset 过期时间，避免冷用户持续占用内存
	// 过期时间应该等于时间窗口的长度，在多宽限 1s
	pipe.Expire(ctx, key, time.Second*time.Duration(period+1))
	// 批量执行
	cmds, err := pipe.Exec(ctx)
	_ = pipe.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cmds)

	// 比较长度是否超标
	return count.Val() <= int64(maxCount)
}
