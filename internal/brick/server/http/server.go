package http

import (
	"SnowBrick-Backend/internal/brick/service"
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
