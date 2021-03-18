
package model

import "time"
import "{{.ModName}}/util/page"

// Model {{.TableName}}  {{.TableComment}}
type {{.SupStructName}} struct {
      {{- range .Fields}}
      {{.FieldName}} {{.FieldType}} `json:"{{.FieldJson}}" gorm:"default:default:(-)"` //{{.ColumeComment}}{{ end }}
}

{{ if .TableName }}
func ({{.SupStructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}

type {{.SupStructName}}Req struct {
	page.PageData
	{{.SupStructName}}
}