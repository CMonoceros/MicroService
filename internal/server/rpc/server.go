package rpc

import (
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/service"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(c *conf.Config, svc *service.Service) *warden.Server {
	ws := warden.NewServer(c.Warden)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
