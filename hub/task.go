package hub

import (
	"encoding/json"
	"futong_server_agent_go/hub/collect"
	"futong_server_agent_go/utils"
	"futong_server_agent_go/utils/conf"
	"futong_server_agent_go/utils/logger"
	"strconv"
	"time"
)

var hostAddr string = conf.Config.MustValue("", "serverAddr")

func sendHostInfo() {
	// TODO 发送主机基本信息到服务端api
	baseInfoUrl := hostAddr + "/api/v1/yw/agent/host/"
	hostInfo, _ := json.Marshal(collect.CollectHostInfo())
	_, err := HttpPost(baseInfoUrl, hostInfo)
	if err != nil {
		logger.Sugar.Error("task send hostInfo err: ", err)
	}
}

func sendHostMetrics() {
	var (
		collectTime int
		ticker      *time.Ticker
	)
	//collectTime = beego.AppConfig.DefaultInt("collect::interval", 15)
	collectTime = conf.Config.MustInt("collect", "interval", 15)

	for {
		ticker = time.NewTicker(1 * time.Second)
		for  range ticker.C {
			// TODO 发送监控数据到服务端api, 服务端每次返回采集频率
			hostMetrics, _ := json.Marshal(collect.CollectHostMetrics())

			postMetricsUrl := hostAddr + "/api/v1/yw/agent/monitor/"

			body, err := HttpPost(postMetricsUrl, hostMetrics)
			if err != nil {
				logger.Sugar.Error("task sendHostMetrics err:", err)
				continue
			}

			interval, err := strconv.Atoi(string(body))
			if interval != collectTime && err == nil && interval >= utils.CollectTimeAllowedMin {
				collectTime = interval
				ticker.Stop()
				logger.Sugar.Infof("采集频率变为:%ds", interval)
				break
			}
		}
	}
}

func sendScriptResult(ret string) {
	// TODO 发送脚本结果到服务端api
	postScriptResultsUrl := "http://127.0.0.1:8899/abc"
	_, err := HttpPost(postScriptResultsUrl, []byte(ret))
	if err != nil {
		logger.Sugar.Infof("task sendScriptResult err:", err)
	}
}

func findScriptTask() {
	// TODO 任务发现
}
