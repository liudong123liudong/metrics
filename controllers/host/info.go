package host

import (
	"futong_server_agent_go/http_common"
	"futong_server_agent_go/hub/collect"
)

type HostInfoController struct {
	http_common.ApiController
}

func (c *HostInfoController) Get() {
	c.ApiSuccess(map[string]interface{}{
		"hostInfo": collect.CollectHostInfo(),
	})
}
