package page

import (
	"code/gen/auto"
	"code/gen/mapper"
	"code/gen/service"
	"code/gen/util/conf"
	"code/gen/util/logger"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"strings"
)

// 代码生成器
/**
 	1.选择表
	2.生成代码
*/
func AutoScreen(win fyne.Window) fyne.CanvasObject {
	tableName := widget.NewEntry()
	tableName.SetPlaceHolder("表名")
	tableName.Resize(fyne.NewSize(30, 5))
	var tableListData []auto.TableInfo
	if err := conf.Database.GetDB(); err == nil {
		err, tableListData = mapper.GetTables(conf.Database.DatabaseName)
		if err != nil {
			dialog.ShowError(err, win)
		}
	} else {
		dialog.ShowError(errors.New("连接失败，请修改数据配置"), win)
	}
	tableList := widget.NewTable(
		func() (int, int) {
			return len(tableListData), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			switch id.Col {
			case 0:
				label := cell.(*widget.Label)
				if tableListData[id.Row].Checked {
					label.SetText("√")
				} else {
					label.SetText("")
				}

			case 1:
				label := cell.(*widget.Label)
				label.SetText(tableListData[id.Row].TableName)
			case 2:
				label := cell.(*widget.Label)
				label.SetText(tableListData[id.Row].TableComment)
			case 3:
				label := cell.(*widget.Label)
				label.SetText(tableListData[id.Row].CreateTime.Format("2006-01-02"))
			}
		})
	tableList.OnSelected = func(id widget.TableCellID) {
		if len(tableListData) > id.Row {
			fmt.Println(len(tableListData), id.Row)
			tableListData[id.Row].Checked = !tableListData[id.Row].Checked
			tableList.Unselect(id)
			tableList.Refresh()
		}
	}

	queryBut := widget.NewButton("查询", func() {
		var err error
		if err = conf.Database.GetDB(); err != nil {
			dialog.ShowError(errors.New("连接失败，请修改数据库配置"), win)
			return
		}

		err, tableListData = mapper.GetTables(conf.Database.DatabaseName)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			if len(tableName.Text) > 0 {
				var temp []auto.TableInfo
				for _, v := range tableListData {
					if strings.Index(v.TableName, tableName.Text) >= 0 {
						temp = append(temp, v)
					}
				}
				tableListData = temp
			}
			if len(tableListData) > 0 {
				tableList.Refresh()
				fr := widget.TableCellID{0, 0}
				tableList.Select(fr)
				//tableList.Unselect(fr)
			}

		}
	})
	//f := 0.0
	//data := binding.BindFloat(&f)
	autocodeBut := widget.NewButton("生成代码", func() {
		logger.Log.WithFields(logrus.Fields{"data": ""}).Info("代码生成开始")
		var selectedTable []auto.TableInfo
		for _, v := range tableListData {
			if v.Checked {
				selectedTable = append(selectedTable, v)
			}
		}
		if len(selectedTable) > 0 {

			//withData := widget.NewProgressBarWithData(data)
			//dialog.ShowCustom("ing...","",container.NewMax(withData),win)
			if err := service.Autocode(selectedTable); err != nil {
				dialog.ShowError(err, win)
			} else {
				dialog.ShowInformation("提示", "代码生成成功", win)
			}
		} else {
			dialog.ShowError(errors.New("请选择表"), win)
		}
	})

	tableList.SetColumnWidth(0, 30)
	tableList.SetColumnWidth(1, 280)
	tableList.SetColumnWidth(2, 180)
	tableList.SetColumnWidth(3, 100)
	return container.NewBorder(container.NewVBox(
		container.NewGridWithColumns(4,
			tableName, queryBut, widget.NewLabel(""), autocodeBut,
		)), nil, nil, nil, tableList,
	)
}
