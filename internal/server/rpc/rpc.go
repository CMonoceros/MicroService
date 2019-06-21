package rpc

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/service"
	"google.golang.org/grpc"
	"net"
)

func NewServer(c *conf.Config, svc *service.Service) *Server {
	listener, err := net.Listen("tcp", c.Grpc.Addr)
	if err != nil {
		log.Error("", err)
	}
	s := new(Server)
	s.server = grpc.NewServer()

	svc.Go(func() error {
		return s.server.Serve(listener)
	})
	return s
}
