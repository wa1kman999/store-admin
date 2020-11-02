package model

import (
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Category     string        `json:"category" gorm:"comment:类别"`
	Ip           string        `json:"ip" gorm:"comment:ip地址"`
	Password     string        `json:"password" gorm:"comment:密码"`
	BusinessIp   string        `json:"business_ip" gorm:"comment:业务ip"`
	BusinessPassword string    `json:"business_password" gorm:"comment:业务ip密码"`
	Sn           string        `json:"sn" gorm:"comment:sn号"`
	Architecture string         `json:"architecture" gorm:"comment:架构"`
	Backboard    string        `json:"backboard" gorm:"comment:背板"`
	RaidCard     string        `json:"raid_card" gorm:"comment:raid卡"`
	Status       int           `json:"status" gorm:"default:1;comment:状态"`
	Region       string        `json:"region" gorm:"comment:地区"`
	Location     string        `json:"location" gorm:"comment:具体位置"`
	Postscript   string        `json:"postscript" gorm:"comment:备注"`
	OccupyTime   string        `json:"occupy_time" gorm:"comment:占用时间"`
	ReleaseTime  string        `json:"release_time" gorm:"comment:释放时间"`
	LastUser     string        `json:"last_user" gorm:"comment:上次用户"`
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User         SysUser       `json:"user"`
}