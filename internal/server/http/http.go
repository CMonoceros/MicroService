package http

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/model"
	"SnowBrick-Backend/internal/service"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func New(c *conf.Config, s *service.Service) (server *Server) {
	svc = s

	server = new(Server)
	gin.DefaultWriter = log.GetDefaultWriter()
	gin.DefaultErrorWriter = log.GetDefaultErrorWriter()
	server.engine = gin.New()
	server.engine.Use(log.GetGinFormatter())
	server.server = &http.Server{Handler: server.engine}

	initRouter(server.engine)

	listener, err := net.Listen("tcp", c.Http.Addr)
	if err != nil {
		log.Error("", err)
	}
	s.Go(func() error {
		return server.server.Serve(listener)
	})

	return server
}

func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
	g := e.Group("/SnowBrick-Backend")
	g.GET("/start", start)
}

func ping(ctx *gin.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func start(c *gin.Context) {
	k := &model.Kratos{
		Content: "Golang 大法好 !!!",
	}
	log.Info("start content(%v)", k.Content)
	c.JSON(200, k)
}
