package Tools

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
	"strings"
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
	DefaultConnectionName string               `mapstructure:"dbDefault"` //默认连接名称
	Db                    *DbStruct            //当前连接的数据
	Connections           map[string]*DbStruct `mapstructure:"dbConnections"` //所有连接的数据
}

func NewDbs() *Dbs {
	if dbConnection == nil {
		dbConnection = &Dbs{}
		dbConnection.LoadMysql()
	}
	return dbConnection
}

var dbConnection *Dbs = nil

func (r *Dbs) LoadMysql() {
	path, _ := os.Getwd()
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
	if err2 := viperDatabase.Unmarshal(dbConnection); err2 != nil {
		fmt.Println()
		panic(fmt.Errorf("配置文件database数据与结构体转换失败:%s \n", err2))
	}
	//判断默认连接是否配置
	defaultDBData, ok := dbConnection.Connections[dbConnection.DefaultConnectionName]
	if ok {
		dbConnection.Db = dbConnection.createDB(defaultDBData)
	}

	for key, dbStruct := range dbConnection.Connections {
		if key != dbConnection.DefaultConnectionName {
			dbConnection.Connections[key] = dbConnection.createDB(dbStruct)
		}
	}
	//for i, k := range viperDatabase.AllSettings() {
	//	fmt.Println(i, "-", k)
	//}
}

// 创建DB
func (r *Dbs) createDB(dbStruct *DbStruct) *DbStruct {
	//默认字符编码
	charset := "utf8mb4"
	if dbStruct.Charset == "" {
		dbStruct.Charset = charset
	}
	dsn := dbConnection.getDsn(dbStruct)
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
				readDbStruct = dbConnection.setDefauleValue(readDbStruct, dbStruct)
				replicasDsn := dbConnection.getDsn(readDbStruct)
				replicas = append(replicas, mysql.Open(replicasDsn))
			}
			//写链接
			for _, writeDbStruct := range dbStruct.Write {
				writeDbStruct = dbConnection.setDefauleValue(writeDbStruct, dbStruct)
				sourcesDsn := dbConnection.getDsn(writeDbStruct)
				sources = append(sources, mysql.Open(sourcesDsn))
			}
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  sources,
				Replicas: replicas,
				// sources/replicas 负载均衡策略
				Policy: dbresolver.RandomPolicy{},
			}).SetMaxIdleConns(dbStruct.MaxIdleConns).SetMaxOpenConns(dbStruct.MaxOpenConns).SetConnMaxLifetime(time.Duration(dbStruct.ConnMaxLifetime) * time.Minute))
		}

		dbStruct.Connection = db
	}

	return dbStruct
}

// 生成连接dsn
func (r *Dbs) getDsn(dbStruct *DbStruct) string {
	return dbStruct.UserName + ":" + dbStruct.Password + "@tcp(" + dbStruct.Host + ":" + dbStruct.Port + ")/" + dbStruct.Database + "?charset=" + dbStruct.Charset + "&parseTime=True&loc=Asia%2FShanghai"
}

// 设置默认值
func (r *Dbs) setDefauleValue(new, old *DbStruct) *DbStruct {
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

// GetDB 获取DB连接
func GetDB(name ...string) (*gorm.DB, error) {
	if len(name) == 0 {
		return dbConnection.Db.Connection, nil
	} else {
		if len(name) != 1 {
			return nil, errors.New("请正确选择要获取的连接名称")
		}
		return dbConnection.Connections[name[0]].Connection, nil
	}
}
