package routers

import (
	"futong_server_agent_go/controllers/file"
	"futong_server_agent_go/controllers/host"
	"futong_server_agent_go/controllers/version"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("api/v1/yw",
			// 获取主机信息
			beego.NSRouter("/host", &host.HostInfoController{}),
			// 获取采集指标
			beego.NSRouter("/metrics", &host.HostMetricsController{}),
			// 脚本上传
			beego.NSRouter("/file", &file.FileController{}),
			// 版本下载
			beego.NSRouter("/download", &version.VersionController{}),
		)
	beego.AddNamespace(ns)
}
