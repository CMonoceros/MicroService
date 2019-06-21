package rpc

import (
	"context"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	ch := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(ch)
	}()
	select {
	case <-ctx.Done():
		s.server.Stop()
		err = ctx.Err()
	case <-ch:
	}
	return
}
