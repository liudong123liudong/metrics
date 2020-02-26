package host

import (
	"futong_server_agent_go/http_common"
	"futong_server_agent_go/hub/collect"
)

type HostMetricsController struct {
	http_common.ApiController
}

// 获取主机监控数据
func (c *HostMetricsController) Get() {
	c.ApiSuccess(map[string]interface{}{
		"hostMetrics": collect.CollectHostMetrics(),
	})
}



