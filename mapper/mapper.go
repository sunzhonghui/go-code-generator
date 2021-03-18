package mapper

import (
	"code/gen/auto"
	"code/gen/util/conf"
	"errors"
)

//获取数据库所有的表
func GetTables(dbName string) (err error, data []auto.TableInfo) {
	if conf.DB == nil {
		return errors.New("未找到数据连接"), data
	}
	err = conf.DB.Raw("select table_name ,table_comment,CREATE_TIME create_time from information_schema.tables where table_schema = ? ", dbName).Scan(&data).Error
	return
}

func GetColumns(dbName, tableName string) (err error, data []auto.TableColumnInfo) {
	if conf.DB == nil {
		return errors.New("未找到数据连接"), data
	}
	err = conf.DB.Raw("SELECT COLUMN_NAME column_name,DATA_TYPE data_type,COLUMN_COMMENT colume_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName).Scan(&data).Error
	return
}
