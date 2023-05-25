package model

import "time"

type UpdateTypesRecord struct {
	Id         int    `gorm:"id" json:"id"`
	Action     string `gorm:"action" json:"action"`
	Content    string `gorm:"content" json:"content"`
	Updatetime int64  `gorm:"updatetime" json:"updatetime"`
	Userid     int    `gorm:"userid" json:"userid"`
}

func (this *UpdateTypesRecord) TableName() string {
	return "UpdateTypesRecord"
}

func AddTypeRecord(userid int, action string, content string) {
	newTypeRecord := UpdateTypesRecord{
		Action:     action,
		Content:    content,
		Updatetime: time.Now().Unix(),
		Userid:     userid,
	}
	_ = db.Create(&newTypeRecord)
}
