package page

import (
	"code/gen/util/conf"
	"code/gen/util/logger"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseScreen(win fyne.Window) fyne.CanvasObject {
	urlName := widget.NewEntry()
	urlName.SetPlaceHolder("本地连接")
	urlName.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入连接名称")
	urlName.SetText(conf.Database.UrlName)

	ip := widget.NewEntry()
	ip.SetPlaceHolder("192.168.1.1")
	ip.Validator = validation.NewRegexp(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`, "输入正确得ip")
	ip.SetText(conf.Database.IP)

	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
	port.Validator = validation.NewRegexp(`^[0-9]*$`, "输入正确得端口")
	port.SetText(conf.Database.Port)

	databaseName := widget.NewEntry()
	databaseName.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入数据库名称")
	databaseName.SetPlaceHolder("database1")
	databaseName.SetText(conf.Database.DatabaseName)

	userName := widget.NewEntry()
	userName.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入用户名")
	userName.SetPlaceHolder("root")
	userName.SetText(conf.Database.UserName)

	password := widget.NewPasswordEntry()
	password.Validator = validation.NewRegexp(`^.{1,50}$`, "请输入密码")
	password.SetPlaceHolder("root")
	password.SetText(conf.Database.Password)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "连接名：", Widget: urlName, HintText: "输入连接名称"},
			{Text: "主机：", Widget: ip, HintText: "输入域名或者IP地址"},
			{Text: "端口：", Widget: port, HintText: "输入数据库端口"},
			{Text: "数据库名：", Widget: databaseName, HintText: "输入数据库名"},
			{Text: "用户名：", Widget: userName, HintText: "输入用户名,默认root"},
			{Text: "密码：", Widget: password, HintText: "输入密码,默认root"},
		},
		OnCancel: func() {
			mysqlUrl := userName.Text + ":" + password.Text + "@(" + ip.Text + ":" + port.Text + ")/" + databaseName.Text + "?charset=utf8&parseTime=True&loc=Local"
			logger.Log.WithFields(logrus.Fields{"data": mysqlUrl}).Info("数据库连接地址")
			db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
			if err != nil {
				dialog.ShowError(errors.New("连接失败"), win)
			} else {
				conf.DB = db
				dialog.ShowInformation("提示", "连接成功", win)
			}
		},
		OnSubmit: func() {
			conf.Database.UrlName = urlName.Text
			conf.Database.IP = ip.Text
			conf.Database.Port = port.Text
			conf.Database.DatabaseName = databaseName.Text
			conf.Database.UserName = userName.Text
			conf.Database.Password = password.Text
			conf.Database.Save()
			//conf.ResetData()
			mysqlUrl := userName.Text + ":" + password.Text + "@(" + ip.Text + ":" + port.Text + ")/" + databaseName.Text + "?charset=utf8&parseTime=True&loc=Local"
			logger.Log.WithFields(logrus.Fields{"data": mysqlUrl}).Info("数据库连接地址")
			db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
			if err != nil {
				dialog.ShowError(errors.New("连接失败，请修改"), win)
			} else {
				conf.DB = db
				dialog.ShowInformation("提示", "保存成功", win)
			}
		},
	}
	form.CancelText = "测试"
	form.SubmitText = "保存"

	return form
}
