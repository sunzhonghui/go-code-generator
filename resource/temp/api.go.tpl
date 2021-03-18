package api

import (
    "strconv"
	"github.com/gin-gonic/gin"
    "{{.ModName}}/util/global"
    "{{.ModName}}/util/result"
    "{{.ModName}}/{{.Abbr}}/model"
	"{{.ModName}}/{{.Abbr}}/service"
)

// @Tags {{.SupStructName}}
// @Summary 创建{{.SupStructName}}
// @accept application/json
// @Produce application/json
// @Param data body model.{{.SupStructName}} true "创建{{.SupStructName}}"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"操作成功"}"
// @Router /{{.RouterName}}/create{{.SupStructName}} [post]
func Create{{.SupStructName}}(c *gin.Context) {

	var req model.{{.SupStructName}}

    if err := c.ShouldBindJSON(&req);err != nil {
		c.JSON(200, result.ReturnFailMsg("获取参数失败"))
	} else {
		res := service.Create{{.SupStructName}}(req)
		c.JSON(200, res)
	}
}

// @Tags {{.SupStructName}}
// @Summary 删除{{.SupStructName}}
// @accept application/json
// @Produce application/json
// @Param data body model.{{.SupStructName}} true "删除{{.SupStructName}}"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"删除成功"}"
// @Router /{{.RouterName}}/delete{{.SupStructName}} [delete]
func Delete{{.SupStructName}}(c *gin.Context) {

    var req model.{{.SupStructName}}

    if err := c.ShouldBindJSON(&req);err != nil {
		c.JSON(200, result.ReturnFailMsg("获取参数失败"))
	} else {
		res := service.Delete{{.SupStructName}}(req)
		c.JSON(200, res)
	}

}

// @Tags {{.SupStructName}}
// @Summary 批量删除{{.SupStructName}}
// @accept application/json
// @Produce application/json
// @Param data body global.IdsReq true "批量删除{{.SupStructName}}"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"删除成功"}"
// @Router /{{.RouterName}}/delete{{.SupStructName}}ByIds [delete]
func Delete{{.SupStructName}}ByIds(c *gin.Context) {

    var req global.IdsReq

    if err := c.ShouldBindJSON(&req);err != nil {
		c.JSON(200, result.ReturnFailMsg("获取参数失败"))
	} else {
		res := service.Delete{{.SupStructName}}ByIds(req.Ids)
		c.JSON(200, res)
	}
}

// @Tags {{.SupStructName}}
// @Summary 更新{{.SupStructName}}
// @accept application/json
// @Produce application/json
// @Param data body model.{{.SupStructName}} true "更新{{.SupStructName}}"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"更新成功"}"
// @Router /{{.RouterName}}/update{{.SupStructName}} [put]
func Update{{.SupStructName}}(c *gin.Context) {
    var req model.{{.SupStructName}}

    if err := c.ShouldBindJSON(&req);err != nil {
		c.JSON(200, result.ReturnFailMsg("获取参数失败"))
	} else {
		res := service.Update{{.SupStructName}}(req)
		c.JSON(200, res)
	}
}

// @Tags {{.SupStructName}}
// @Summary 用id查询{{.SupStructName}}
// @accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"查询成功"}"
// @Router /{{.RouterName}}/get/{id} [get]
func Get{{.SupStructName}}(c *gin.Context) {

    idstr:=c.Param("id") //查询路径Path参数
    id, err := strconv.ParseInt(idstr, 10, 64)
    if err!=nil{
        c.JSON(200, result.ReturnFailMsg("获取参数失败"))
    }else{
        res := service.Get{{.SupStructName}}(id)
    	c.JSON(200, res)
    }

}

// @Tags {{.SupStructName}}
// @Summary 分页获取{{.SupStructName}}列表
// @accept application/json
// @Produce application/json
// @Param data body model.{{.SupStructName}}Req true "分页获取{{.SupStructName}}列表"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"获取成功"}"
// @Router /{{.RouterName}}/get{{.SupStructName}}List [post]
func GetPageLimit{{.SupStructName}}(c *gin.Context) {

    var req model.{{.SupStructName}}Req

    if err := c.ShouldBindJSON(&req);err != nil {
		c.JSON(200, result.ReturnFailMsg("获取参数失败"))
	} else {
		res := service.GetPageLimit{{.SupStructName}}(req)
		c.JSON(200, res)
	}
}
