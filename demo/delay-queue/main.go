package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	t := float64(time.Now().Unix())

	rdb.Del(ctx, "queue:msg:delay")

	// produce
	for i := 0; i < 10; i++ {
		if i > 3 {
			t += float64(i + 1) // 测试，未来 i+1 s
		}
		rdb.ZAdd(ctx, "queue:msg:delay", redis.Z{
			Score:  t,
			Member: fmt.Sprintf("msg%d:%f", i, t),
		})
	}
	fmt.Println(rdb.ZRange(ctx, "queue:msg:delay", 0, -1))

	// consume
	for {
		nt := float64(time.Now().Unix())

		ssc := rdb.ZRangeByScore(ctx, "queue:msg:delay", &redis.ZRangeBy{
			Min:    strconv.Itoa(0),
			Max:    strconv.FormatFloat(nt, 'E', -1, 64),
			Offset: 0,
			Count:  1,
		})
		s := ssc.Val()
		if 0 == len(s) {
			time.Sleep(time.Second * 1)
			continue
		}
		// handle msg
		msgInfo := strings.Split(s[0], ":")
		it, err := strconv.ParseFloat(msgInfo[1], 64)
		if err != nil {
			panic(err)
		}
		fmt.Println(msgInfo[0], time.Unix(int64(it), 0).Format("2006-01-02 15:04:05"))

		// delete
		rdb.ZRem(ctx, "queue:msg:delay", s)
	}

}
