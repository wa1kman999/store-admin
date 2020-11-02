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
func GetDiskList(c *gin.Context) {
	var sp request.SearchDiskParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetDiskInfoList(sp.Disk, sp.PageInfo, sp.OrderKey, sp.Desc)
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
func CreateDisk(c *gin.Context) {
	var disk model.Disk
	_ = c.ShouldBindJSON(&disk)
	DiskVerify := utils.Rules{
		"Vendor":         {utils.NotEmpty()},
		"Series":         {utils.NotEmpty()},
		"ModelNumber":    {utils.NotEmpty()},
		"Sn":             {utils.NotEmpty()},
		"Section":        {utils.NotEmpty()},
		"Capacity":       {utils.NotEmpty()},
		"Interfaces":     {utils.NotEmpty()},
		"Type":           {utils.NotEmpty()},
		"Status":         {utils.NotEmpty()},
		"Region":         {utils.NotEmpty()},
	}
	DiskVerifyErr := utils.Verify(disk, DiskVerify)
	if DiskVerifyErr != nil {
		response.FailWithMessage(DiskVerifyErr.Error(), c)
		return
	}
	err := service.CreateDisk(disk)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//更新硬盘
func UpdateDisk(c *gin.Context) {
	var disk model.Disk
	_ = c.ShouldBindJSON(&disk)
	DiskVerify := utils.Rules{
		"Vendor":         {utils.NotEmpty()},
		"Series":         {utils.NotEmpty()},
		"ModelNumber":    {utils.NotEmpty()},
		"Sn":             {utils.NotEmpty()},
		"Section":        {utils.NotEmpty()},
		"Capacity":       {utils.NotEmpty()},
		"Interfaces":     {utils.NotEmpty()},
		"Type":           {utils.NotEmpty()},
		"Status":         {utils.NotEmpty()},
		"Region":         {utils.NotEmpty()},
	}
	DiskVerifyErr := utils.Verify(disk, DiskVerify)
	if DiskVerifyErr != nil {
		response.FailWithMessage(DiskVerifyErr.Error(), c)
		return
	}
	err := service.UpdateDisk(disk)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

//根据id获取盘
func GetDiskById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, disk := service.GetDiskById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.DiskResponse{Disk: disk}, c)
	}
}

//删除硬盘
func DeleteDisk(c *gin.Context) {
	var a model.Disk
	_ = c.ShouldBindJSON(&a)
	ApiVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.DeleteDisk(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

