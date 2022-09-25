package service

import (
	"context"
	"fmt"
	pb "github.com/miiy/go-examples/trace/api/order"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceServer() pb.OrderServiceServer {
	return &OrderServiceServer{}
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.Order, error) {
	//tracer := otel.Tracer("order-service")
	//_, span := tracer.Start(ctx, "service/GetOrder")
	//defer span.End()
	span := trace.SpanFromContext(ctx)

	fmt.Printf("TraceId: %s\n", span.SpanContext().TraceID())

	name := request.GetName()
	span.SetAttributes(attribute.String("Name", name))

	return &pb.Order{
		Id:    0,
		Title: "computer",
	}, nil
}
