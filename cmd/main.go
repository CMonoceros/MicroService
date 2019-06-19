package main

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/server/rpc"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"SnowBrick-Backend/internal/server/http"
	"SnowBrick-Backend/internal/service"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}

	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("SnowBrick-Backend start")

	// init service
	svc := service.New(conf.Conf)
	// init rpc
	rpcSrv := rpc.New(conf.Conf, svc)
	// init http
	httpSrv := http.New(conf.Conf, svc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Conf.TimeoutSecond)*time.Second)
			if err := rpcSrv.Shutdown(ctx); err != nil {
				log.Error("grpcSrv.Shutdown error(%v)", err)
			}
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}

			log.Info("SnowBrick-Backend exit")
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
