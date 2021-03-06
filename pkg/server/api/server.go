package api

import (
	"context"
	"fmt"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"go.uber.org/zap"

	api_pb "github.com/srvc/glx/api"
	"github.com/srvc/glx/pkg/glx/domain"
	"github.com/srvc/glx/pkg/server/api/internal"
	"github.com/srvc/glx/pkg/server/proxy"
)

var (
	defaultNetwork = "tcp"
	defaultPort    = 80
)

type Server struct {
	appRepo domain.AppRepository
	proxies proxy.Manager

	server *grapiserver.Engine
	log    *zap.Logger
}

func NewServer(
	appRepo domain.AppRepository,
	proxies proxy.Manager,
) *Server {
	return &Server{
		appRepo: appRepo,
		proxies: proxies,
		log:     zap.L().Named("api"),
	}
}

func (s *Server) Serve(ctx context.Context) error {
	app := &api_pb.App{
		Name:     "srvc.tools/glx/api",
		Hostname: "api.glx.local",
		Type:     api_pb.App_TYPE_LOCAL,
		Ip:       "127.0.0.1",
		Ports: []*api_pb.App_Port{
			{
				Network:      api_pb.App_Port_TCP,
				ExposedPort:  80,
				InternalPort: 80,
			},
		},
	}
	err := s.appRepo.Create(ctx, app)
	if err != nil {
		return nil
	}
	s.server = grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithSignalHandling(false),
		grapiserver.WithAddr(defaultNetwork, fmt.Sprintf("%s:%d", app.GetIp(), defaultPort)),
		grapiserver.WithServers(
			internal.NewAppServiceServer(s.appRepo, s.proxies),
		),
	)
	return s.server.ServeContext(ctx)
}
