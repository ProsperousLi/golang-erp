package logs

import (
	//"strings"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var FileLogs *logs.BeeLogger

//运行方式
var runmode string

func InitLogs() {
	// consoleLogs = logs.NewLogger(1)
	// consoleLogs.SetLogger(logs.AdapterConsole)
	// consoleLogs.Async() //异步
	// fileLogs = logs.NewLogger(10000)
	// level := beego.AppConfig.String("logs::level")
	// fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/rms.log",
	// 	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
	// 	"level":`+level+`,
	// 	"daily":true,
	// 	"maxdays":10}`)
	// fileLogs.Async() //异步
	// runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	// if runmode == "" {
	// 	runmode = "dev"
	// }
	// 设置配置文件
	jsonConfig := `{
        "filename" : "/home/dragon/gocode/src/erpweb/log/test.log",
        "maxlines" : 1000,       
        "maxsize"  : 10240       
    }`
	FileLogs = logs.NewLogger(10000)
	FileLogs.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	FileLogs.SetLevel(logs.LevelError)     // 设置日志写入缓冲区的等级
	FileLogs.SetLevel(logs.LevelInfo)
	FileLogs.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	FileLogs.Async()
	Info("log start...")
	//log.Info("log start...")
}
func LogEmergency(v interface{}) {
	log("emergency", v)
}
func LogAlert(v interface{}) {
	log("alert", v)
}
func LogCritical(v interface{}) {
	log("critical", v)
}
func Error(v interface{}) {
	log("error", v)
}
func Warn(v interface{}) {
	log("warning", v)
}
func LogNotice(v interface{}) {
	log("notice", v)
}
func Info(v interface{}) {
	log("info", v)
}
func LogDebug(v interface{}) {
	log("debug", v)
}

func LogTrace(v interface{}) {
	log("trace", v)
}

//Log 输出日志
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	//if runmode == "dev" {
	switch level {
	case "emergency":
		FileLogs.Emergency(format, v)
	case "alert":
		FileLogs.Alert(format, v)
	case "critical":
		FileLogs.Critical(format, v)
	case "error":
		FileLogs.Error(format, v)
	case "warning":
		FileLogs.Warning(format, v)
	case "notice":
		FileLogs.Notice(format, v)
	case "info":
		FileLogs.Info(format, v)
	case "debug":
		FileLogs.Debug(format, v)
	case "trace":
		FileLogs.Trace(format, v)
	default:
		FileLogs.Debug(format, v)
	}
	//}
	// switch level {
	// case "emergency":
	// 	consoleLogs.Emergency(format, v)
	// case "alert":
	// 	consoleLogs.Alert(format, v)
	// case "critical":
	// 	consoleLogs.Critical(format, v)
	// case "error":
	// 	consoleLogs.Error(format, v)
	// case "warning":
	// 	consoleLogs.Warning(format, v)
	// case "notice":
	// 	consoleLogs.Notice(format, v)
	// case "info":
	// 	consoleLogs.Info(format, v)
	// case "debug":
	// 	consoleLogs.Debug(format, v)
	// case "trace":
	// 	consoleLogs.Trace(format, v)
	// default:
	// 	consoleLogs.Debug(format, v)
	// }
}
