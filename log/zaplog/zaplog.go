package zaplog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// import log "github.com/yangtizi/go/log/zaplog"

// theZap 新的日志库,据说性能更好
var theZap *zap.SugaredLogger

// 不要扩展名的文件
func noExt(path string) string {
	path = filepath.Base(path)
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[0:i]
		}
	}
	return path
}

func autologfilename() string {
	strNoExt := noExt(os.Args[0])
	strPath := "./log/" + strNoExt + "/"
	fmt.Println("[√] zaplog 正常保存日志", strPath)
	os.MkdirAll(strPath, os.ModePerm)

	strFilename := strPath + strNoExt + ".log"
	return strFilename
}

// NewZapLog 得到新的ZapLog
func init() {
	level := zapcore.DebugLevel
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  autologfilename(),
		MaxSize:   50,
		LocalTime: true,
		Compress:  true, // 是否压缩
		MaxAge:    60,   // 文件最多保存多少天
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))

	theZap = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

// Printf 为了兼容
func Printf(template string, args ...interface{}) {
	theZap.Debugf(template, args...)
}

// Print 为了兼容
func Print(args ...interface{}) {
	theZap.Info(args...)
}

// Println 为了兼容
func Println(args ...interface{}) {
	theZap.Info(args...)
}

// Debug ()
func Debug(args ...interface{}) {
	theZap.Debug(args...)
}

// Debugf ()
func Debugf(template string, args ...interface{}) {
	theZap.Debugf("[+] "+template, args...)
}

// Info ()
func Info(args ...interface{}) {
	theZap.Info(args...)
}

// Infof ()
func Infof(template string, args ...interface{}) {
	theZap.Infof("[√] "+template, args...)
}

// Warn ()
func Warn(args ...interface{}) {
	theZap.Warn(args...)
}

// Warnf ()
func Warnf(template string, args ...interface{}) {
	theZap.Warnf("[!] "+template, args...)
}

// Error ()
func Error(args ...interface{}) {
	fmt.Println(args...)
	theZap.Error(args...)
}

// Errorf ()
func Errorf(template string, args ...interface{}) {
	fmt.Printf(template, args...)
	theZap.Errorf("[x] "+template, args...)
}

// DPanic ()
func DPanic(args ...interface{}) {
	theZap.DPanic(args...)
}

// DPanicf ()
func DPanicf(template string, args ...interface{}) {
	theZap.DPanicf(template, args...)
}

// Panic ()
func Panic(args ...interface{}) {
	theZap.Panic(args...)
}

// Panicf ()
func Panicf(template string, args ...interface{}) {
	theZap.Panicf(template, args...)
}

// Fatal ()
func Fatal(args ...interface{}) {
	theZap.Fatal(args...)
}

// Fatalf ()
func Fatalf(template string, args ...interface{}) {
	theZap.Fatalf(template, args...)
}

// Flush ()
func Flush() {
	theZap.Sync()
}

// Since 给 defer 用的函数 defer zaplog.Since(time.Now())
func Since(t time.Time) time.Duration {
	return time.Since(t)
}
