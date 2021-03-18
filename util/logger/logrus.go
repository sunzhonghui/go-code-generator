package logger

import (
	"github.com/sirupsen/logrus"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"github.com/rifflock/lfshook"

	"os"
)

var Log = logrus.New()

func InitLog() {
	Log.SetReportCaller(true)
	// 设置日志级别为xx以及以上
	if level, err := logrus.ParseLevel("info"); err != nil {
		Log.Fatal(err.Error())
	} else {
		Log.SetLevel(level)
	}
	// 设置日志格式为json格式
	//Log.SetFormatter(&logrus.JSONFormatter{
	//	PrettyPrint: true,//格式化json
	//	TimestampFormat: "2006-01-02 15:04:05",//时间格式化
	//})
	Log.SetReportCaller(false)
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		// FullTimestamp:true,
		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
		// DisableLevelTruncation:true,
	})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）

	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)
	infoWriter := getWriter("info")
	errorWriter := getWriter("error")
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: infoWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  infoWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: errorWriter,
		logrus.PanicLevel: errorWriter,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"})

	Log.Hooks.Add(lfsHook)
}
func getWriter(ty string) *rotatelogs.RotateLogs {
	logPath := "log/"
	writer, err := rotatelogs.New(
		//这是分割代码的命名规则，要和下面WithRotationTime时间精度一致。要是分钟都是分钟
		logPath+ty,
		rotatelogs.WithLinkName(logPath+ty+".out"),
		rotatelogs.WithMaxAge(-1), //需要手动禁用禁用  默认情况下不清除日志，
		rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		Log.WithField("err", err).Panic("配置日志错误")
		return nil
	}
	return writer
}
