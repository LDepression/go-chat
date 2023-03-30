/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:41
 */

package setting

type group struct {
	Dao    mdao
	Log    elog
	Va     va
	Worker worker
	Maker  maker
	Pager  page
}

var Group = new(group)

func InitAll() {
	Group.Dao.Init()
	Group.Log.Init()
	Group.Va.InitTrans("zh")
	Group.Worker.Init()
	Group.Maker.Init()
	Group.Pager.Init()
}
