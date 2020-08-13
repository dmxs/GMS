package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"time"
)

var (
	Handle  *zap.Logger
	LogPath string
)

type KV map[string]interface{}

func Info(msg ...string) {
	Handle.Info(ConcatWithPart(", ", msg...))
}

func InfoKV(msg string, param KV) {
	Handle.Info(msg, kvToField(param)...)
}

func InfoF(format string, a ...interface{}) {
	Handle.Info(fmt.Sprintf(format, a...))
}

func Error(msg ...string) {
	Handle.Error(ConcatWithPart(", ", msg...))
}

func ErrorKV(msg string, param KV) {
	Handle.Error(msg, kvToField(param)...)
}

func ErrorF(format string, a ...interface{}) {
	Handle.Error(fmt.Sprintf(format, a...))
}

func Warn(msg ...string) {
	Handle.Warn(ConcatWithPart(", ", msg...))
}

func WarnKV(msg string, param KV) {
	Handle.Warn(msg, kvToField(param)...)
}

func Debug(msg ...string) {
	Handle.Debug(ConcatWithPart(", ", msg...))
}

func DebugKV(msg string, param KV) {
	Handle.Debug(msg, kvToField(param)...)
}

func DebugF(format string, a ...interface{}) {
	Handle.Debug(fmt.Sprintf(format, a...))
}

func InitLog() {
	applicationPath, _ := os.Getwd()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeCaller: CustomCallerEncoder,
	}

	// 生成日志目录
	LogPath = applicationPath + "/log/"
	if err := os.MkdirAll(LogPath, os.ModePerm); err != nil {
		panic("log dir init failed")
	}

	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取io.Writer的实现
	infoWriter := getWriter(LogPath + "info.log")
	errorWriter := getWriter(LogPath + "error.log")

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(errorWriter), warnLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), infoLevel), //同时将日志输出到控制台
	)

	Handle = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.InfoLevel))

	Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	Info("日志初始化成功")
	Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func kvToField(param KV) (list []zap.Field) {
	for key, value := range param {
		list = append(list, Any(key, value))
	}
	return list
}

func Any(key string, value interface{}) zap.Field {
	switch val := value.(type) {
	case error:
		return zap.Any(key, value.(error).Error())
	default:
		return zap.Any(key, val)
	}
}

func ConcatWithPart(part string, s ...string) (result string) {
	for i, v := range s {
		if i != 0 {
			result += part
		}
		result += v
	}
	return result
}

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     2, // days
	}
}

func CustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	pc, file, line, ok := runtime.Caller(6)
	caller = zapcore.EntryCaller{
		Defined: ok,
		PC:      pc,
		File:    file,
		Line:    line,
	}
	enc.AppendString(caller.TrimmedPath()) // 短路径
	//enc.AppendString(caller.String())// 长路径
}
