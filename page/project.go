package page

import (
	"code/gen/util/conf"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

//项目配置
/**
 	1.项目名称（注释用）
	2.项目英文缩写（生成的业务文件夹名称）
	3.项目mod名称 （import 的项目名称）
*/
func ProjcetScreen(win fyne.Window) fyne.CanvasObject {
	projectName := widget.NewEntry()
	projectName.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入项目名称")
	projectName.SetPlaceHolder("Go代码生成器v1")
	projectName.SetText(conf.Project.Name)

	projectAbbr := widget.NewEntry()
	projectAbbr.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入模块缩写")
	projectAbbr.SetPlaceHolder("go-gen")
	projectAbbr.SetText(conf.Project.Abbr)

	projectMod := widget.NewEntry()
	projectMod.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入mod名称")
	projectMod.SetPlaceHolder("go-gen")
	projectMod.SetText(conf.Project.ModName)

	projectRouter := widget.NewEntry()
	projectRouter.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入路由前缀地址")
	projectRouter.SetPlaceHolder("api/test")
	projectRouter.SetText(conf.Project.RouterName)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "项目名称：", Widget: projectName, HintText: "输入项目名称（注释用）"},
			{Text: "mod名称：", Widget: projectMod, HintText: "输入项目mod名称 （import 的项目名称）"},
			{Text: "模块缩写：", Widget: projectAbbr, HintText: "输入模块英文缩写（生成的业务文件夹名称,可带'/'但不要'/'结尾）"},
			{Text: "路由前缀：", Widget: projectRouter, HintText: "输入路由前缀"},
		},
		OnCancel: func() {
			projectName.SetText(conf.Project.Name)
			projectAbbr.SetText(conf.Project.Abbr)
			projectMod.SetText(conf.Project.ModName)
			projectRouter.SetText(conf.Project.RouterName)
		},
		OnSubmit: func() {
			conf.Project.Name = projectName.Text
			conf.Project.Abbr = projectAbbr.Text
			conf.Project.ModName = projectMod.Text
			conf.Project.RouterName = projectRouter.Text
			conf.Project.Save()
			//conf.ResetData()
			dialog.ShowInformation("提示", "保存成功", win)
		},
	}
	form.CancelText = "重置"
	form.SubmitText = "保存"

	return form
}
