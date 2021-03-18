// Package main provides various examples of Fyne API capabilities.
package main

import (
	"code/gen/page"
	"code/gen/runner"
	"code/gen/util/logger"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func init() {
	runner.Runner()
}

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

func main() {
	a := app.NewWithID("com.idmiss.generator")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("Golang 代码生成器")
	w.SetFixedSize(true)
	topWindow = w
	tutorial := page.Pages["welcome"].SetView(w)

	databaseItem := fyne.NewMenuItem("数据库配置", func() {
		logger.Log.WithFields(logrus.Fields{"data": ""}).Info("数据库配置")
		//tutorial = container.NewBorder(container.NewVBox(widget.NewLabelWithStyle(page.Pages["database"].Title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), widget.NewSeparator()), nil, nil, nil,page.Pages["database"].View(w))
		tutorial = page.Pages["database"].SetView(w)
		w.SetContent(tutorial)
	})
	projectItem := fyne.NewMenuItem("项目配置", func() {
		tutorial = page.Pages["project"].SetView(w)
		w.SetContent(tutorial)
		logger.Log.WithFields(logrus.Fields{"data": ""}).Info("项目配置")
	})
	fileItem := fyne.NewMenuItem("模板文件", func() {
		pwd, _ := os.Getwd()
		u, _ := url.Parse(pwd + "/resource/temp/")
		_ = a.OpenURL(u)
	})

	helpMenu := fyne.NewMenu("帮助",
		fyne.NewMenuItem("查看文档", func() {
			u, _ := url.Parse("http://www.idmiss.com/")
			_ = a.OpenURL(u)
		}))

	welcome := fyne.NewMenuItem("首页", func() {
		//tutorial := container.NewBorder(container.NewVBox(widget.NewLabelWithStyle(page.Pages["welcome"].Title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), widget.NewSeparator()), nil, nil, nil,page.Pages["welcome"].View(w))
		tutorial = page.Pages["welcome"].SetView(w)
		w.SetContent(tutorial)
	})
	autocodeMenu := fyne.NewMenuItem("代码生成", func() {
		tutorial = page.Pages["autocode"].SetView(w)
		w.SetContent(tutorial)
		logger.Log.WithFields(logrus.Fields{"data": ""}).Info("代码生成")
	})
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("文件", welcome),
		fyne.NewMenu("设置", databaseItem, fyne.NewMenuItemSeparator(), projectItem, fyne.NewMenuItemSeparator(), fileItem),
		fyne.NewMenu("生成器", autocodeMenu),
		helpMenu,
	)

	w.SetMainMenu(mainMenu)
	w.SetMaster()

	w.SetContent(tutorial)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()

	os.Unsetenv("FYNE_FONT")
}
