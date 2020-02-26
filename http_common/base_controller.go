package http_common

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

// 基础控制器
type BaseController struct {
	beego.Controller
	// 服务器时间，所有请求时间以此时间为准
	ServerTime int64
}

// 前置方法
func (c *BaseController) Prepare() {
	c.ServerTime = time.Now().Unix()
	c.requestTrace()
}

// 跟踪日志
func (c *BaseController) trace(isRequest bool, data interface{}) {
	// 是否开启日志，默认开启
	open := beego.AppConfig.DefaultBool("trace::open", true)
	if !open {
		return
	}

	var LogText = ""

	if isRequest {
		// 解析form和params数据
		c.Ctx.Request.ParseForm()
		m := make(map[string]interface{})
		for k, v := range c.Ctx.Request.Form {
			m[k] = v[0]
		}

		// 格式化form和params数据
		reqParams, err := json.MarshalIndent(m, "", " ")
		if err != nil {
			logs.Info("format form and params fail:", err)
		}

		// 格式化body数据
		jsonBody, err := json.MarshalIndent(c.Ctx.Request.Body, "", " ")
		if err != nil {
			logs.Error("marshal json body err", err)
		}

		LogText += fmt.Sprintf("[%s] [%s] [%s]\n requestParams:\n %s\n requestBody: \n %s",
			c.Ctx.Input.IP(),
			c.Ctx.Request.RequestURI,
			c.Ctx.Request.Method,
			string(reqParams),
			string(jsonBody),
		)
	} else {
		rspData, _ := json.MarshalIndent(data, "", " ")

		LogText += fmt.Sprintf("[%s] [%s] [%s]  code:[%d]\n responseData:\n %s",
			c.Ctx.Input.IP(),
			c.Ctx.Request.RequestURI,
			c.Ctx.Request.Method,
			c.Ctx.ResponseWriter.Status,
			string(rspData),
		)
	}

	// 输出到控制台
	console := beego.AppConfig.DefaultBool("trace::console", true)
	if console {
		logs.Info(LogText)
	}

	// 记录到文件
	TraceLog.Info(LogText)
}

// 入参跟踪
func (c *BaseController) requestTrace() {
	c.trace(true, "")
}

// 出参跟踪
func (c *BaseController) responseTrace(data interface{}) {
	c.trace(false, data)
}
