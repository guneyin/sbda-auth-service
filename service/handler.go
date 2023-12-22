package service

import (
	"context"
	"github.com/guneyin/sbda-auth-service/config"
	"github.com/guneyin/sbda-auth-service/usecase"
	pb "github.com/guneyin/sbda-sdk/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	config *config.Config
	pb.UnimplementedAuthServiceServer
}

func newGrpcHandler(cfg *config.Config) *grpcHandler {
	return &grpcHandler{
		config: cfg,
	}
}

func (h *grpcHandler) InitAuth(ctx context.Context, in *pb.InitAuthRequest) (*pb.InitAuthResponse, error) {
	r, err := usecase.InitAuth(ctx, in.CallbackUrl, h.config.GoogleOauthClientId, h.config.GoogleOauthClientSecret)
	if err != nil {
		return nil, err
	}

	return &pb.InitAuthResponse{
		Url: r.Url,
	}, nil
}

func (h *grpcHandler) Callback(ctx context.Context, in *pb.CallbackRequest) (*pb.CallbackResponse, error) {
	res, err := usecase.Callback(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	return &pb.CallbackResponse{
		Id:      res.Id,
		Email:   res.Email,
		Picture: res.Picture,
		Token: &pb.CallbackToken{
			AccessToken:  res.Token.AccessToken,
			RefreshToken: res.Token.RefreshToken,
			Expiry:       res.Token.Expiry,
		},
	}, nil
}

func (h *grpcHandler) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h *grpcHandler) Watch(in *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
