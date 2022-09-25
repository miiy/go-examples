package main

import (
	"context"
	"flag"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	pb "github.com/miiy/go-examples/trace/api/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = "zhang san"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()), //
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	if err = callGetUser(c); err != nil {
		log.Fatal(err)
	}

}

func callGetUser(c pb.UserServiceClient) error {
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id",
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.GetUser(ctx, &pb.GetUserRequest{Name: name})
	if err != nil {
		log.Fatalf("calling GetUser: %s", err)
	}
	log.Printf("Response from GetUser: %s", r.GetName())
	return nil
}
