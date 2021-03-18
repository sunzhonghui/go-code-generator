package router

import (
    "{{.ModName}}/util/global"
    "{{.ModName}}/auth/service"
    "{{.ModName}}/{{.Abbr}}/api"
)

func init() {
	// 路由权限相关
	r := global.Router
	v2 := r.Group("{{.RouterName}}")
	{
//		v2.POST("", api.XX) 不走权限
		auth := v2.Group("")
		auth.Use(service.AuthRequired())
		{
			// 需要权限
			auth.POST("create{{.SupStructName}}", api.Create{{.SupStructName}})
			auth.DELETE("Delete{{.SupStructName}}", api.Delete{{.SupStructName}})
			auth.DELETE("Delete{{.SupStructName}}ByIds", api.Delete{{.SupStructName}}ByIds)
			auth.PUT("Update{{.SupStructName}}", api.Update{{.SupStructName}})
			auth.GET("get/:id", api.Get{{.SupStructName}})
			auth.POST("GetPageLimit{{.SupStructName}}", api.GetPageLimit{{.SupStructName}})

		}
	}
}
