package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	orderpb "github.com/miiy/go-examples/trace/api/order"
	userpb "github.com/miiy/go-examples/trace/api/user"
	"github.com/miiy/go-examples/trace/common/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

var userClient userpb.UserServiceClient
var orderClient orderpb.OrderServiceClient

func connGrpc() func() {
	// Set up a connection to the server.
	connUser, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	userClient = userpb.NewUserServiceClient(connUser)

	// Set up a connection to the server.
	connOrder, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	orderClient = orderpb.NewOrderServiceClient(connOrder)

	return func() {
		connUser.Close()
		connOrder.Close()
	}

}

func main() {

	tp, err := tracing.NewTracerProvider(&tracing.TraceConfig{
		Endpoint:    "http://localhost:14268/api/traces",
		ServiceName: "gw-service",
		Environment: "production",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3.New())

	//
	cleanup := connGrpc()
	defer cleanup()

	r := gin.Default()
	r.Use(otelgin.Middleware("gw-server"))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user", userHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func userHandler(ctx *gin.Context) {
	// Use the global TracerProvider.
	//tr := otel.Tracer("gateway")
	span := trace.SpanFromContext(ctx.Request.Context())
	//_, span := tr.Start(ctx, "gateway-userHandler")
	//defer span.End()

	name := "zhang san"
	span.SetAttributes(attribute.String("Name", name))

	fmt.Println(span.SpanContext().TraceID())

	// user
	userReq := userpb.GetUserRequest{Name: name}
	userMd := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "gateway-1",
		"user-name", name,
	)
	uCtx := metadata.NewOutgoingContext(ctx.Request.Context(), userMd)
	resp, err := userClient.GetUser(uCtx, &userReq)
	if err != nil {
		log.Fatalf("user-service.GetUser(_) = _, %v", err)
	}
	fmt.Println(resp)

	// order
	orderReq := orderpb.GetOrderRequest{Name: name}
	orderResp, err := orderClient.GetOrder(uCtx, &orderReq)
	if err != nil {
		log.Fatalf("order-service.GetOrder(_) = _, %v", err)
	}
	fmt.Println(orderResp)

}
