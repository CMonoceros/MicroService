package http

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/common/response"
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/service"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func getGinEngine() (engine *gin.Engine) {
	gin.DefaultWriter = log.GetGinDefaultWriter()
	gin.DefaultErrorWriter = log.GetGinDefaultErrorWriter()

	engine = gin.New()
	engine.Use(log.GetGinFormatter())
	engine.Use(func(c *gin.Context) {
		defer response.CatchPanicFunc(c)()
		c.Next()
	})
	return
}

func New(c *conf.Config, s *service.Service) (server *Server) {
	svc = s

	server = new(Server)
	server.engine = getGinEngine()
	server.server = &http.Server{Handler: server.engine}

	initRouter(server.engine)

	listener, err := net.Listen("tcp", c.HTTP.Addr)
	if err != nil {
		log.Error("HTTP New server error=%v", err)
	}

	s.Go(func() error {
		return server.server.Serve(listener)
	})

	return server
}

func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)

	setGroup := e.Group("/set")
	setGroup.GET("/list", listSets)
}

func ping(ctx *gin.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}

	response.JSON(ctx, nil, nil)
}
