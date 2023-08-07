package handler

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	"github.com/finexblock-dev/gofinexblock/pkg/goaws"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
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

func (s *Server) Check(ctx context.Context, input *health.HealthCheckInput) (*health.HealthCheckOutput, error) {
	return &health.HealthCheckOutput{Message: fmt.Sprintf("Hello %s", input.Name)}, nil
}

// ProxyHealthCheck @ProxyHealthCheck
//
//	@security		BearerAuth
//	@description	gRPC Health check in same VPC.
//	@tags			gRPC
//	@accept			json
//	@produce		json
//	@param			ProxyHealthCheckInput	query		dto.ProxyHealthCheckInput	true	"ProxyHealthCheckInput"
//	@success		200					{object}	dto.ProxyHealthCheckOutput	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin [get]
func ProxyHealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var client health.HealthCheckClient
		var conn *grpc.ClientConn
		var err error
		var resp *health.HealthCheckOutput
		var query = new(dto.ProxyHealthCheckInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		conn, err = grpc.Dial(fmt.Sprintf("%s:50051", query.Domain), grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: false,
		})))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		defer func(conn *grpc.ClientConn) {
			_ = conn.Close()
		}(conn)

		client = health.NewHealthCheckClient(conn)

		resp, err = client.Check(context.Background(), &health.HealthCheckInput{Name: "Proxy"})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(resp.Message)
	}
}

func NewServer() *Server {
	return &Server{
		UnimplementedHealthCheckServer: health.UnimplementedHealthCheckServer{},
	}
}

func (s *Server) mustEmbedUnimplementedHealthCheckServer() {

}