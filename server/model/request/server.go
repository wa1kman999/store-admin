package request


import "gin-vue-admin/model"

// server分页条件查询及排序结构体
type SearchServerParams struct {
	model.Server
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}


