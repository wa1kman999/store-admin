package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)
// 获取硬盘列表
func GetDiskInfoList(disk model.Disk, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Disk{})
	var diskList []model.Disk
	if disk.Vendor != "" {
		db = db.Where("vendor LIKE ?", "%"+disk.Vendor+"%")
	}

	if disk.Series != "" {
		db = db.Where("series LIKE ?", "%"+disk.Series+"%")
	}

	if disk.ModelNumber != "" {
		db = db.Where("model_number = ?", disk.ModelNumber)
	}
	if disk.Sn != "" {
		db = db.Where("sn = ?", disk.Sn)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("User").Find(&diskList).Error

	if err != nil {
		return err, diskList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Preload("User").Find(&diskList).Error
		}
		//} else {
		//	err = db.Order("api_group").Find(&diskList).Error
		//}
	}
	return err, diskList, total
	//err = db.Limit(limit).Offset(offset).Preload("User").Find(&diskList).Error
	//return err, diskList, total
}

// 新增硬盘
func CreateDisk(disk model.Disk) (err error) {
	if !errors.Is(global.GVA_DB.Where("vendor = ? AND series = ? AND model_number = ? AND sn = ? AND section = ? AND capacity = ? AND interfaces = ? AND type = ? AND status = ? AND region = ? AND postscript = ?", disk.Vendor, disk.Series, disk.ModelNumber, disk.Sn, disk.Section, disk.Capacity, disk.Interfaces, disk.Type, disk.Status, disk.Region, disk.Postscript).First(&model.Disk{}).Error, gorm.ErrRecordNotFound){
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&disk).Error
}

//更新硬盘
func UpdateDisk(disk model.Disk) (err error) {
	var oldA model.Disk

	err = global.GVA_DB.Where("id = ?", disk.ID).First(&oldA).Error

	if oldA.Vendor != disk.Vendor || oldA.Series != disk.Series || oldA.ModelNumber != disk.ModelNumber {
		if !errors.Is(global.GVA_DB.Where("sn = ?", disk.Sn).First(&model.Disk{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同SN相同的盘")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.GVA_DB.Save(&disk).Error
		}
	return err
}

//根据id获取盘
func GetDiskById(id float64) (err error, disk model.Disk) {
	err = global.GVA_DB.Where("id = ?", id).First(&disk).Error
	return
}

//删除硬盘
func DeleteDisk(disk model.Disk) (err error) {
	err = global.GVA_DB.Delete(disk).Error
	return err
}
