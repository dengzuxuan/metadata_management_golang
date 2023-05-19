package model

import (
	"others-part/utils"
	"time"
)

type ClassificationReqAtlas struct {
	ClassificationDefs []struct {
		Name          string        `json:"name"`
		Description   string        `json:"description"`
		SuperTypes    []interface{} `json:"superTypes"`
		AttributeDefs []struct {
			Name           string `json:"name"`
			Description    string `json:"description"`
			TypeName       string `json:"typeName"`
			IsOptional     bool   `json:"isOptional"`
			Cardinality    string `json:"cardinality"`
			ValuesMinCount int    `json:"valuesMinCount"`
			ValuesMaxCount int    `json:"valuesMaxCount"`
			IsUnique       bool   `json:"isUnique"`
			IsIndexable    bool   `json:"isIndexable"`
		} `json:"attributeDefs"`
	} `json:"classificationDefs"`
	EntityDefs []interface{} `json:"entityDefs"`
	EnumDefs   []interface{} `json:"enumDefs"`
	StructDefs []interface{} `json:"structDefs"`
}
type ClassificationRespAtlas struct {
	ClassificationDefs []struct {
		AttributeDefs []struct {
			Cardinality           string `json:"cardinality"`
			IncludeInNotification bool   `json:"includeInNotification"`
			IsIndexable           bool   `json:"isIndexable"`
			IsOptional            bool   `json:"isOptional"`
			IsUnique              bool   `json:"isUnique"`
			Name                  string `json:"name"`
			SearchWeight          int    `json:"searchWeight"`
			TypeName              string `json:"typeName"`
			ValuesMaxCount        int    `json:"valuesMaxCount"`
			ValuesMinCount        int    `json:"valuesMinCount"`
		} `json:"attributeDefs"`
		Category    string        `json:"category"`
		CreateTime  int64         `json:"createTime"`
		CreatedBy   string        `json:"createdBy"`
		Description string        `json:"description"`
		EntityTypes []interface{} `json:"entityTypes"`
		GUID        string        `json:"guid"`
		Name        string        `json:"name"`
		SubTypes    []interface{} `json:"subTypes"`
		SuperTypes  []interface{} `json:"superTypes"`
		TypeVersion string        `json:"typeVersion"`
		UpdateTime  int64         `json:"updateTime"`
		UpdatedBy   string        `json:"updatedBy"`
		Version     int           `json:"version"`
	} `json:"classificationDefs"`
}
type Addclassification struct {
	Classification struct {
		TypeName   string `json:"typeName"`
		Attributes struct {
		} `json:"attributes"`
		Propagate                        bool          `json:"propagate"`
		RemovePropagationsOnEntityDelete bool          `json:"removePropagationsOnEntityDelete"`
		ValidityPeriods                  []interface{} `json:"validityPeriods"`
	} `json:"classification"`
	EntityGuids []string `json:"entityGuids"`
}
type ClassificationInfo struct {
	Id                 int    `gorm:"id" json:"id"`
	Classificationname string `gorm:"classificationname" json:"classificationname"`
	Userid             int    `gorm:"userid" json:"userid"`
	Username           string `gorm:"username" json:"username"`
	Avatar             string `gorm:"avatar" json:"avatar"`
	Attributenumber    int    `gorm:"attributeenumber" json:"attributenumber"`
	Createtime         string `gorm:"createtime" json:"createtime"`
	Description        string `gorm:"description" json:"description"`
	Guid               string `gorm:"guid" json:"guid"`
}

type ClassificationAttributeInfo struct {
	Id                 int    `gorm:"id" json:"id"`
	Attributename      string `gorm:"attributename" json:"attributename"`
	Classificationname string `gorm:"classificationname" json:"classificationname"`
	Userid             int    `gorm:"userid" json:"userid"`
	Username           string `gorm:"username" json:"username"`
	Avatar             string `gorm:"avatar" json:"avatar"`
	Createtime         string `gorm:"createtime" json:"createtime"`
	Description        string `gorm:"description" json:"description"`
	Guid               string `gorm:"guid" json:"guid"`
}

func (this *ClassificationInfo) TableName() string {
	return "ClassificationInfo"
}
func (this *ClassificationAttributeInfo) TableName() string {
	return "ClassificationAttributeInfo"
}
func AddClassification(classificationName string, userid int, username string, avatar string, attributeenumber int, description string, guid string) int {
	newClassification := ClassificationInfo{
		Classificationname: string(classificationName),
		Userid:             userid,
		Username:           username,
		Avatar:             avatar,
		Attributenumber:    attributeenumber,
		Description:        description,
		Guid:               guid,
		Createtime:         time.Now().Format("2006-01-02 15:05:05"),
	}
	err = db.Create(&newClassification).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}
func GetAllClassification() []ClassificationInfo {
	classification := []ClassificationInfo{}
	_ = db.Find(&classification)
	return classification
}
func AddClassificationAttribute(classificationName string, attributename string, userid int, username string, avatar string, description string, guid string) int {
	newClassificationAttribute := ClassificationAttributeInfo{
		Attributename:      attributename,
		Classificationname: classificationName,
		Userid:             userid,
		Username:           username,
		Avatar:             avatar,
		Description:        description,
		Guid:               guid,
		Createtime:         time.Now().Format("2006-01-02 15:05:05"),
	}
	err = db.Create(&newClassificationAttribute).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}
func GetClassificationAttribute(guid string) []ClassificationAttributeInfo {
	classificationAttribute := []ClassificationAttributeInfo{}
	_ = db.Where("guid=?", guid).Find(&classificationAttribute)
	return classificationAttribute
}
func GetClassificatioInfo(name string) ClassificationInfo {
	classificationInfo := ClassificationInfo{}
	_ = db.Where("Classificationname=?", name).Find(&classificationInfo)
	username, avatar := GetUserInfo(classificationInfo.Userid)
	classificationInfo.Username = username
	classificationInfo.Avatar = avatar
	return classificationInfo
}
func GetClassificationName(guid string) ClassificationInfo {
	classificationAttribute := ClassificationInfo{}
	_ = db.Select("Classificationname").Where("guid=?", guid).Find(&classificationAttribute)
	username, avatar := GetUserInfo(classificationAttribute.Userid)
	classificationAttribute.Username = username
	classificationAttribute.Avatar = avatar
	return classificationAttribute
}
