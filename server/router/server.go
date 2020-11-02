package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("server").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
		//Use(middleware.CasbinHandler())
	{
		ServerRouter.POST("getServerList", v1.GetServerList)  // 分页获取服务器列表
		ServerRouter.POST("createServer", v1.CreateServer)    //创建服务器
		ServerRouter.POST("updateServer", v1.UpdateServer)    //更新服务器信息
		ServerRouter.POST("getServerById", v1.GetServerById)  //获取服务器通过id
		ServerRouter.POST("deleteServer", v1.DeleteServer)    //删除服务器
	}
}

