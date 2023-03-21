/**
 * @Author: lenovo
 * @Description:
 * @File:  log
 * @Version: 1.0.0
 * @Date: 2023/03/20 15:32
 */

package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/pkg/logger"
	"go.uber.org/zap"
)

type elog struct{}

func (elog) Init() {
	//初始化
	//logger, _ := zap.NewDevelopment()

	/*
				当我们使用生产者模式的时候,我们可以使用Debug
			因为product时候,日志级别是info,要高于debug,所以debug打印不出来
		S和L函数提供了安全访问全局logger的功能(),因为里面加了锁
	*/
	//zap.ReplaceGlobals(logger)
	global.Logger = logger.NewLogger(&logger.InitStruct{
		LogSavePath:   global.Settings.Log.LogSavePath,
		LogFileExt:    global.Settings.Log.LogFileExt,
		MaxSize:       global.Settings.Log.MaxSize,
		MaxBackups:    global.Settings.Log.MaxBackups,
		MaxAge:        global.Settings.Log.MaxAge,
		Compress:      global.Settings.Log.Compress,
		LowLevelFile:  global.Settings.Log.LowLevelFile,
		HighLevelFile: global.Settings.Log.HighLevelFile,
	}, global.Settings.Log.Level)
	zap.ReplaceGlobals(global.Logger.Logger)
}
