package middleware

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		fmt.Println(claims,111)
		waitUse := claims.(*request.CustomClaims)
		fmt.Println(waitUse,222)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		fmt.Println(obj,333)
		// 获取请求方法
		act := c.Request.Method
		fmt.Println(act,444)
		// 获取用户的角色
		sub := waitUse.AuthorityId
		fmt.Println(sub,555)
		e := service.Casbin()
		// 判断策略中是否存在
		fmt.Println(e,666)
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.Result(response.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
