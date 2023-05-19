package model

import (
	"time"
)

type BusinessMetaInfo struct {
	Id           int    `gorm:"id" json:"id"`
	Businessname string `gorm:"businessname" json:"businessname"`
	Userid       int    `gorm:"userid" json:"userid"`
	Username     string `gorm:"-" json:"username"`
	Avatar       string `gorm:"-" json:"avatar"`
	Createtime   string `gorm:"createtime" json:"createtime"`
	Description  string `gorm:"description" json:"description"`
	Guid         string `gorm:"guid" json:"guid"`
}
type BusinessMetaEntityInfo struct {
	Id             int    `gorm:"id" json:"id"`
	Entityguid     string `gorm:"enetityguid" json:"entityguid"`
	Entityname     string `gorm:"enetityname" json:"entityname"`
	Userid         int    `gorm:"userid" json:"userid"`
	Username       string `gorm:"-" json:"username"`
	Avatar         string `gorm:"-" json:"avatar"`
	Attributename  string `gorm:"attributename" json:"attributename"`
	Attributevalue string `gorm:"attributevalue" json:"attributevalue"`
	Createtime     string `gorm:"createtime" json:"createtime"`
	Guid           string `gorm:"guid" json:"guid"`
}
type BusinessMetaAttributeInfo struct {
	Id            int    `gorm:"id" json:"id"`
	Attributename string `gorm:"attributename" json:"attributename"`
	Businessname  string `gorm:"businessname" json:"businessname"`
	Businessguid  string `gorm:"businessguid" json:"businessguid"`
	Userid        int    `gorm:"userid" json:"userid"`
	Username      string `gorm:"-" json:"username"`
	Avatar        string `gorm:"-" json:"avatar"`
	Description   string `gorm:"description" json:"description"`
	Createtime    string `gorm:"createtime" json:"createtime"`
}
type BusinessMetaAttributeTypeInfo struct {
	Id            int    `gorm:"id" json:"id"`
	Businessname  string `gorm:"businessname" json:"businessname"`
	Attributename string `gorm:"attributename" json:"attributename"`
	Typename      string `gorm:"typename" json:"typename"`
	Userid        int    `gorm:"userid" json:"userid"`
	Username      string `gorm:"-" json:"username"`
	Avatar        string `gorm:"-" json:"avatar"`
	Createtime    string `gorm:"createtime" json:"createtime"`
}

func (this *BusinessMetaInfo) TableName() string {
	return "BusinessMetaInfo"
}
func (this *BusinessMetaEntityInfo) TableName() string {
	return "BusinessMetaEntityInfo"
}
func (this *BusinessMetaAttributeInfo) TableName() string {
	return "BusinessMetaAttributeInfo"
}
func (this *BusinessMetaAttributeTypeInfo) TableName() string {
	return "BusinessMetaAttributeTypeInfo"
}
func AddBusinessMeta(business AtlasBusinessMeta) {
	for _, businessInfo := range business.BusinessMetadataDefs {
		timestamp := int64(businessInfo.CreateTime) / 1000 // 转换为以秒为单位的时间戳
		t := time.Unix(timestamp, 0)
		// 将time.Time对象格式化为指定格式的字符串
		formattedTime := t.Format("2006-01-02 15:04")
		newbusiness := BusinessMetaInfo{
			Businessname: businessInfo.Name,
			Userid:       1,
			Createtime:   formattedTime,
			Description:  businessInfo.Description,
			Guid:         businessInfo.GUID,
		}
		err = db.Create(&newbusiness).Error
		for _, attribute := range businessInfo.AttributeDefs {
			timestamp2 := int64(businessInfo.CreateTime) / 1000 // 转换为以秒为单位的时间戳
			t2 := time.Unix(timestamp2, 0)
			// 将time.Time对象格式化为指定格式的字符串
			formattedTime2 := t2.Format("2006-01-02 15:04")
			newbusinessAttribute := BusinessMetaAttributeInfo{
				Attributename: attribute.Name,
				Businessname:  businessInfo.Name,
				Businessguid:  businessInfo.GUID,
				Userid:        1,
				Description:   "",
				Createtime:    formattedTime2,
			}
			err = db.Create(&newbusinessAttribute).Error
			newbusinessAttributeType := BusinessMetaAttributeTypeInfo{
				Attributename: attribute.Name,
				Businessname:  businessInfo.Name,
				Userid:        1,
				Createtime:    formattedTime2,
			}
			err = db.Create(&newbusinessAttributeType).Error
		}
	}
}

func GetBusinessMeta(guid string) BusinessMetaInfo {
	businessInfo := BusinessMetaInfo{}
	_ = db.Where("guid=?", guid).Find(&businessInfo)
	username, avatar := GetUserInfo(businessInfo.Userid)
	businessInfo.Username = username
	businessInfo.Avatar = avatar
	return businessInfo
}
func GetBusinessMetaAttribute(guid string) []BusinessMetaAttributeInfo {
	businessInfoAttribute := []BusinessMetaAttributeInfo{}
	businessInfoAttributes := []BusinessMetaAttributeInfo{}
	_ = db.Where("businessguid=?", guid).Find(&businessInfoAttribute)
	for _, info := range businessInfoAttribute {
		username, avatar := GetUserInfo(info.Userid)
		info.Username = username
		info.Avatar = avatar
		businessInfoAttributes = append(businessInfoAttributes, info)
	}
	return businessInfoAttributes
}
