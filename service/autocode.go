package service

import (
	"code/gen/auto"
	"code/gen/mapper"
	"code/gen/util/autocode"
	"code/gen/util/common"
	"code/gen/util/conf"
	"code/gen/util/logger"
	"code/gen/util/strcase"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func Autocode(tableList []auto.TableInfo) error {

	autoData := make([]auto.AutoCodeStruct, 0, len(tableList))

	for _, v := range tableList {
		v.SupStructName = strcase.UpperSnakeCase(v.TableName)
		_, data := mapper.GetColumns(conf.Database.DatabaseName, v.TableName)
		for i, j := range data {
			j.FieldName = strcase.UpperSnakeCase(j.ColumnName)
			if j.ColumnName == "id" {
				j.FieldName = "ID"
			}
			j.FieldJson = strcase.SnakeCase(j.ColumnName)
			j.FieldType = autocode.GetDbType(j.DataType)
			data[i] = j
		}
		autoData = append(autoData, auto.AutoCodeStruct{*conf.Project, v, data, ""})
	}

	var allTempFile []string
	pathName := "resource/temp"
	files, err := ioutil.ReadDir(pathName) // 找出所有模板文件
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": err}).Error("代码生成出错")
		return err
	}
	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".tpl") {
			allTempFile = append(allTempFile, pathName+"/"+v.Name())
		}
	}

	//DirExistAndMake(autoPath)// 检查 文件夹是否存在

	marshal, _ := json.Marshal(autoData)
	fmt.Println(string(marshal))
	fmt.Println(allTempFile)

	for _, tv := range autoData { //数据列表
		for _, fv := range allTempFile { // 文件列表
			if err := autocodeFile(tv, fv); err != nil {
				return err
			}
		}
	}

	return nil
}

func autocodeFile(tv auto.AutoCodeStruct, fv string) error {
	// 开始生成 代码
	autoPath := "resource/autocode/"
	if strings.Index(fv, "router") >= 0 {
		autoPath += "router/"
	} else {
		autoPath += tv.Abbr + "/" + fv[14:strings.Index(fv, ".")] + "/"
	}
	if err := DirExistAndMake(autoPath); err != nil {
		return err
	}

	files, err := template.ParseFiles(fv)

	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": err}).Error("代码生成出错")
		return err
	}
	if err == nil {
		structName := strcase.SnakeCase(tv.TableName)
		tv.StructName = structName

		file, _ := os.OpenFile(autoPath+structName+".go", os.O_CREATE|os.O_WRONLY, 0755)
		err := files.Execute(file, tv)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(autoPath)
	return nil
}

func DirExistAndMake(autoPath string) error {
	if !common.Exists(autoPath) { // 检查 文件夹是否存在
		if err := os.MkdirAll(autoPath, os.ModePerm); err != nil {
			logger.Log.WithFields(logrus.Fields{"data": err}).Warn("文件夹不存在，创建文件夹出错")
			return err
		}
	}
	return nil
}
