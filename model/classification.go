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
type ClassificationUpdateAtlas struct {
	ClassificationDefs []ClassificationDef `json:"classificationDefs"`
	EntityDefs         []interface{}       `json:"entityDefs"`
	EnumDefs           []interface{}       `json:"enumDefs"`
	StructDefs         []interface{}       `json:"structDefs"`
}
type ClassificationDef struct {
	Category      string             `json:"category"`
	GUID          string             `json:"guid"`
	CreatedBy     string             `json:"createdBy"`
	UpdatedBy     string             `json:"updatedBy"`
	CreateTime    int64              `json:"createTime"`
	UpdateTime    int64              `json:"updateTime"`
	Version       int                `json:"version"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	TypeVersion   string             `json:"typeVersion"`
	AttributeDefs []AttributeDefType `json:"attributeDefs"`
	SuperTypes    []interface{}      `json:"superTypes"`
	EntityTypes   []interface{}      `json:"entityTypes"`
	SubTypes      []interface{}      `json:"subTypes"`
}
type AttributeDefType struct {
	Name                  string `json:"name"`
	TypeName              string `json:"typeName"`
	IsOptional            bool   `json:"isOptional"`
	Cardinality           string `json:"cardinality"`
	ValuesMinCount        int    `json:"valuesMinCount"`
	ValuesMaxCount        int    `json:"valuesMaxCount"`
	IsUnique              bool   `json:"isUnique"`
	IsIndexable           bool   `json:"isIndexable"`
	IncludeInNotification bool   `json:"includeInNotification"`
	SearchWeight          int    `json:"searchWeight"`
}
type UpdateAttribue []struct {
	ID             int    `json:"id"`
	TypeName       string `json:"typeName"`
	IsOptional     bool   `json:"isOptional"`
	Cardinality    string `json:"cardinality"`
	ValuesMinCount int    `json:"valuesMinCount"`
	ValuesMaxCount int    `json:"valuesMaxCount"`
	IsUnique       bool   `json:"isUnique"`
	IsIndexable    bool   `json:"isIndexable"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

type AddClassificationReq []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	Username           string `gorm:"-" json:"username"`
	Avatar             string `gorm:"-" json:"avatar"`
	Attributenumber    int    `gorm:"attributeenumber" json:"attributenumber"`
	Createtime         string `gorm:"createtime" json:"createtime"`
	Description        string `gorm:"description" json:"description"`
	Guid               string `gorm:"guid" json:"guid"`
	Updatetime         string `gorm:"updatetime" json:"updatetime"`
	Version            int    `json:"version"`
	Updateuserid       int    `gorm:"updateuserid" json:"updateuserid"`
}

type ClassificationAttributeInfo struct {
	Id                 int    `gorm:"id" json:"id"`
	Attributename      string `gorm:"attributename" json:"attributename"`
	Classificationname string `gorm:"classificationname" json:"classificationname"`
	Userid             int    `gorm:"userid" json:"userid"`
	Username           string `gorm:"-" json:"username"`
	Avatar             string `gorm:"-" json:"avatar"`
	Createtime         string `gorm:"createtime" json:"createtime"`
	Description        string `gorm:"description" json:"description"`
	Guid               string `gorm:"guid" json:"guid"`
	Attributeguid      string `gorm:"attributeguid" json:"attributeguid"`
	Updatetime         string `gorm:"updatetime" json:"updatetime"`
	Updateuserid       int    `gorm:"updateuserid" json:"updateuserid"`
	UpdateUsername     string `gorm:"-" json:"update_username"`
	UpdateAvatar       string `gorm:"-" json:"update_avatar"`
}

func (this *ClassificationInfo) TableName() string {
	return "ClassificationInfo"
}
func (this *ClassificationAttributeInfo) TableName() string {
	return "ClassificationAttributeInfo"
}
func AddClassification(classificationName string, userid int, attributeenumber int, version int, description string, guid string) int {
	newClassification := ClassificationInfo{
		Classificationname: string(classificationName),
		Userid:             userid,
		Attributenumber:    attributeenumber,
		Description:        description,
		Guid:               guid,
		Version:            version,
		Createtime:         time.Now().Format("2006-01-02 15:04:05"),
	}
	err = db.Create(&newClassification).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}

func UpdateClassificationVersion(classificationName string, userid int, version int) int {
	classificationInfo := ClassificationInfo{}
	_ = db.Model(&classificationInfo).Where("classificationname=?", classificationName).Update("version", version)
	_ = db.Model(&classificationInfo).Where("classificationname=?", classificationName).Update("updateuserid", userid)
	_ = db.Model(&classificationInfo).Where("classificationname=?", classificationName).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
	return utils.SUCCESS
}

func UpdateClassification(name string, desc string, userid int) {
	classificationInfo := ClassificationInfo{}
	_ = db.Model(&classificationInfo).Where("classificationname=?", name).Update("description", desc)
	_ = db.Model(&classificationInfo).Where("classificationname=?", name).Update("updateuserid", userid)
	_ = db.Model(&classificationInfo).Where("classificationname=?", name).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
}
func CheckClassificationAttribute(attributename string, classificationname string, userid int, description string, guid string) string {
	attribute := ClassificationAttributeInfo{}
	_ = db.Where("attributename=?", attributename).Where("classificationname=?", classificationname).Find(&attribute)
	if attribute.Id == 0 {
		newClassificationAttribute := ClassificationAttributeInfo{
			Attributename:      attributename,
			Classificationname: classificationname,
			Userid:             userid,
			Description:        description,
			Guid:               guid,
			Createtime:         time.Now().Format("2006-01-02 15:04:05"),
		}
		_ = db.Create(&newClassificationAttribute)
		return "Create Attribute"
	} else {
		if attribute.Description != description {
			_ = db.Model(&attribute).Where("attributename=?", attributename).Where("classificationname=?", classificationname).
				Update("description", description).
				Update("updateuserid", userid).
				Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
			return "Update Attribute"
		}
	}
	return "none"
}
func GetAddAttributesAtlasReq(atts AddClassificationReq, username string, classificationname string) ClassificationUpdateAtlas {
	attrTypes := []AttributeDefType{}
	oriAttribute := GetAttribute(classificationname)
	for _, info := range oriAttribute {
		attrTypes = append(attrTypes, AttributeDefType{
			Name:                  info.Attributename,
			TypeName:              "string",
			IsOptional:            true,
			Cardinality:           "SINGLE",
			ValuesMinCount:        0,
			ValuesMaxCount:        1,
			IsUnique:              false,
			IsIndexable:           true,
			IncludeInNotification: false,
			SearchWeight:          -1,
		})
	}
	for _, info := range atts {
		attrTypes = append(attrTypes, AttributeDefType{
			Name:                  info.Name,
			TypeName:              "string",
			IsOptional:            true,
			Cardinality:           "SINGLE",
			ValuesMinCount:        0,
			ValuesMaxCount:        1,
			IsUnique:              false,
			IsIndexable:           true,
			IncludeInNotification: false,
			SearchWeight:          -1,
		})
	}
	oriInfo := GetClassificatioInfo(classificationname)
	classifications := ClassificationDef{
		Category:      "CLASSIFICATION",
		GUID:          oriInfo.Guid,
		CreatedBy:     oriInfo.Username,
		UpdatedBy:     username,
		CreateTime:    utils.TimeStringToUnix(oriInfo.Createtime),
		UpdateTime:    time.Now().Unix(),
		Version:       oriInfo.Version,
		Name:          classificationname,
		Description:   oriInfo.Description,
		TypeVersion:   "1.0",
		AttributeDefs: attrTypes,
		SuperTypes:    nil,
		EntityTypes:   nil,
		SubTypes:      nil,
	}
	classificationInfos := ClassificationUpdateAtlas{
		ClassificationDefs: []ClassificationDef{classifications},
		EntityDefs:         nil,
		EnumDefs:           nil,
		StructDefs:         nil,
	}
	return classificationInfos
}

func GetAllClassification() []ClassificationInfo {
	classification := []ClassificationInfo{}
	_ = db.Find(&classification)
	return classification
}
func AddClassificationAttribute(classificationName string, attributename string, userid int, description string, guid string) int {
	newClassificationAttribute := ClassificationAttributeInfo{
		Attributename:      attributename,
		Classificationname: classificationName,
		Userid:             userid,
		Description:        description,
		Guid:               guid,
		Createtime:         time.Now().Format("2006-01-02 15:04:05"),
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

type ClassificationInfoType struct {
	Id                 int                           `gorm:"id" json:"id"`
	Version            int                           `gorm:"version" json:"version"`
	Classificationname string                        `gorm:"classificationname" json:"classificationname"`
	Userid             int                           `gorm:"userid" json:"userid"`
	Username           string                        `gorm:"-" json:"username"`
	Avatar             string                        `gorm:"-" json:"avatar"`
	Attributenumber    int                           `gorm:"attributeenumber" json:"attributenumber"`
	Createtime         string                        `gorm:"createtime" json:"createtime"`
	Description        string                        `gorm:"description" json:"description"`
	Guid               string                        `gorm:"guid" json:"guid"`
	Updatetime         string                        `gorm:"updatetime" json:"updatetime"`
	Updateuserid       int                           `gorm:"updateuserid" json:"updateuserid"`
	AttributeInfos     []ClassificationAttributeInfo `json:"attribute_infos"`
	UpdateUsername     string                        `gorm:"-" json:"update_username"`
	UpdateAvatar       string                        `gorm:"-" json:"update_avatar"`
}
type ClassificationAttributeUpdateInfo struct {
	Id                 int    `json:"id"`
	Attributename      string `json:"attributename"`
	Classificationname string `json:"classificationname"`
	Userid             int    `json:"userid"`
	Username           string `json:"username"`
	Avatar             string `json:"avatar"`
	Createtime         string `json:"createtime"`
	Description        string `json:"description"`
	Guid               string `json:"guid"`
	Attributeguid      string `json:"attributeguid"`
	Updatetime         string `json:"updatetime"`
	Updateuserid       int    `json:"updateuserid"`
	UpdateUsername     string `json:"update_username"`
	UpdateAvatar       string `json:"update_avatar"`
}

func GetClassificatioInfo(name string) ClassificationInfoType {
	classificationInfo := ClassificationInfo{}
	_ = db.Where("classificationname=?", name).Find(&classificationInfo)
	attributeInfos2 := []ClassificationAttributeInfo{}
	attributeInfos := []ClassificationAttributeInfo{}
	_ = db.Debug().Where("classificationname=?", name).Find(&attributeInfos)
	for _, info := range attributeInfos {
		info.Username, info.Avatar = GetUserInfo(info.Userid)
		info.UpdateUsername, info.UpdateAvatar = GetUserInfo(info.Updateuserid)
		attributeInfos2 = append(attributeInfos2, info)
	}
	username, avatar := GetUserInfo(classificationInfo.Userid)
	updateusername, updateavatar := GetUserInfo(classificationInfo.Updateuserid)
	classificationInfos := ClassificationInfoType{
		Id:                 classificationInfo.Id,
		Version:            classificationInfo.Version,
		Classificationname: classificationInfo.Classificationname,
		Userid:             classificationInfo.Userid,
		Username:           username,
		Avatar:             avatar,
		Attributenumber:    classificationInfo.Attributenumber,
		Createtime:         classificationInfo.Createtime,
		Description:        classificationInfo.Description,
		Guid:               classificationInfo.Guid,
		Updatetime:         classificationInfo.Updatetime,
		Updateuserid:       classificationInfo.Updateuserid,
		AttributeInfos:     attributeInfos2,
		UpdateAvatar:       updateavatar,
		UpdateUsername:     updateusername,
	}
	return classificationInfos
}

func GetAttributeInfo(name string, attributename string) ClassificationAttributeInfo {
	attributes := ClassificationAttributeInfo{}
	_ = db.Where("classificationname=?", name).Where("attributename=?", attributename).Find(&attributes)
	return attributes
}

func UpdateAttribute(name string, attributename string, description string, userid int) {
	newClassificationAttribute := ClassificationAttributeInfo{
		Attributename:      attributename,
		Classificationname: name,
	}
	db.Find(&newClassificationAttribute).Update("description", description)
	db.Find(&newClassificationAttribute).Update("updateuserid", userid)
	db.Find(&newClassificationAttribute).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))

}
func GetAttribute(name string) []ClassificationAttributeInfo {
	attributes := []ClassificationAttributeInfo{}
	_ = db.Where("classificationname=?", name).Find(&attributes)
	return attributes
}

func GetClassificationName(guid string) ClassificationInfo {
	classificationAttribute := ClassificationInfo{}
	_ = db.Select("classificationname").Where("guid=?", guid).Find(&classificationAttribute)
	username, avatar := GetUserInfo(classificationAttribute.Userid)
	classificationAttribute.Username = username
	classificationAttribute.Avatar = avatar
	return classificationAttribute
}
