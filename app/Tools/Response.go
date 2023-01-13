package Tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	REQUEST_SUCCESS = 200
	REQUEST_ERROR   = 500
	PARAMETER_ERROR = 400
	USER_AUTH_ERROR = 401
	CSRF_ERROR      = 419
)

type Response struct {
}

func (r *Response) ReturnJsonSuccess(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"result": data,
	})
}

func (r *Response) ReturnJsonError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
