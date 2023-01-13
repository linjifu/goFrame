package Tools

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
	"time"
)

// 单个连接结构体
type DbStruct struct {
	Name            string      `mapstructure:"name"`            //链接名称
	Driver          string      `mapstructure:"driver"`          //驱动
	Host            string      `mapstructure:"host"`            //host
	Port            string      `mapstructure:"port"`            //端口
	UserName        string      `mapstructure:"userName"`        //用户名
	Password        string      `mapstructure:"password"`        //密码
	Charset         string      `mapstructure:"charset"`         //字符编码
	Database        string      `mapstructure:"database"`        //数据库名称
	MaxIdleConns    int         `mapstructure:"maxIdleConns"`    //空闲连接池中连接的最大数量
	MaxOpenConns    int         `mapstructure:"maxOpenConns"`    //打开数据库连接的最大数量
	ConnMaxLifetime int         `mapstructure:"connMaxLifetime"` //连接可复用的最大时间（分钟）
	Read            []*DbStruct `mapstructure:"read"`            //读连接
	Write           []*DbStruct `mapstructure:"write"`           //写连接
	Connection      *gorm.DB    //真实连接
}

type Dbs struct {
	DefaultConnectionName string               `mapstructure:"default"` //默认连接名称
	Db                    *DbStruct            //当前连接的数据
	Connections           map[string]*DbStruct `mapstructure:"connections"` //所有连接的数据
}

var DBConnection = new(Dbs)

func init() {
	path, _ := os.Getwd()
	viperModel := viper.New()
	viperModel.SetDefault("charset", "utf8mb4")
	//导入配置文件
	viperModel.SetConfigType("yaml")
	viperModel.SetConfigFile(path + "/config/database.yml")
	//读取配置文件
	err := viperModel.ReadInConfig()
	if err != nil {
		fmt.Println()
		panic(fmt.Errorf("读取配置文件database数据与结构体转换失败:%s \n", err))

	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viperModel.Unmarshal(DBConnection); err != nil {
		fmt.Println()
		panic(fmt.Errorf("配置文件database数据与结构体转换失败:%s \n", err))
	}

	//判断默认连接是否配置
	defaultData, ok := DBConnection.Connections[DBConnection.DefaultConnectionName]
	if ok {
		DBConnection.Db = createDB(defaultData)
	}

	for key, dbStruct := range DBConnection.Connections {
		if key != DBConnection.DefaultConnectionName {
			DBConnection.Connections[key] = createDB(dbStruct)
		}
	}
}

// 创建DB
func createDB(dbStruct *DbStruct) *DbStruct {
	//默认字符编码
	charset := "utf8mb4"
	if dbStruct.Charset == "" {
		dbStruct.Charset = charset
	}
	dsn := getDsn(dbStruct)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println()
		panic(fmt.Errorf(dbStruct.Name+"数据库链接失败:%s \n", err))
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println()
			panic(fmt.Errorf(dbStruct.Name+"设置连接池失败:%s \n", err))
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(dbStruct.MaxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(dbStruct.MaxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Duration(dbStruct.ConnMaxLifetime) * time.Minute)

		fmt.Println(dbStruct.Name + "数据库链接成功")
		if dbStruct.Read != nil || dbStruct.Write != nil {
			sources := []gorm.Dialector{}
			replicas := []gorm.Dialector{}
			//读链接
			for _, readDbStruct := range dbStruct.Read {
				readDbStruct = setDefauleValue(readDbStruct, dbStruct)
				replicasDsn := getDsn(readDbStruct)
				replicas = append(replicas, mysql.Open(replicasDsn))
			}
			//写链接
			for _, writeDbStruct := range dbStruct.Write {
				writeDbStruct = setDefauleValue(writeDbStruct, dbStruct)
				sourcesDsn := getDsn(writeDbStruct)
				sources = append(sources, mysql.Open(sourcesDsn))
			}
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  sources,
				Replicas: replicas,
				// sources/replicas 负载均衡策略
				Policy: dbresolver.RandomPolicy{},
			}))
		}

		dbStruct.Connection = db
	}

	return dbStruct
}

// 生成连接dsn
func getDsn(dbStruct *DbStruct) string {
	return dbStruct.UserName + ":" + dbStruct.Password + "@tcp(" + dbStruct.Host + ":" + dbStruct.Port + ")/" + dbStruct.Database + "?charset=" + dbStruct.Charset + "&parseTime=True&loc=Asia%2FShanghai"
}

// 设置默认值
func setDefauleValue(new, old *DbStruct) *DbStruct {
	if new.Charset == "" {
		new.Charset = old.Charset
	}
	if new.UserName == "" {
		new.UserName = old.UserName
	}
	if new.Password == "" {
		new.Password = old.Password
	}
	if new.Port == "" {
		new.Port = old.Port
	}
	if new.Database == "" {
		new.Database = old.Database
	}
	if new.Driver == "" {
		new.Driver = old.Driver
	}
	if new.MaxIdleConns == 0 {
		new.MaxIdleConns = old.MaxIdleConns
	}
	if new.MaxOpenConns == 0 {
		new.MaxOpenConns = old.MaxOpenConns
	}
	if new.ConnMaxLifetime == 0 {
		new.ConnMaxLifetime = old.ConnMaxLifetime
	}
	return new
}
