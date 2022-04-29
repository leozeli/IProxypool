package main

import (
	"github.com/leozeli/IProxypool/api"
	"github.com/leozeli/IProxypool/cmd"
	"github.com/leozeli/IProxypool/middleware/config"
	"github.com/leozeli/IProxypool/middleware/database"
	"github.com/leozeli/IProxypool/middleware/logutil"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 检查或设置命令行参数
	cmd.Execute()

	setting := config.ServerSetting

	// 将日志写入文件或打印到控制台
	logutil.InitLog(&setting.Log)
	// 初始化数据库连接
	database.InitDB(&setting.Database)

	// Start HTTP
	//go func() {
	//	api.Run(&setting.System)
	//}()
	for {
		api.Run(&setting.System)
	}

	// Start Task
	//todo: 改为api 触发
	//run.Task()
}
