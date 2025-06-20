package main

import (
	"context"
	"fmt"
	klog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"
	"time"
)

// Let's say this is a service that means to register itself.
// First, we will set up some context.
var (
	etcdServer = "127.0.0.1:2379"      // in the change from v2 to v3, the schema is no longer necessary if connecting directly to an etcd v3 instance
	prefix     = "/services/foosvc/"  // known at compile time
	instance   = "127.0.0.1:50051"       // taken from runtime or platform, somehow
	key        = prefix + instance    // should be globally unique
	value      = instance // based on our transport
	ctx        = context.Background()
)

func main() {
	client, err := newClient([]string{etcdServer})
	if err != nil {
		panic(err)
	}

	_, cleanup := register(client)
	defer cleanup()


	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newClient(machines []string) (etcdv3.Client, error) {
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
	client, err := etcdv3.NewClient(ctx, machines, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func register(client etcdv3.Client) (*etcdv3.Registrar, func()) {
	// Build the registrar.
	registrar :=  etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   key,
		Value: value,
	}, klog.NewNopLogger())

	// Register our instance.
	registrar.Register()

	// At the end of our service lifecycle, for example at the end of func main,
	// we should make sure to deregister ourselves. This is important! Don't
	// accidentally skip this step by invoking a log.Fatal or os.Exit in the
	// interim, which bypasses the defer stack.
	cleanup := func() {
		defer registrar.Deregister()
	}

	return registrar, cleanup
}
