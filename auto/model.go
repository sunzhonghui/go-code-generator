package auto

import (
	"code/gen/util/conf"
	"time"
)

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	conf.ProjectConf
	TableInfo
	Fields     []TableColumnInfo `json:"fields"`
	StructName string            `json:"structName"`
}

type GLOBALMODEL struct {
	ID         int64     `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updatedTime"`
}
