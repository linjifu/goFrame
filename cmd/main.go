package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"goFrame/app/Tools"
	"goFrame/cmd/Api"
	"os"
	"strings"
)

func init() {
	// 环境变量载入
	loadEnv()
	// MySql载入
	Tools.NewDbs()
	// Redis载入
	Tools.NewRedises()
}

func main() {
	// 命令行根，用于后续新增新的命令行
	rootCmd := &cobra.Command{
		Use:   "Hello World",
		Short: "您好，世界！",
		Long:  `启动程序默认命令`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
		},
	}
	// 增加Api模块
	rootCmd.AddCommand(Api.NewApi().Cmd())
	// 启动程序
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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
