package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	klog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"io"
	"log"
	"time"
)

// Let's say this is a service that means to register itself.
// First, we will set up some context.
var (
	etcdServer = "127.0.0.1:2379"      // in the change from v2 to v3, the schema is no longer necessary if connecting directly to an etcd v3 instance
	ctx        = context.Background()
)

func main() {
	clientOptions := etcdv3.ClientOptions{
		Cert:          "",
		Key:           "",
		CACert:        "",
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
		DialOptions:   nil,
		Username:      "",
		Password:      "",
	}
	// Build the client.
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, clientOptions)
	if err != nil {
		panic(err)
	}

	callFoo(client)

}

func callFoo(client etcdv3.Client) error {
	fooPrefix := "/services/foosvc"
	logger := klog.NewNopLogger()
	instancer, err := etcdv3.NewInstancer(client, fooPrefix, logger)
	if err != nil {
		return err
	}
	endpointer := sd.NewEndpointer(instancer, fooFactory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 3*time.Second, balancer)

	// And now retry can be used like any other endpoint.
	req := struct{}{}
	if _, err = retry(ctx, req); err != nil {
		panic(err)
	}
	return nil
}

func fooFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println(instanceAddr)
		conn, err := grpc.Dial(instanceAddr,  grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "golang"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		return nil, nil
	}, nil, nil
}