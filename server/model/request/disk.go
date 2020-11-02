package request

import "gin-vue-admin/model"

// api分页条件查询及排序结构体
type SearchDiskParams struct {
	model.Disk
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}

