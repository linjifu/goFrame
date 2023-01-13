package routes

import (
	"github.com/gin-gonic/gin"
	"goFrame/app/Http/Controllers/Api"
)

func LoadApi(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		v1.GET("/index", new(Api.EventController).Index)
		////足球模块
		//v1Zu := v1.Group("/zu")
		//{
		//	v1Zu.POST("/odds/get-3004-all", new(zu.OddsController).Get3004All)
		//}
	}
}
