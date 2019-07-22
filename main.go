package main

import (
	"erpweb/logs"
	"erpweb/models"
	_ "erpweb/routers"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
)

// func InitLog(log *logs.BeeLogger) {

// 	// 设置配置文件
// 	jsonConfig := `{
//         "filename" : "/home/dragon/gocode/src/erpweb/test.log",
//         "maxlines" : 1000,
//         "maxsize"  : 10240
//     }`
// 	log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
// 	log.SetLevel(logs.LevelError)     // 设置日志写入缓冲区的等级
// 	log.SetLevel(logs.LevelInfo)
// 	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
// 	log.Async()
// 	log.Info("log start...")
// }

func main() {
	logs.InitLogs()

	logs.Info("hahah")

	//init sql
	models.Init()

	beego.Run()
}
