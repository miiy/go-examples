package main

import (
	"context"
	"flag"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	pb "github.com/miiy/go-examples/trace/api/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = "zhang san"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)

	if err = callGetOrder(c); err != nil {
		log.Fatal(err)
	}

}

func callGetOrder(c pb.OrderServiceClient) error {
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id",
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.GetOrder(ctx, &pb.GetOrderRequest{Name: name})
	if err != nil {
		log.Fatalf("calling GetOrder: %s", err)
	}
	log.Printf("Response from GetOrder: %s", r.GetTitle())
	return nil
}
