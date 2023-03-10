package Tools

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

// 单个连接结构体
type RedisStruct struct {
	Name            string        `mapstructure:"name"`            //链接名称
	Host            string        `mapstructure:"host"`            //host
	Port            string        `mapstructure:"port"`            //端口
	UserName        string        `mapstructure:"userName"`        //用户名
	Password        string        `mapstructure:"password"`        //密码
	Database        int           `mapstructure:"database"`        //数据库名称
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`    //空闲连接池中连接的最大数量
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`    //打开数据库连接的最大数量
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"` //连接可复用的最大时间（秒）
	Pool            *redis.Pool   //连接池
}

type Redises struct {
	DefaultConnectionName string                  `mapstructure:"redisDefault"` //默认连接名称
	Redis                 *RedisStruct            //当前连接的数据
	Connections           map[string]*RedisStruct `mapstructure:"redisConnections"` //所有连接的数据
}

func NewRedises() *Redises {
	if redisConnection == nil {
		redisConnection = &Redises{}
		redisConnection.LoadRedis()
	}
	return redisConnection
}

var redisConnection *Redises = nil

func (r *Redises) LoadRedis() {
	path, _ := os.Getwd()
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
	if redisErr2 := viperRedis.Unmarshal(redisConnection); redisErr2 != nil {
		fmt.Println()
		panic(fmt.Errorf("配置文件redis数据与结构体转换失败:%s \n", redisErr2))
	}
	//判断默认连接是否配置
	defaultRedisData, ok := redisConnection.Connections[redisConnection.DefaultConnectionName]
	if ok {
		redisConnection.Redis = redisConnection.createRedis(defaultRedisData)
	}

	for key, redisStruct := range redisConnection.Connections {
		if key != redisConnection.DefaultConnectionName {
			redisConnection.Connections[key] = redisConnection.createRedis(redisStruct)
		}
	}
}

// 创建DB
func (r *Redises) createRedis(redisStruct *RedisStruct) *RedisStruct {
	//默认值设置
	redisStruct = redisConnection.setDefaultValue(redisStruct)
	//实例化一个连接池
	pool := &redis.Pool{
		MaxIdle:     redisStruct.MaxIdleConns,    //最初的连接数量
		MaxActive:   redisStruct.MaxOpenConns,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: redisStruct.ConnMaxLifetime, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp",
				redisStruct.Host+":"+redisStruct.Port,
				redis.DialUsername(redisStruct.UserName),
				redis.DialPassword(redisStruct.Password),
				redis.DialDatabase(redisStruct.Database),
			)
		},
	}

	redisStruct.Pool = pool

	fmt.Println(redisStruct.Name + "-redis链接成功")

	return redisStruct
}

// 设置默认值
func (r *Redises) setDefaultValue(redisStruct *RedisStruct) *RedisStruct {
	if redisStruct.Database == 0 {
		redisStruct.Database = 0
	}
	if redisStruct.MaxIdleConns == 0 {
		redisStruct.MaxIdleConns = 10
	}
	if redisStruct.MaxOpenConns == 0 {
		redisStruct.MaxOpenConns = 100
	}
	if redisStruct.ConnMaxLifetime == 0 {
		redisStruct.ConnMaxLifetime = 300
	}
	return redisStruct
}

// GetRedisPool 获取redis连接池
func GetRedisPool(name ...string) (*redis.Pool, error) {
	if len(name) == 0 {
		return redisConnection.Redis.Pool, nil
	} else {
		if len(name) != 1 {
			return nil, errors.New("请正确选择要获取的连接池名称")
		}
		return redisConnection.Connections[name[0]].Pool, nil
	}
}
