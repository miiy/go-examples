package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	err := NewClient()
	if err != nil {
		panic(err)
	}

}

func NewClient() error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"http://127.0.0.1:2379"},
		DialTimeout:          5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer cli.Close()


	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	putResp, err := cli.Put(ctx, "/service/pay", "pay")
	cancel()
	if err != nil {
		return err
	}
	fmt.Println(putResp)

	getResp, err := cli.Get(context.Background(), "/service/pay")
	if err != nil {
		return err
	}
	fmt.Println(getResp)
	for _, kv := range getResp.Kvs {
		fmt.Printf("key: %s, value: %s\n", kv.Key, kv.Value)
	}

	deleteResp, err := cli.Delete(context.Background(), "/service/pay")
	if err != nil {
		return err
	}
	fmt.Println(deleteResp)
	return nil
}
