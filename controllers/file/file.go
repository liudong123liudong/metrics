package file

import (
	"fmt"
	"futong_server_agent_go/http_common"
	"futong_server_agent_go/hub"
	"futong_server_agent_go/utils"
	"futong_server_agent_go/utils/logger"
	"github.com/astaxie/beego/validation"
	"net/http"
	"os"
)

type FileController struct {
	http_common.ApiController
}

// 上传脚本并执行
func (c *FileController) Post() {
	// 接受
	checkNum := c.GetString("check_num")

	// 验证
	valid := validation.Validation{}
	valid.Required(checkNum, "check_num")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.Sugar.Error(err.Key, err.Message)
		}
		c.ApiError(http.StatusBadRequest, "check_num can not empty")
	}

	// 接受文件
	f, h, err := c.GetFile("file")
	if err != nil {
		logger.Sugar.Error(err)
		c.ApiError(http.StatusBadRequest, "no such file")
	}
	defer f.Close()

	// 完整性校验
	strSha1, err := utils.GetFileSha1(f)
	if err != nil {
		logger.Sugar.Error(err)
		c.ApiError(http.StatusInternalServerError, "file sha1 error")
	}

	if strSha1 != checkNum {
		logger.Sugar.Error("checkNum err not equal")
		//c.ApiError(http.StatusBadRequest, "file check_num err")
	}

	// 创建文件夹
	if _, err := os.Stat("upload"); err != nil {
		if err := os.Mkdir("upload", os.ModePerm); err != nil {
			logger.Sugar.Error(err)
			c.ApiError(http.StatusInternalServerError, "")
		}
	}

	// 保存文件
	filePos := fmt.Sprintf("upload/%s", h.Filename)
	err = c.SaveToFile("file", filePos)
	if err != nil {
		logger.Sugar.Error("save to file err:", err)
		c.ApiError(http.StatusInternalServerError, "")
	}

	// 执行脚本
	go hub.MyHub.ScriptTask.Exec(checkNum, h.Filename)

	// 返回数据
	c.ApiSuccess(map[string]interface{}{})
}

// 获取脚本执行结果
func (c *FileController) Get() {
	// 接受
	checkNum := c.GetString("check_num")

	valid := validation.Validation{}
	valid.Required(checkNum, "check_num")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.Sugar.Error(err.Key, err.Message)
		}
		c.ApiError(http.StatusBadRequest, "params error")
	}

	data := hub.MyHub.ScriptTask.Cache[checkNum]

	// 删除缓存
	delete(hub.MyHub.ScriptTask.Cache, checkNum)

	// 返回
	c.ApiSuccess(map[string]interface{}{
		"result": data,
	})
}
