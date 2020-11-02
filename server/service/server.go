package service


import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)
// 获取server列表
func GetServerInfoList(server model.Server, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Server{})
	var ServerList []model.Server
	if server.Category != "" {
		db = db.Where("category LIKE ?", "%"+server.Category+"%")
	}

	//if server.Status != "" {
	//	db = db.Where("server LIKE ?", "%"+disk.Series+"%")
	//}

	if server.Status != 0 {
		db = db.Where("status = ?", server.Status)
	}
	if server.Ip != "" {
		db = db.Where("ip = ?", server.Ip)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("User").Find(&ServerList).Error

	if err != nil {
		return err, ServerList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Preload("User").Find(&ServerList).Error
		}
		//} else {
		//	err = db.Order("api_group").Find(&diskList).Error
		//}
	}
	return err, ServerList, total
	//err = db.Limit(limit).Offset(offset).Preload("User").Find(&diskList).Error
	//return err, diskList, total
}

// 新增server
func CreateServer(server model.Server) (err error) {
	if !errors.Is(global.GVA_DB.Where("category = ? AND sn = ? AND architecture = ? ", server.Category, server.Sn, server.Architecture).First(&model.Server{}).Error, gorm.ErrRecordNotFound){
		return errors.New("存在相同服务器")
	}
	return global.GVA_DB.Create(&server).Error
}

//更新server
func UpdateServer(server model.Server) (err error) {
	var oldA model.Server

	err = global.GVA_DB.Where("id = ?", server.ID).First(&oldA).Error

	if oldA.Category != server.Category {
		if !errors.Is(global.GVA_DB.Where(" = ?", server.Ip).First(&model.Server{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同ip相同的服务器")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.GVA_DB.Save(&server).Error
	}
	return err
}

//根据id获取server
func GetServerById(id float64) (err error, server model.Server) {
	err = global.GVA_DB.Where("id = ?", id).First(&server).Error
	return
}

//删除server
func DeleteServer(server model.Server) (err error) {
	err = global.GVA_DB.Delete(server).Error
	return err
}
