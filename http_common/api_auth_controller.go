package http_common

import (
	"futong_server_agent_go/utils/logger"
	"futong_server_agent_go/utils/token"
	"net/http"
)

type ApiAuthController struct {
	ApiController
}

func (c *ApiAuthController) Prepare() {
	c.ApiController.Prepare()
	c.checkToken()
}

func (c *ApiAuthController) checkToken() {
	if err := token.DefaultCredential.Verify(c.Ctx.Request); err != nil {
		logger.Sugar.Error(err)
		c.ApiError(http.StatusUnauthorized, "Unauthorized")
	}
}
