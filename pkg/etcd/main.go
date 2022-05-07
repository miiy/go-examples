package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	cli, cleanup, err := NewClient()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = basic(cli); err != nil {
		panic(err)
	}

	if err = lease(cli); err != nil {
		panic(err)
	}
}

func NewClient() (*clientv3.Client, func(), error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"http://127.0.0.1:2379"},
		DialTimeout:          5 * time.Second,
	})
	if err != nil {
		return nil, nil, err
	}

	 cleanup := func() {
		defer cli.Close()
	}

	return cli, cleanup, nil
}

func basic(cli *clientv3.Client) error {
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

func lease(cli *clientv3.Client) error {
	grantResp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		return err
	}
	fmt.Println(grantResp)

	putResp, err := cli.Put(context.Background(), "/test/lease", "10s", clientv3.WithLease(grantResp.ID))
	if err != nil {
		return err
	}
	fmt.Print(putResp)

	return nil
}