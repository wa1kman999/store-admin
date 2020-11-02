package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDiskRouter(Router *gin.RouterGroup) {
	DiskRouter := Router.Group("disk").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord()).
		Use(middleware.CasbinHandler())
	{
		DiskRouter.POST("getDiskList", v1.GetDiskList)  // 分页获取硬盘列表
		DiskRouter.POST("createDisk", v1.CreateDisk)    //创建硬盘
		DiskRouter.POST("updateDisk", v1.UpdateDisk)    //更新硬盘信息
		DiskRouter.POST("getDiskById", v1.GetDiskById)  //获取盘通过id
		DiskRouter.POST("deleteDisk", v1.DeleteDisk)    //删除盘
	}
}

