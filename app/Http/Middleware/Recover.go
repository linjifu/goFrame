package Middleware

import (
	"github.com/gin-gonic/gin"
	"goFrame/app/Tools"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var response = new(Tools.Response)
				var Err *Tools.ExceptionHandle

				if e, ok := err.(*Tools.ExceptionHandle); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = Err.OtherError(e.Error())
				} else {
					Err = Err.OtherError(errorToString(err))
				}

				response.ReturnJsonError(c, Err.Code, Err.Error())
				c.Abort()
			}
		}()
		c.Next()
	}
}

//	func Recover(c *gin.Context) {
//		defer func() {
//			if r := recover(); r != nil {
//
//				//打印错误堆栈信息
//				log.Printf("panic: %v\n", r)
//				debug.PrintStack()
//				//封装通用json返回
//				//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
//				//Result.Fail不是本例的重点，因此用下面代码代替
//				c.JSON(http.StatusOK, gin.H{
//					"code": "1",
//					"msg":  errorToString(r),
//					"data": nil,
//				})
//				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
//				c.Abort()
//			}
//		}()
//		//加载完 defer recover，继续后续接口调用
//		c.Next()
//	}
//
// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
