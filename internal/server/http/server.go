package http

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"SnowBrick-Backend/internal/service"
)

var (
	svc *service.Service
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

func (svr *Server) Shutdown(ctx context.Context) error {
	server := svr.server
	if server == nil {
		return errors.New("http: no server")
	}
	return errors.WithStack(server.Shutdown(ctx))
}
