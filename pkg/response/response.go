package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/utils"
)

const (
	MessageOK                  = "OK"
	MessageBadRequest          = "未提供合法的请求参数"
	MessageNotFound            = "未找到匹配的记录"
	MessageInternalServerError = ""
)

var Res = NewResponse()

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
		Code:    errno.Success.Code(),
		Message: errno.Success.Msg(),
		Data:    data,
	})
}

func (r *Response) Error(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code:    errno.Success.Code(),
			Message: errno.Success.Msg(),
			Data:    gin.H{},
		})
		return
	}

	switch typed := err.(type) {
	case *errno.Error:
		response := Response{
			Code:    typed.Code(),
			Message: typed.Msg(),
			Data:    gin.H{},
			Details: []string{},
		}
		details := typed.Details()
		if len(details) > 0 {
			response.Details = details
		}
		c.JSON(errno.ToHTTPStatusCode(typed.Code()), response)
		return
	}
}

//func (r *Response) Error(c *gin.Context, err error) {
//	if err == nil {
//		c.JSON(http.StatusOK, Response{
//			Code: http.StatusOK,,
//			Message: MessageOK,
//			Data:    gin.H{},
//		})
//		return
//	}
//	switch typed := err.(type) {
//	case *errno.Error:
//		response := Response{
//			Code:    typed.Code(),
//			Message: typed.Msg(),
//			Data:    gin.H{},
//			Details: []string{},
//		}
//		details := typed.Details()
//		if len(details) > 0 {
//			response.Details = details
//		}
//		c.JSON(errno.ToHTTPStatusCode(typed.Code()), response)
//		return
//	}
//}

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
