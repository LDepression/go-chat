/**
 * @Author: lenovo
 * @Description:
 * @File:  work
 * @Version: 1.0.0
 * @Date: 2023/03/21 8:08
 */

package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/pkg/goroutine/work"
)

type worker struct {
}

func (w worker) Init() {
	global.Worker = work.Init(work.Config{
		TaskChanCapacity:   global.Settings.Work.TaskChanCapacity,
		WorkerChanCapacity: global.Settings.Work.WorkerChanCapacity,
		WorkerNum:          global.Settings.Work.WorkerNum,
	})
}
