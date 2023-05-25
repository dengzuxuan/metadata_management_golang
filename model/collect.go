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

func DeleteCollect(id int) {
	info := UserCollectInfo{}
	db.Where("id=?", id).Delete(&info)
	record := UserCollectRecord{}
	db.Where("collectid=?", id).Delete(&record)
}
func DeleteEntityCollect(userid int, guid string) {
	info := UserCollectInfo{}
	db.Where("userid=?", userid).Where("collectguid", guid).Delete(&info)
	record := UserCollectRecord{}
	db.Where("collectid=?", info.Id).Delete(&record)
}
func DeleteTypeCollect(userid int, typename string, findname string) {
	info := UserCollectInfo{}
	db.Where("userid=?", userid).Where("typename=?", typename).Where("findname=?", findname).Delete(&info)
	record := UserCollectRecord{}
	db.Where("collectid=?", info.Id).Delete(&record)
}
