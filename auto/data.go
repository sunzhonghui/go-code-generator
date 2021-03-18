package auto

import (
	"time"
)

//数据库表信息
type TableInfo struct {
	TableName     string    `json:"tableName"`
	TableComment  string    `json:"tableComment"`
	CreateTime    time.Time `json:"createTime"`
	Checked       bool      `json:"checked"`
	SupStructName string    `json:"supStructName"`
}

//表列信息
type TableColumnInfo struct {
	ColumnName    string `json:"columnName"`
	DataType      string `json:"dataType"`
	ColumeComment string `json:"columeComment"`

	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
	FieldJson string `json:"fieldJson"`
}
