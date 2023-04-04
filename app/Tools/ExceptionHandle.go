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

	s := "%s %s \"%s %s\" " +
		"%s %d %s " +
		"\"%s\""

	layout := "2006-01-02 15:04:05"
	timeNow := time.Now().Format(layout)

	Log.ErrorLog.Errorf(s,
		GetRealIp(c),
		timeNow,
		c.Request.Method,
		c.Request.RequestURI,
		c.Request.Proto,
		404,
		"找不到路由或者方法",
		c.Request.Header.Get("User-Agent"),
	)

	defer Log.ErrorLog.Sync()

	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}
