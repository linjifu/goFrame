package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"goFrame/app/Tools"
	"goFrame/cmd/Api"
	"goFrame/cmd/Artisan"
	"goFrame/cmd/Queue"
	"goFrame/cmd/Schedule"
	"os"
	"strings"
)

// 环境变量载入
func loadEnv() {
	path, _ := os.Getwd()
	_, envExistErr := os.Stat(path + "/.env")
	if envExistErr == nil && !os.IsNotExist(envExistErr) {
		viperEnv := viper.New()
		viperEnv.SetConfigFile(path + "/.env")
		envErr := viperEnv.ReadInConfig()
		if envErr != nil {
			fmt.Println()
			panic(fmt.Errorf("读取配置文件.env文件失败:%s \n", envErr))
		}
		for i, k := range viperEnv.AllSettings() {
			fmt.Println(i, "-", k)
			os.Setenv(strings.ToUpper(i), k.(string))
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   os.Getenv("APP_NAME"),
	Short: "主程序",
	Long:  `主程序` + os.Getenv("APP_NAME") + `启动入口`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("主程序：", os.Getenv("APP_NAME"), "启动成功！")
	},
}

// addModelCmd 统一新增项目模块的启动命令
func addModelCmd() {
	// 增加Api模块
	rootCmd.AddCommand(Api.NewApi().Cmd())
	// 增加定时计划器模块
	rootCmd.AddCommand(Schedule.NewSchedule().Cmd())
	// 自定义artisan命令模块
	rootCmd.AddCommand(Artisan.NewArtisan().Cmd())
	// 队列模块
	rootCmd.AddCommand(Queue.NewQueue().Cmd())
}

func init() {
	// 环境变量载入
	loadEnv()
	// MySql载入
	Tools.NewDbs()
	// Redis载入
	Tools.NewRedises()
}

func main() {

	a := new(Tools.AsyncQueue)
	a.Serialize(func() error {
		return nil
	}, nil)

	//增加项目模块
	addModelCmd()
	// 启动程序
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
