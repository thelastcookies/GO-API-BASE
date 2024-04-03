package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/utils"
)

const (
	MessageOK                  = "OK"
	MessageBadRequest          = "未提供合法的请求参数"
	MessageNotFound            = "未找到匹配的记录"
	MessageInternalServerError = ""
)

// Response define a response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Details []string    `json:"details,omitempty"`
}

// NewResponse return a response
func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Send(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    errno.Success.GetCode(),
		Message: errno.Success.GetMsg(),
		Data:    data,
	})
}

func (r *Response) Error(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code:    errno.Success.GetCode(),
			Message: errno.Success.GetMsg(),
			Data:    gin.H{},
		})
		return
	}

	switch typed := err.(type) {
	case *errno.Error:
		response := Response{
			Code:    typed.GetCode(),
			Message: typed.GetMsg(),
			Data:    gin.H{},
			Details: []string{},
		}
		details := typed.GetDetails()
		if len(details) > 0 {
			response.Details = details
		}
		c.JSON(errno.ToHTTPStatusCode(typed.GetCode()), response)
		return
	}
}

// Send return a response
func Send(c *gin.Context, code int, msg string, data interface{}) {
	resp := NewResponse()
	resp.Send(c, code, msg, data)
}

// Success return a success response
func Success(c *gin.Context, data interface{}) {
	resp := NewResponse()
	resp.Success(c, data)
}

// Error return a error response
func Error(c *gin.Context, err error) {
	resp := NewResponse()
	resp.Error(c, err)
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "The route not found, please check your request url.")
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the connection is healthy.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}
