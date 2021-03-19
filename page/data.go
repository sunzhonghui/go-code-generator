package page

import (
	"code/gen/util/conf"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type PageDetail struct {
	Title, Intro string
	View         func(win fyne.Window) fyne.CanvasObject
}

var (
	Pages = map[string]*PageDetail{
		"welcome":  {Title: "Golang 代码生成器 go-code-generator " + conf.Version, View: welcomeScreen},
		"database": {Title: "设置数据库", View: DatabaseScreen},
		"project":  {Title: "项目设置", View: ProjcetScreen},
		"autocode": {Title: "代码生成", View: AutoScreen},
	}
)

func (d *PageDetail) SetView(win fyne.Window) *fyne.Container {
	return container.NewBorder(container.NewVBox(widget.NewLabelWithStyle(d.Title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), widget.NewSeparator()), nil, nil, nil, d.View(win))
}
