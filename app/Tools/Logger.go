package Tools

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	ApiLog      *zap.SugaredLogger
	QueueLog    *zap.SugaredLogger
	ScheduleLog *zap.SugaredLogger
	ArtisanLog  *zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{
		ApiLog: newSugaredLogger("api.log"),
	}
}

func newSugaredLogger(path string) *zap.SugaredLogger {
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
	return zapcore.NewConsoleEncoder(encoderConfig)
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
