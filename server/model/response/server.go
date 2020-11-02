package response

import "gin-vue-admin/model"

type ServerResponse struct {
	Server     model.Server    `json:"server"`
}
