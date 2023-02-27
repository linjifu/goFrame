package Tools

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	path, _ := os.Getwd()
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

	//数据库配置读取==========================================================================================
	viperDatabase := viper.New()
	viperDatabase.SetDefault("charset", "utf8mb4")
	//自动获取全部的env加入到viper中。（如果环境变量多就全部加进来）默认别名和环境变量名一致
	viperDatabase.AutomaticEnv()
	//将加入的环境变量*_*_格式替换成 *.*格式
	//（因为从环境变量读是按"a.b.c.d"的格式读取，所以要给在viper维护一个别名对象，给环境变量一个别名）
	viperDatabase.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//导入配置文件
	viperDatabase.SetConfigType("yaml")
	viperDatabase.SetConfigFile(path + "/config/database.yml")
	//读取配置文件
	err := viperDatabase.ReadInConfig()
	if err != nil {
		fmt.Println()
		panic(fmt.Errorf("读取配置文件database文件失败:%s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viperDatabase.Unmarshal(DBConnection); err != nil {
		fmt.Println()
		panic(fmt.Errorf("配置文件database数据与结构体转换失败:%s \n", err))
	}
	//判断默认连接是否配置
	defaultDBData, ok := DBConnection.Connections[DBConnection.DefaultConnectionName]
	if ok {
		DBConnection.Db = DBConnection.createDB(defaultDBData)
	}

	for key, dbStruct := range DBConnection.Connections {
		if key != DBConnection.DefaultConnectionName {
			DBConnection.Connections[key] = DBConnection.createDB(dbStruct)
		}
	}
	//for i, k := range viperDatabase.AllSettings() {
	//	fmt.Println(i, "-", k)
	//}
	//数据库配置读取结束==========================================================================================

	//redis配置读取开始==========================================================================================

	viperRedis := viper.New()
	//自动获取全部的env加入到viper中。（如果环境变量多就全部加进来）默认别名和环境变量名一致
	viperRedis.AutomaticEnv()
	//将加入的环境变量*_*_格式替换成 *.*格式
	//（因为从环境变量读是按"a.b.c.d"的格式读取，所以要给在viper维护一个别名对象，给环境变量一个别名）
	viperRedis.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//导入配置文件
	viperRedis.SetConfigType("yaml")
	viperRedis.SetConfigFile(path + "/config/redis.yml")
	//读取配置文件
	redisErr := viperRedis.ReadInConfig()
	if redisErr != nil {
		fmt.Println()
		panic(fmt.Errorf("读取配置文件redis文件失败:%s \n", redisErr))
	}
	// 将读取的配置信息保存至全局变量Conf
	if redisErr2 := viperRedis.Unmarshal(RedisConnection); err != nil {
		fmt.Println()
		panic(fmt.Errorf("配置文件redis数据与结构体转换失败:%s \n", redisErr2))
	}
	//判断默认连接是否配置
	defaultRedisData, ok := RedisConnection.Connections[RedisConnection.DefaultConnectionName]
	if ok {
		RedisConnection.Redis = RedisConnection.createRedis(defaultRedisData)
	}

	for key, redisStruct := range RedisConnection.Connections {
		if key != RedisConnection.DefaultConnectionName {
			RedisConnection.Connections[key] = RedisConnection.createRedis(redisStruct)
		}
	}
	//for i, k := range viperRedis.AllSettings() {
	//	fmt.Println(i, "-", k)
	//}

	//redis配置读取结束==========================================================================================
}
