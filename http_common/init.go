package http_common

import (
	"flag"
	"fmt"
	"futong_server_agent_go/utils/global"
	"github.com/astaxie/beego/logs"
)

var (
	TraceLog *logs.BeeLogger
	Addr     string
)

func init() {
	//initCommandLineParams()
	initTraceLog()
}

// 初始化追踪日志
func initTraceLog() {
	fmt.Println("init trace.log......")
	if TraceLog != nil {
		return
	}
	// 初始化日志 和系统日志区分开
	TraceLog = logs.NewLogger()
	err := TraceLog.SetLogger(logs.AdapterFile, `{"filename":"` + global.Global.Path + `/logs/trace.log","perm":"666"}`)
	if err != nil {
		logs.Error("init trace.log err:", err)
	}
}

// 解析命令行参数，默认为"127.0.0.1:17086"
func initCommandLineParams() {
	host := flag.String("h", "", "host")
	port := flag.String("p", "17086", "port")
	flag.Parse()
	Addr = *host + ":" + *port
}
