package Tools

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	ApiLog      *zap.SugaredLogger
	QueueLog    *zap.SugaredLogger
	ScheduleLog *zap.SugaredLogger
	ArtisanLog  *zap.SugaredLogger
	AccessLog   *zap.SugaredLogger
	ErrorLog    *zap.SugaredLogger
}

var Log *Logger

func NewLogger() {
	Log = &Logger{
		ApiLog:      newSugaredLogger("api"),
		QueueLog:    newSugaredLogger("queue"),
		ScheduleLog: newSugaredLogger("schedule"),
		ArtisanLog:  newSugaredLogger("artisan"),
		AccessLog:   newSugaredLogger("access"),
		ErrorLog:    newSugaredLogger("error"),
	}
}

func newSugaredLogger(path string) *zap.SugaredLogger {
	absPath, _ := os.Getwd()
	path = absPath + "/storage/logs/" + path
	_, fileExistErr := os.Stat(path + "/log.txt")
	if fileExistErr != nil && os.IsNotExist(fileExistErr) {
		os.MkdirAll(path, os.ModePerm)
		_, error2 := os.Create(path + "/log.txt")
		if error2 != nil {
			fmt.Println("【"+path+"】日志文件创建失败：", error2)
			os.Exit(1)
		}
	}
	path = path + "/log.txt"

	writeSyncer := getLogWriter(path)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    2,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
