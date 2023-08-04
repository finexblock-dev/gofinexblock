package handler

import (
	"context"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	"github.com/finexblock-dev/gofinexblock/pkg/goaws"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	health.UnimplementedHealthCheckServer
}

func (s *Server) WhoAmI(ctx context.Context, input *health.WhoAmIInput) (*health.WhoAmIOutput, error) {
	privateIP, err := goaws.OwnPrivateIP()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	return &health.WhoAmIOutput{Message: fmt.Sprintf("Hello %s, I am %s", input.Name, privateIP)}, nil
}

func NewServer() *Server {
	return &Server{
		UnimplementedHealthCheckServer: health.UnimplementedHealthCheckServer{},
	}
}

func (s *Server) Check(ctx context.Context, input *health.HealthCheckInput) (*health.HealthCheckOutput, error) {
	return &health.HealthCheckOutput{Message: fmt.Sprintf("Hello %s", input.Name)}, nil
}

func (s *Server) mustEmbedUnimplementedHealthCheckServer() {

}