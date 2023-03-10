package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goFrame/app/Http/Middleware"
	"goFrame/app/Tools"
	"goFrame/routes"
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
		fmt.Printf("Api服务启动失败:%v\n", err)
	}
}
