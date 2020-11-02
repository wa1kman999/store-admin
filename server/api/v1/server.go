package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// 获取全部
func GetServerList(c *gin.Context) {
	var sp request.SearchServerParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetServerInfoList(sp.Server, sp.PageInfo, sp.OrderKey, sp.Desc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		}, c)
	}
}

// 新增硬盘
func CreateServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindJSON(&server)
	ServerVerify := utils.Rules{
		"Category":         {utils.NotEmpty()},
		"Architecture":         {utils.NotEmpty()},
		//"ModelNumber":    {utils.NotEmpty()},
		"Sn":             {utils.NotEmpty()},
		"Location":        {utils.NotEmpty()},
		//"Capacity":       {utils.NotEmpty()},
		//"Interfaces":     {utils.NotEmpty()},
		//"Type":           {utils.NotEmpty()},
		//"Status":         {utils.NotEmpty()},
		"Region":         {utils.NotEmpty()},
	}
	ServerVerifyErr := utils.Verify(server, ServerVerify)
	if ServerVerifyErr != nil {
		response.FailWithMessage(ServerVerifyErr.Error(), c)
		return
	}
	err := service.CreateServer(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//更新硬盘
func UpdateServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindJSON(&server)
	ServerVerify := utils.Rules{
		"Category":         {utils.NotEmpty()},
		"Architecture":         {utils.NotEmpty()},
		//"ModelNumber":    {utils.NotEmpty()},
		"Sn":             {utils.NotEmpty()},
		"Location":        {utils.NotEmpty()},
		//"Capacity":       {utils.NotEmpty()},
		//"Interfaces":     {utils.NotEmpty()},
		//"Type":           {utils.NotEmpty()},
		"Status":         {utils.NotEmpty()},
		"Region":         {utils.NotEmpty()},
	}
	ServerVerifyErr := utils.Verify(server, ServerVerify)
	if ServerVerifyErr != nil {
		response.FailWithMessage(ServerVerifyErr.Error(), c)
		return
	}
	err := service.UpdateServer(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

//根据id获取盘
func GetServerById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, server := service.GetServerById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.ServerResponse{Server: server}, c)
	}
}

//删除硬盘
func DeleteServer(c *gin.Context) {
	var a model.Server
	_ = c.ShouldBindJSON(&a)
	ApiVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.DeleteServer(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


