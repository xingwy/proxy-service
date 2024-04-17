package main

import (
	"flag"
	"log"
	"proxy-service/api/handler"
	"proxy-service/api/http"
	"proxy-service/conf"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
)

type program struct {
	handler    *handler.Handler
	httpEngine *gin.Engine
}

func (p *program) Start(s service.Service) error {
	// 初始化处理模块
	p.handler = handler.New(conf.Conf)
	// 注册路由
	p.httpEngine = http.Reginster(gin.Default())
	// 运行
	err := p.httpEngine.Run(conf.Conf.GinServer.Addr())
	if err != nil {
		return err
	}
	return nil
}

func (p *program) Stop(s service.Service) error {
	// TODO
	return nil
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	svcConfig := &service.Config{
		Name:        conf.Conf.ServiceName,
		DisplayName: "数据中心",
		Description: "商品中台-数据中心",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
