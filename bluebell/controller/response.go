/*
	通过结构体定义response返回的数据
*/

type ResponseData struct {
	Code int `json:"code"`
	MSg interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

//定义返回函数
func ResponseError(c *gin.Context){}