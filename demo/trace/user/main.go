package main

import (
	"flag"
	"fmt"
	pb "github.com/miiy/go-examples/trace/api/user"
	"github.com/miiy/go-examples/trace/common/pkg/tracing"
	"github.com/miiy/go-examples/trace/user/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tp, err := tracing.NewTracerProvider(&tracing.TraceConfig{
		Endpoint:    "http://localhost:14268/api/traces",
		ServiceName: "user-service",
		Environment: "production",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3.New())

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	pb.RegisterUserServiceServer(s, service.NewUserServiceServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
