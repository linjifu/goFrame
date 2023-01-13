package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goFrame/app/Http/Middleware"
	"goFrame/app/Tools"
	"goFrame/routes"
)

func main() {
	//框架初始化
	r := gin.Default()
	//载入中间件
	r.NoMethod(Tools.HandleNotFound)
	r.NoRoute(Tools.HandleNotFound)
	//全局异常中间件
	r.Use(Middleware.ErrHandler())

	//载入路由
	routes.LoadApi(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
