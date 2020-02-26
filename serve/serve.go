package serve

import (
	"futong_server_agent_go/controllers/error_handler"
	"futong_server_agent_go/hub"
	_ "futong_server_agent_go/routers"
	"futong_server_agent_go/utils/conf"
	"futong_server_agent_go/utils/logger"
	"github.com/astaxie/beego"
	"github.com/kardianos/service"
	"os"
)

var Service = NewServe()

type Serve struct{}

func NewServe() *Serve {
	return &Serve{}
}

func (s *Serve) RunServe() {
	// 开启hub服务
	go hub.MyHub.Run()

	// 错误处理
	beego.ErrorController(&error_handler.ErrorController{})

	// 启动http服务
	port := conf.Config.MustValue("", "httpport", "17086")
	addr := ":" + port
	beego.Run(addr)
}

func (s *Serve) Start(srv service.Service) error {
	go s.RunServe()
	return nil
}

func (s *Serve) Stop(srv service.Service) error {
	os.Exit(0)
	return nil
}

func (s *Serve) StartServe() {
	daemonCfg := &service.Config{
		Name: "futongAgent",
	}

	daemonService, err := service.New(Service, daemonCfg)
	if err != nil {
		logger.Sugar.Error("new daemon service err:", err)
	}

	if len(os.Args) > 1 {
		err := service.Control(daemonService, os.Args[1])
		if err != nil {
			logger.Sugar.Error(err)
			logger.Sugar.Infof("valid signals: %q\n", service.ControlAction)
		}
		return
	}

	err = daemonService.Run()
	if err != nil {
		logger.Sugar.Error("run daemon service err:", err)
	}
}
