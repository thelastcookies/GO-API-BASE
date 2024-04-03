package errno

import (
	"fmt"
	"net/http"
	"sync"
)

// Error 返回错误码和消息的结构体
// nolint: govet
type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var errorCodes = map[int]struct{}{}
var toStatus sync.Map

// NewError create a error
func NewError(code int, msg string) *Error {
	if _, ok := errorCodes[code]; ok {
		panic(fmt.Sprintf("code %d is exsit, please change one", code))
	}
	errorCodes[code] = struct{}{}
	return &Error{Code: code, Msg: msg}
}

// Error return a error string
func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.GetCode(), e.GetMsg())
}

// GetCode return error code
func (e *Error) GetCode() int {
	return e.Code
}

// GetMsg return error msg
func (e *Error) GetMsg() string {
	return e.Msg
}

// GetMsgF format error string
func (e *Error) GetMsgF(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

// GetDetails return more error details
func (e *Error) GetDetails() []string {
	return e.Details
}

// WithDetails return err with detail
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	newError.Details = append(details)

	return &newError
}

// SetHTTPStatusCode set a specific http status code to err
func SetHTTPStatusCode(err *Error, status int) {
	toStatus.Store(err.GetCode(), status)
}

// ToHTTPStatusCode convert custom error code to http status code and avoid return unknown status code.
func ToHTTPStatusCode(code int) int {
	if status, ok := toStatus.Load(code); ok {
		return status.(int)
	}

	return http.StatusBadRequest
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

// Error return error string
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.Code, Success.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Error:
		return typed.Code, typed.Msg
	default:
	}

	return ErrInternalServer.GetCode(), err.Error()
}

func initToStatus() {
	for code, status := range map[int]int{
		Success.GetCode():               http.StatusOK,
		ErrInternalServer.GetCode():     http.StatusInternalServerError,
		ErrNotFound.GetCode():           http.StatusNotFound,
		ErrInvalidParam.GetCode():       http.StatusBadRequest,
		ErrToken.GetCode():              http.StatusUnauthorized,
		ErrInvalidToken.GetCode():       http.StatusUnauthorized,
		ErrTokenTimeout.GetCode():       http.StatusUnauthorized,
		ErrTooManyRequests.GetCode():    http.StatusTooManyRequests,
		ErrServiceUnavailable.GetCode(): http.StatusServiceUnavailable,
	} {
		toStatus.Store(code, status)
	}
}

func init() {
	initToStatus()
}
