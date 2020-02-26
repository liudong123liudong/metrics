package version

import (
	"fmt"
	"futong_server_agent_go/http_common"
	"futong_server_agent_go/utils/logger"
	"net/http"
	"runtime"
)

type VersionController struct {
	http_common.ApiController
}

// 版本更新
func (c *VersionController) Get() {
	versionNum := c.GetString("version_num", "100")

	if versionNum == "" {
		logger.Sugar.Errorf("versionNum can not be empty")
		c.ApiError(http.StatusBadRequest, "params error")
	}

	fileName := fmt.Sprintf("download/futong_cm_agent_%s_%s.txt", runtime.GOOS, versionNum)

	c.Ctx.Output.Download(fileName, fileName)
}
