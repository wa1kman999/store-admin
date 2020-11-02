package response

import "gin-vue-admin/model"

type DiskResponse struct {
	Disk     model.Disk    `json:"disk"`
} 
