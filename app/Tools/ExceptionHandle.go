package Tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ExceptionHandle struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func NewError(statusCode, Code int, msg string) *ExceptionHandle {
	return &ExceptionHandle{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

func (e *ExceptionHandle) Error() string {
	return e.Msg
}

// 参数错误异常
func (e *ExceptionHandle) ValidationException(message string) *ExceptionHandle {
	return NewError(http.StatusOK, PARAMETER_ERROR, message)
}

// csrf验证异常
func (e *ExceptionHandle) TokenMismatchException(message string) *ExceptionHandle {
	return NewError(http.StatusOK, CSRF_ERROR, message)
}

// 用户认证异常
func (e *ExceptionHandle) AuthenticationException(message string) *ExceptionHandle {
	return NewError(http.StatusOK, USER_AUTH_ERROR, message)
}

// 404异常
func (e *ExceptionHandle) NotFoundException(message string) *ExceptionHandle {
	return NewError(http.StatusOK, 404, http.StatusText(http.StatusNotFound))
}

// 其他异常
func (e *ExceptionHandle) OtherError(message string) *ExceptionHandle {
	return NewError(http.StatusOK, REQUEST_ERROR, message)
}

var (
	NotFound = NewError(http.StatusNotFound, 404, http.StatusText(http.StatusNotFound))
)

// 404处理
func HandleNotFound(c *gin.Context) {

	beginTime := time.Now().UnixNano()
	c.Next()
	endTime := time.Now().UnixNano()
	duration := endTime - beginTime

	s := "%s %s \"%s %s\" " +
		"%s %d %d %dµs " +
		"\"%s\""

	layout := "2006-01-02 15:04:05"
	timeNow := time.Now().Format(layout)

	Tools.Log.AccessLog.Infof(s,
		Tools.GetRealIp(c),
		timeNow,
		c.Request.Method,
		c.Request.RequestURI,
		c.Request.Proto,
		bodyWriter.Status(),
		bodyWriter.body.Len(),
		duration/1000,
		c.Request.Header.Get("User-Agent"),
	)

	defer Tools.Log.AccessLog.Sync()

	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}
