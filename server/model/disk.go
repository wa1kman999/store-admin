package model

import "gorm.io/gorm"

type Disk struct {
	gorm.Model
	Vendor       string        `json:"vendor" gorm:"comment:供应商"`
	Series       string        `json:"series" gorm:"comment:系列"`
	ModelNumber  string        `json:"model_number" gorm:"comment:model号"`
	Sn           string        `json:"sn" gorm:"comment:sn号"`
	Section      string        `json:"section" gorm:"comment:扇区"`
	Capacity     string        `json:"capacity" gorm:"comment:容量"`
	Interfaces   int           `json:"interfaces" gorm:"comment:接口"`
	Type         int           `json:"type" gorm:"comment:类型"`
	Status       int           `json:"status" gorm:"comment:状态"`
	Region       string        `json:"region" gorm:"comment:地区"`
	//Healthy      bool          `json:"healthy" gorm:"comment:健康"`
	Postscript   string        `json:"postscript" gorm:"comment:备注"`
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User         SysUser       `json:"user"`

}
