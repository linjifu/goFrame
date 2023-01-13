package Tools

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validation(c *gin.Context, v any) any {
	if err := c.ShouldBind(v); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			panic(new(ExceptionHandle).ValidationException(err.Error()))
		}
		for _, v := range errs.Translate(Trans) {
			panic(new(ExceptionHandle).ValidationException(v))
		}
	}
	return v
}
