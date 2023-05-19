package model

import (
	"others-part/utils"
	"time"
)

type EntityTypeInfo struct {
	Id          int    `gorm:"id" json:"id"`
	Typename    string `gorm:"typename" json:"typename"`
	Userid      int    `gorm:"userid" json:"userid"`
	Username    string `gorm:"-" json:"username"`
	Avatar      string `gorm:"-" json:"avatar"`
	Createtime  string `gorm:"createtime" json:"createtime"`
	Description string `gorm:"description" json:"description"`
}

func (this *EntityTypeInfo) TableName() string {
	return "EntityTypeInfo"
}

func AddEntityType(typename string) int {
	newEntity := EntityTypeInfo{
		Typename:    typename,
		Userid:      1,
		Createtime:  time.Now().Add(-1 * time.Hour * 24 * 3).Format("2006-01-02 15:05:05"),
		Description: "",
	}
	err = db.Create(&newEntity).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}

func GetEntityType(name string) EntityTypeInfo {
	entityInfo := EntityTypeInfo{}
	_ = db.Where("Typename=?", name).Find(&entityInfo)
	username, avatar := GetUserInfo(entityInfo.Userid)
	entityInfo.Username = username
	entityInfo.Avatar = avatar
	return entityInfo
}
