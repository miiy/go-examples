package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/miiy/go-examples/trace/api/user"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"time"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer() pb.UserServiceServer {
	return &UserServiceServer{}
}

func (s *UserServiceServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	span := trace.SpanFromContext(ctx)
	fmt.Printf("TraceId: %s\n", span.SpanContext().TraceID())

	name := request.GetName()
	span.SetAttributes(attribute.String("Name", name))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not get metadata from incoming stream context")
	}

	fmt.Printf("MD: %+v\n", md)

	if err := validateGetUserRequest(ctx, request); err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   0,
		Name: name,
	}, nil
}

func validateGetUserRequest(ctx context.Context, request *pb.GetUserRequest) error {
	//tracer := otel.Tracer("user-service")
	//_, span := tracer.Start(ctx, "server/validateGetUserRequest")
	//defer span.End()
	span := trace.SpanFromContext(ctx)
	span.AddEvent("validateGetUserRequest")

	time.Sleep(50 * time.Millisecond)

	return nil
}
