package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goFrame/app/Http/Middleware"
	"goFrame/app/Tools"
	"goFrame/routes"
	"os"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (a Api) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "api模块",
		Long:  `api模块的启动监听`,
		Run: func(cmd *cobra.Command, args []string) {
			a.run()
		},
	}
}

func (a *Api) run() {

	//环境
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if os.Getenv("APP_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	//框架初始化
	r := gin.Default()
	//载入中间件
	r.NoMethod(Tools.HandleNotFound)
	r.NoRoute(Tools.HandleNotFound)
	//全局异常中间件
	r.Use(Middleware.ErrHandler())

	//载入路由
	routes.LoadApi(r)

	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Api服务启动失败:%v\n", err)
	}

	fmt.Println("Api服务启动成功")
}
