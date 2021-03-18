
package service

import (
     "errors"
     "github.com/sirupsen/logrus"
     "gorm.io/gorm"

	 "{{.ModName}}/{{.Abbr}}/mapper"
     "{{.ModName}}/{{.Abbr}}/model"

     "{{.ModName}}/util/logger"
     "{{.ModName}}/util/result"
)

//创建
func Create{{.SupStructName}}(data model.{{.SupStructName}})  *result.Response {

	if err := mapper.Create{{.SupStructName}}(data);err!=nil{
        logger.Log.WithFields(logrus.Fields{"err": err, "data": data}).Error("保存异常-{{.SupStructName}}")
        return result.ReturnFailMsg("保存失败")
	}
	return result.ReturnSucNil()
}

//根据ID删除
func Delete{{.SupStructName}}(data model.{{.SupStructName}})  *result.Response {

    if data.ID == 0{
     return result.ReturnFailMsg("参数获取失败")
    }

	if err := mapper.Delete{{.SupStructName}}(data);err!=nil{
            logger.Log.WithFields(logrus.Fields{"err": err, "data": data}).Error("删除异常-{{.SupStructName}}")
            return result.ReturnFailMsg("删除失败")
    }

    return result.ReturnSucNil()
}

//根据ID批量删除
func Delete{{.SupStructName}}ByIds(ids []int64)  *result.Response {
	if len(ids) == 0{
         return result.ReturnFailMsg("参数获取失败")
    }

    if err := mapper.Delete{{.SupStructName}}ByIds(ids);err!=nil{
       logger.Log.WithFields(logrus.Fields{"err": err, "data": ids}).Error("批量删除异常-{{.SupStructName}}")
       return result.ReturnFailMsg("批量删除失败")
    }

    return result.ReturnSucNil()
}

//根据id 更新 ，排除零值
func Update{{.SupStructName}}(data model.{{.SupStructName}})  *result.Response {
	if data.ID == 0{
        return result.ReturnFailMsg("参数获取失败")
    }

    if err := mapper.Update{{.SupStructName}}(data);err!=nil{
        logger.Log.WithFields(logrus.Fields{"err": err, "data": data}).Error("Update异常-{{.SupStructName}}")
        return result.ReturnFailMsg("更新失败")
    }

    return result.ReturnSucNil()
}

//根据id 更新， Save 会保存所有的字段，即使字段是零值 没有记录则insert
func Save{{.SupStructName}}(data model.{{.SupStructName}})  *result.Response {
	if data.ID == 0{
        return result.ReturnFailMsg("参数获取失败")
    }

    if err := mapper.Save{{.SupStructName}}(data);err!=nil{
        logger.Log.WithFields(logrus.Fields{"err": err, "data": data}).Error("save异常-{{.SupStructName}}")
        return result.ReturnFailMsg("更新失败")
    }

    return result.ReturnSucNil()
}

//根据id获取model
func Get{{.SupStructName}}(id int64)  *result.Response {

    if id == 0{
        return result.ReturnFailMsg("参数获取失败")
    }

    if err,data := mapper.Get{{.SupStructName}}(id);err!=nil && !errors.Is(err, gorm.ErrRecordNotFound){
        logger.Log.WithFields(logrus.Fields{"err": err, "data": id}).Error("Get异常-{{.SupStructName}}")
        return result.ReturnFailMsg("获取数据失败")
    }else{
        return result.ReturnSuc(data)
    }

}

//获取所有的model
func GetAll{{.SupStructName}}() *result.Response {

	if err,data := mapper.GetAll{{.SupStructName}}();err!=nil {
        logger.Log.WithFields(logrus.Fields{"err": err}).Error("GetAll异常-{{.SupStructName}}")
        return result.ReturnFailMsg("获取数据失败")
    }else{
        return result.ReturnSuc(data)
    }
}

//按条件分页查询 limit offset ,参数用指针&, 数据会自动填充到req对象
func GetPageLimit{{.SupStructName}}(data model.{{.SupStructName}}Req) *result.Response {

    if err := mapper.GetPageLimit{{.SupStructName}}(&data);err!=nil {
        logger.Log.WithFields(logrus.Fields{"err": err,"data":data}).Error("GetPageLimit异常-{{.SupStructName}}")
        return result.ReturnFailMsg("获取数据失败")
    }else{
        return result.ReturnSuc(data.PageData)
    }

}
