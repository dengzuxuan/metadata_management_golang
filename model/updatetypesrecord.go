package model

import "time"

type UpdateTypesRecord struct {
	Id         int    `gorm:"id" json:"id"`
	Action     string `gorm:"action" json:"action"`
	Content    string `gorm:"content" json:"content"`
	Updatetime int64  `gorm:"updatetime" json:"updatetime"`
	Userid     int    `gorm:"userid" json:"userid"`
	Typename   string `gorm:"typename" json:"typename"`
}

func (this *UpdateTypesRecord) TableName() string {
	return "UpdateTypesRecord"
}

func AddTypeRecord(userid int, action string, content string, typename string) {
	newTypeRecord := UpdateTypesRecord{
		Action:     action,
		Content:    content,
		Updatetime: time.Now().Unix(),
		Userid:     userid,
		Typename:   typename,
	}
	_ = db.Create(&newTypeRecord)
}

func GetUserTypeRecord(userid int) []UpdateTypesRecord {
	typeRecord := []UpdateTypesRecord{}
	_ = db.Debug().Where("userid=?", userid).Find(&typeRecord)
	return typeRecord
}
func GetTypeRecord(typename string) []UpdateTypesRecord {
	typeRecord := []UpdateTypesRecord{}
	_ = db.Where("typename=?", typename).Find(&typeRecord)
	return typeRecord
}
