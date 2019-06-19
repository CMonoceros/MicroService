package http

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/conf"
	"net/http"

	"SnowBrick-Backend/internal/model"
	"SnowBrick-Backend/internal/service"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine) {
	svc = s
	engine = bm.DefaultServer(c.BM)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/SnowBrick-Backend")
	g.GET("/start", start)
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func start(c *bm.Context) {
	k := &model.Kratos{
		Content: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
