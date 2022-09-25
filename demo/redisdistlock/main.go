package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type RedisDistLock struct {
	expire time.Duration
	rdb *redis.Client
}

var delIfEqualsScript = redis.NewScript(`
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	lock := &RedisDistLock{
		expire: time.Second * 5,
		rdb: rdb,
	}
	ctx := context.Background()

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			randomStr := strconv.Itoa(rand.Int())
			if lock.lock(ctx, "test-lock", randomStr) {
				fmt.Printf("%d: get lock! %s\n", i, randomStr)
				lock.unlock(ctx, "test-lock", randomStr)
			} else {
				fmt.Printf("%d: not get lock\n", i)
			}
		}(i)
	}

	wg.Wait()
}

func (l *RedisDistLock) lock(ctx context.Context, key string, val string) bool {
	// set key "" ex 5 nx
	bc := l.rdb.SetNX(ctx, key, val, l.expire)
	if bc.Err() != nil {
		panic(bc.Err())
	}
	return bc.Val()
}


func (l *RedisDistLock) unlock(ctx context.Context, key string, val string) bool {
	ic := delIfEqualsScript.Run(ctx, l.rdb, []string{key}, val)
	if ic.Err() != nil {
		panic(ic.Err())
	}
	// log.Println(ic)
	return ic.Val() == 1
}

func (l *RedisDistLock) unlock2(ctx context.Context, key string) bool {
	ic := l.rdb.Del(ctx, key)
	if ic.Err() != nil {
		panic(ic.Err())
	}
	return ic.Val() == 1
}
