package http_common

import "encoding/json"

type ApiController struct {
	BaseController
	BodyData []byte
}

func (c *ApiController) Prepare() {
	c.BaseController.Prepare()
	c.CopyBody()

	// 解析body中json数据
	c.formJsonBody()
}

// 获取 Request Body 里的 JSON 或 XML 的数据，c.Ctx.Input.RequestBody, 这样可以避免必须在配置文件conf中配置 copyrequestbody = true
// 也可以直接通过c.Ctx.Request.Body 获取
func (c *ApiController) CopyBody() {
	c.BodyData = c.Ctx.Input.CopyBody(2048)
}

// 解析body中json数据
func (c *ApiController) formJsonBody() {
	jsonData := make(map[string]interface{})

	_ = json.Unmarshal(c.BodyData, &jsonData)

	for k, v := range jsonData {
		if str, ok := v.(string); ok {
			c.Ctx.Input.SetParam(k, str)
		}
	}
}

func (c *ApiController) WriteString(s string) {
	c.responseTrace(s)
	c.Ctx.WriteString(s)
}

func (c *ApiController) ApiResponse(code int, msg string,  data map[string]interface{}) {
	if code != 0 {
		c.Ctx.ResponseWriter.WriteHeader(code)
	}

	result := make(map[string]interface{})
	result["code"] = code
	result["msg"] = msg
	result["data"] = data
	c.Data["json"] = result
	//c.SetData(result)
	c.ServeJSON()

	// 出参跟踪
	c.responseTrace(result)

	// 结束请求
	c.StopRun()
}

//重写父类方法， 父类方法默认在非prod模式会进行json格式化操作
//这里设置为不格式化
func (c *ApiController) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = false
		hasEncoding = len(encoding) > 0 && encoding[0]
	)

	c.Ctx.Output.JSON(c.Data["json"], hasIndent, hasEncoding)
}

func (c *ApiController) ApiSuccess(data map[string]interface{}) {
	c.ApiResponse(0, "success", data)
}

func (c *ApiController) ApiError(code int, msg string) {
	c.ApiResponse(code, msg, make(map[string]interface{}))
}


