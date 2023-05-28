package model

import (
	"time"
)

type CollectInfoType struct {
	Id               int     `json:"id"`
	Collectguid      string  `json:"collectguid"`
	Createtime       string  `json:"createtime"`
	EntityName       string  `json:"entity_name"`
	Typename         string  `json:"typename"`
	Type             string  `json:"type"`
	Created          float64 `json:"created"`
	Updated          float64 `json:"updated"`
	Created2         string  `json:"created2"`
	Updated2         string  `json:"updated2"`
	CreateUserId     int     `json:"createUserId"`
	CreateUserName   string  `json:"createUserName"`
	CreateUserAvatar string  `json:"createUserAvatar"`
	CreateUserRole   string  `json:"createUserRole"`
	UpdateUserId     int     `json:"updateUserId"`
	UpdateUserName   string  `json:"updateUserName"`
	UpdateUserRole   string  `json:"updateUserRole"`
	UpdateUserAvatar string  `json:"updateUserAvatar"`
	Desc             string  `json:"desc"`
}

type UserCollectInfo struct {
	Id          int    `gorm:"id" json:"id"`
	Userid      int    `gorm:"userid" json:"userid"`
	Collectname string `gorm:"collectname" json:"collectname"`
	Touserid    int    `gorm:"touserid" json:"touserid"`
	Collectguid string `gorm:"collectguid" json:"collectguid"`
	Createtime  string `gorm:"createtime" json:"createtime"`
	Description string `gorm:"description" json:"description"`
	Typename    string `gorm:"typename" json:"typename"`
	Findname    string `gorm:"findname" json:"findname"`
	Type        string `gorm:"type" json:"type"`
}

func (this *UserCollectInfo) TableName() string {
	return "UserCollectInfo"
}

type UserCollectRecord struct {
	Id         int    `gorm:"id" json:"id"`
	Userid     int    `gorm:"userid" json:"userid"`
	Collectid  int    `gorm:"collectid" json:"collectid"`
	Touserid   int    `gorm:"touserid" json:"touserid"`
	Createtime string `gorm:"createtime" json:"createtime"`
}

func (this *UserCollectRecord) TableName() string {
	return "UserCollectRecord"
}

type UserCollectRecordInfos struct {
	Id         int    `gorm:"id" json:"id"`
	Userid     int    `gorm:"userid" json:"userid"`
	Content    string `json:"content"`
	Collectid  int    `gorm:"collectid" json:"collectid"`
	Touserid   int    `gorm:"touserid" json:"touserid"`
	TypeInfo   string `json:"type_info"`
	Createtime string `gorm:"createtime" json:"createtime"`
}

func AddCollect(userid, touserid int, guid, name, typename, findname, desc, typeinfo string) {
	newCollectInfo := UserCollectInfo{
		Userid:      userid,
		Collectname: name,
		Touserid:    touserid,
		Collectguid: guid,
		Createtime:  time.Now().Format("2006-01-02 15:04:05"),
		Description: desc,
		Typename:    typename,
		Findname:    findname,
		Type:        typeinfo,
	}
	_ = db.Create(&newCollectInfo)
	newCollectInfoRecord := UserCollectRecord{
		Userid:     userid,
		Collectid:  newCollectInfo.Id,
		Touserid:   touserid,
		Createtime: newCollectInfo.Createtime,
	}
	_ = db.Create(&newCollectInfoRecord)
}

func CheckEntityCollect(userid int, guid string) bool {
	info := UserCollectInfo{}
	db.Where("userid=?", userid).Where("collectguid=?", guid).Find(&info)
	if info.Id != 0 {
		return true
	}
	return false
}

func CheckTypeCollect(userid int, typename string, findname string) bool {
	info := UserCollectInfo{}
	db.Where("userid=?", userid).Where("typename=?", typename).Where("findname=?", findname).Find(&info)
	if info.Id != 0 {
		return true
	}
	return false
}

func GetCollect(userid int) []UserCollectInfo {
	userCollectInfos := []UserCollectInfo{}
	_ = db.Where("userid=?", userid).Find(&userCollectInfos)
	return userCollectInfos
}

func GetCollectRecord(userid int) []UserCollectRecordInfos {
	userCollectRecordInfos := []UserCollectRecordInfos{}
	userCollectInfos := []UserCollectRecord{}
	_ = db.Where("touserid=?", userid).Find(&userCollectInfos)
	for _, info := range userCollectInfos {
		collectInfo := UserCollectInfo{}
		_ = db.Where("id=?", info.Collectid).First(&collectInfo)
		userCollectRecordInfos = append(userCollectRecordInfos, UserCollectRecordInfos{
			Id:         info.Id,
			Userid:     info.Userid,
			Content:    collectInfo.Collectname,
			TypeInfo:   getTypeName(collectInfo.Findname),
			Createtime: info.Createtime,
		})
	}
	return userCollectRecordInfos
}

func getTypeName(name string) string {
	switch name {
	case "glossary":
		return "术语表"
	case "term":
		return "术语"
	case "business":
		return "业务元数据"
	case "classification":
		return "分类类型"
	}
	return "实体"
}

func DeleteCollect(id int) {
	info1 := UserCollectInfo{}
	info2 := UserCollectInfo{}
	db.Where("id=?", id).Find(&info1)
	db.Where("id=?", id).Delete(&info2)
	record := UserCollectRecord{}
	db.Where("collectid=?", info1.Id).Delete(&record)
}
func DeleteEntityCollect(userid int, guid string) {
	info1 := UserCollectInfo{}
	info2 := UserCollectInfo{}
	db.Where("userid=?", userid).Where("collectguid", guid).Find(&info1)
	db.Where("userid=?", userid).Where("collectguid", guid).Delete(&info2)
	record := UserCollectRecord{}
	db.Where("collectid=?", info1.Id).Delete(&record)
}
func DeleteTypeCollect(userid int, typename string, findname string) {
	info1 := UserCollectInfo{}
	info2 := UserCollectInfo{}
	db.Where("userid=?", userid).Where("typename=?", typename).Where("findname=?", findname).Find(&info1)
	db.Where("userid=?", userid).Where("typename=?", typename).Where("findname=?", findname).Delete(&info2)
	record := UserCollectRecord{}
	db.Where("collectid=?", info1.Id).Delete(&record)
}
