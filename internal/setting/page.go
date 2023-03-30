package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/pkg/app"
)

type page struct {
}

func (page) Init() {
	global.Pager = app.InitPage(
		global.Settings.Page.DefaultPageSize,
		global.Settings.Page.MaxPageSize,
		global.Settings.Page.PageKey,
		global.Settings.Page.PageSizeKey,
	)
}
