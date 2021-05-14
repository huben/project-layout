package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder  {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer  {
	lumberjackLogger := &lumberjack.Logger{
		Filename: "./log/test.log",
		MaxSize: 1, // MB
		MaxBackups: 5, //个数
		MaxAge: 30, // day
		Compress: false,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func Debug(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args)
}

func Info(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

func Warn(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args)
}

func Panic(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args)
}

func Fatal(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args)
}

func Sync()  {
	sugarLogger.Sync()
}