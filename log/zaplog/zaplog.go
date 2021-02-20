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

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel int8 = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

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
	os.MkdirAll(strPath, os.ModePerm)

	strFilename := strPath + strNoExt + ".log"
	return strFilename
}

// NewZapLog 得到新的ZapLog
func init() {
	if theZap != nil {
		fmt.Println("[√] zaplog 正常保存日志")
		return
	}
	NewSugar(autologfilename(), 50, true, true, 60, "2006-01-02 15:04:05.000", DebugLevel)

}

// NewSugar 新建一个糖
// usage : NewSugar(autologfilename(), 50, true, true, 60, "2006-01-02 15:04:05.000", DebugLevel)
// @strFilename 保存的文件名
// @nMaxSizeMB 截取的文件大小, 每隔多少MB截取
// @bLocalTime 是否使用本地时间
// @bCompress 是否压缩
// @nMaxAge 文件最多保存多少天
// @strTimeFormat 时间格式
// @nLevel 日志的保存等级
func NewSugar(strFilename string, nMaxSizeMB int, bLocalTime bool, bCompress bool, nMaxAge int, strTimeFormat string, nLevel int8) {
	if theZap != nil {
		theZap.Sync()
		theZap = nil
	}

	fmt.Println("[√] zaplog 正常保存日志", strFilename)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  strFilename,
		MaxSize:   nMaxSizeMB,
		LocalTime: bLocalTime,
		Compress:  bCompress, // 是否压缩
		MaxAge:    nMaxAge,   // 文件最多保存多少天
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(strTimeFormat)) // 时间格式
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zapcore.Level(nLevel)))

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
