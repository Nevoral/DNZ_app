package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
)

var logger *zap.Logger

func InitLogger() {
	logWriter := &lumberjack.Logger{
		Filename:   "./logs/appLogger.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logWriter),
		zapcore.DebugLevel,
	)

	logger = zap.New(core)
}

func getCaller() zap.Field {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return zap.String("caller", "unknown")
	}

	pc, _, _, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)
	return zap.String("caller", fn.Name()+" "+file+":"+string(rune(line)))
}

func DebugLog(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func InfoLog(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func WarnLog(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func ErrorLog(msg string, fields ...zap.Field) {
	logger.Error(msg, append(fields, getCaller())...)
}

func DPanicLog(msg string, fields ...zap.Field) {
	logger.DPanic(msg, append(fields, getCaller())...)
}

func PanicLog(msg string, fields ...zap.Field) {
	logger.Panic(msg, append(fields, getCaller())...)
}

func FatalLog(msg string, fields ...zap.Field) {
	logger.Fatal(msg, append(fields, getCaller())...)
}
