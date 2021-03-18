package mapper

import (
	"{{.ModName}}/util/global"
    "{{.ModName}}/{{.Abbr}}/model"
)
//创建
func Create{{.SupStructName}}(data model.{{.SupStructName}}) (err error) {
	err = global.DB.Create(&data).Error
	return
}
//根据ID删除
func Delete{{.SupStructName}}(data model.{{.SupStructName}}) (err error) {
	err = global.DB.Delete(&data).Error
	return
}
//根据ID批量删除
func Delete{{.SupStructName}}ByIds(ids []int64) (err error) {
	err = global.DB.Delete(&model.{{.SupStructName}}{},"id in ?",ids).Error
	return
}

//根据id 更新 ，排除零值
func Update{{.SupStructName}}(data model.{{.SupStructName}}) (err error) {
	err = global.DB.Updates(&data).Error
	return
}

//根据id 更新， Save 会保存所有的字段，即使字段是零值 没有记录则insert
func Save{{.SupStructName}}(data model.{{.SupStructName}}) (err error) {
	err = global.DB.Save(&data).Error
	return
}

//根据id获取model
func Get{{.SupStructName}}(id int64) (err error, data model.{{.SupStructName}}) {
	err = global.DB.Where("id = ?", id).First(&data).Error
	return
}

//获取所有的model
func GetAll{{.SupStructName}}() (err error, list []model.{{.SupStructName}}) {
	err = global.DB.Find(&list).Error
	return
}

//按条件分页查询 limit offset ,参数用指针&, 数据会自动填充到req对象
func GetPageLimit{{.SupStructName}}(data *model.{{.SupStructName}}Req) (err error) {

    var list []model.{{.SupStructName}}
    if err = global.DB.Model(&data.{{.SupStructName}}).Where(&data.{{.SupStructName}}).Count(&data.Total).Error; err != nil {
        return
    }
    if err = global.DB.Where(&data.{{.SupStructName}}).Limit(int(data.PageSize)).Offset(int(data.Offset())).Find(&list).Error; err != nil {
        return
    }
    data.Data=list
    return
}