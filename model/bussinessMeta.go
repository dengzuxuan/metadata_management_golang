package model

import (
	"encoding/json"
	"fmt"
	"others-part/utils"
	"time"
)

type BMRespAtlas struct {
	EnumDefs             []interface{} `json:"enumDefs"`
	StructDefs           []interface{} `json:"structDefs"`
	ClassificationDefs   []interface{} `json:"classificationDefs"`
	EntityDefs           []interface{} `json:"entityDefs"`
	RelationshipDefs     []interface{} `json:"relationshipDefs"`
	BusinessMetadataDefs []struct {
		Category      string `json:"category"`
		GUID          string `json:"guid"`
		CreatedBy     string `json:"createdBy"`
		UpdatedBy     string `json:"updatedBy"`
		CreateTime    int64  `json:"createTime"`
		UpdateTime    int64  `json:"updateTime"`
		Version       int    `json:"version"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		TypeVersion   string `json:"typeVersion"`
		AttributeDefs []struct {
			Name                  string `json:"name"`
			Desc                  string `json:"desc"`
			TypeName              string `json:"typeName"`
			IsOptional            bool   `json:"isOptional"`
			Cardinality           string `json:"cardinality"`
			ValuesMinCount        int    `json:"valuesMinCount"`
			ValuesMaxCount        int    `json:"valuesMaxCount"`
			IsUnique              bool   `json:"isUnique"`
			IsIndexable           bool   `json:"isIndexable"`
			IncludeInNotification bool   `json:"includeInNotification"`
			SearchWeight          int    `json:"searchWeight"`
			Options               struct {
				ApplicableEntityTypes string `json:"applicableEntityTypes"`
				MaxStrLength          string `json:"maxStrLength"`
			} `json:"options"`
		} `json:"attributeDefs"`
	} `json:"businessMetadataDefs"`
}
type BusinessAddAtlas struct {
	EnumDefs             []interface{} `json:"enumDefs"`
	StructDefs           []interface{} `json:"structDefs"`
	ClassificationDefs   []interface{} `json:"classificationDefs"`
	EntityDefs           []interface{} `json:"entityDefs"`
	RelationshipDefs     []interface{} `json:"relationshipDefs"`
	BusinessMetadataDefs []BusinessDef `json:"businessMetadataDefs"`
}
type BusinessDef struct {
	Category      string         `json:"category"`
	GUID          string         `json:"guid"`
	CreatedBy     string         `json:"createdBy"`
	UpdatedBy     string         `json:"updatedBy"`
	CreateTime    int64          `json:"createTime"`
	UpdateTime    int64          `json:"updateTime"`
	Version       int            `json:"version"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	TypeVersion   string         `json:"typeVersion"`
	AttributeDefs []AttributeDef `json:"attributeDefs"`
}
type AttributeDef struct {
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
	Options               struct {
		ApplicableEntityTypes string `json:"applicableEntityTypes"`
		MaxStrLength          string `json:"maxStrLength"`
	} `json:"options"`
}
type BusinessMetaInfo struct {
	Id             int    `gorm:"id" json:"id"`
	Businessname   string `gorm:"businessname" json:"businessname"`
	Userid         int    `gorm:"userid" json:"userid"`
	Username       string `gorm:"-" json:"username"`
	Avatar         string `gorm:"-" json:"avatar"`
	Createtime     string `gorm:"createtime" json:"createtime"`
	Description    string `gorm:"description" json:"description"`
	Guid           string `gorm:"guid" json:"guid"`
	Version        int    `gorm:"version" json:"version"`
	Updatetime     string `gorm:"updatetime" json:"updatetime"`
	Updateuserid   int    `gorm:"updateuserid" json:"updateuserid"`
	UpdateUsername string `gorm:"-" json:"update_username"`
	UpdateAvatar   string `gorm:"-" json:"update_avatar"`
}
type BusinessMetaEntityInfo struct {
	Id             int    `gorm:"id" json:"id"`
	Entityguid     string `gorm:"entityguid" json:"entityguid"`
	Entityname     string `gorm:"entityname" json:"entityname"`
	Userid         int    `gorm:"userid" json:"userid"`
	Username       string `gorm:"-" json:"username"`
	Avatar         string `gorm:"-" json:"avatar"`
	Attributename  string `gorm:"attributename" json:"attributename"`
	Attributevalue string `gorm:"attributevalue" json:"attributevalue"`
	Createtime     string `gorm:"createtime" json:"createtime"`
	Guid           string `gorm:"guid" json:"guid"`
	Typename       string `gorm:"typename" json:"typename"`
	Businessname   string `gorm:"businessname" json:"businessname"`
}
type BusinessMetaAttributeInfo struct {
	Id             int      `gorm:"id" json:"id"`
	Attributename  string   `gorm:"attributename" json:"attributename"`
	Businessname   string   `gorm:"businessname" json:"businessname"`
	Businessguid   string   `gorm:"businessguid" json:"businessguid"`
	Userid         int      `gorm:"userid" json:"userid"`
	Username       string   `gorm:"-" json:"username"`
	Avatar         string   `gorm:"-" json:"avatar"`
	Guid           string   `gorm:"guid" json:"guid"`
	Description    string   `gorm:"description" json:"description"`
	Createtime     string   `gorm:"createtime" json:"createtime"`
	Attributeguid  string   `gorm:"attributeguid" json:"attributeguid"`
	Updatetime     string   `gorm:"updatetime" json:"updatetime"`
	Updateuserid   int      `gorm:"updateuserid" json:"updateuserid"`
	UpdateUsername string   `gorm:"-" json:"update_username"`
	UpdateAvatar   string   `gorm:"-" json:"update_avatar"`
	Weight         int      `gorm:"weight" json:"weight"`
	Types          []string `gorm:"-" json:"types"`
}
type BusinessMetaAttributeTypeInfo struct {
	Id            int    `gorm:"id" json:"id"`
	Businessname  string `gorm:"businessname" json:"businessname"`
	Desc          string `gorm:"-" json:"desc"`
	Attributename string `gorm:"attributename" json:"attributename"`
	Typename      string `gorm:"typename" json:"typename"`
	Userid        int    `gorm:"userid" json:"userid"`
	Username      string `gorm:"-" json:"username"`
	Avatar        string `gorm:"-" json:"avatar"`
	Guid          string `gorm:"guid" json:"guid"`
	Createtime    string `gorm:"createtime" json:"createtime"`
}
type BusinessMetaInfoType struct {
	Id             int                         `gorm:"id" json:"id"`
	Businessname   string                      `gorm:"businessname" json:"businessname"`
	Userid         int                         `gorm:"userid" json:"userid"`
	Username       string                      `gorm:"-" json:"username"`
	Avatar         string                      `gorm:"-" json:"avatar"`
	Createtime     string                      `gorm:"createtime" json:"createtime"`
	Description    string                      `gorm:"description" json:"description"`
	Guid           string                      `gorm:"guid" json:"guid"`
	Updatetime     string                      `gorm:"updatetime" json:"updatetime"`
	Updateuserid   int                         `gorm:"updateuserid" json:"updateuserid"`
	AttributeInfos []BusinessMetaAttributeInfo `json:"attribute_infos"`
	UpdateUsername string                      `gorm:"-" json:"update_username"`
	UpdateAvatar   string                      `gorm:"-" json:"update_avatar"`
	Version        int                         `gorm:"-" json:"version"`
}

// todo:把这些信息都附加上一个人的头像，意味着是这个人操作的
type guidBusinessInfo struct {
	Guid           string              `json:"guid"`
	UserId         int                 `json:"user_id"`
	BusinessName   string              `json:"business_name"`
	BusinessDesc   string              `json:"business_desc"`
	CreateTime     string              `json:"create_time"`
	AttributeInfos []guidAttributeInfo `json:"attribute_infos"`
}

type guidAttributeInfo struct {
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
	AttributeDesc  string `json:"attribute_desc"`
	UserId         int    `json:"user_id"`
	CreateTime     string `json:"create_time"`
}

type guidAttributeLists struct {
	Guid          string `json:"guid"`
	BusinessName  string `json:"business_name"`
	BusinessDesc  string `json:"business_desc"`
	AttributeName string `json:"attribute_name"`
	AttributeDesc string `json:"attribute_desc"`
	IsExist       bool   `json:"is_exist"`
}

type OriBmReq []struct {
	GUID           string `json:"guid"`
	UserID         int    `json:"user_id"`
	BusinessName   string `json:"business_name"`
	BusinessDesc   string `json:"business_desc"`
	CreateTime     string `json:"create_time"`
	AttributeInfos []struct {
		AttributeName  string `json:"attribute_name"`
		AttributeValue string `json:"attribute_value"`
		AttributeDesc  string `json:"attribute_desc"`
		UserID         int    `json:"user_id"`
		CreateTime     string `json:"create_time"`
	} `json:"attribute_infos"`
}
type AddBmReq []struct {
	BusinessName  string `json:"business_name"`
	AttributeName struct {
		ID            int    `json:"id"`
		Disabled      bool   `json:"disabled"`
		Value         string `json:"value"`
		Label         string `json:"label"`
		BusinessName  string `json:"business_name"`
		AttributeName string `json:"attribute_name"`
	} `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}
type BusinessAttributeAdd []BusinessAttribueAddItem
type BusinessAttribueAddItem struct {
	ID     int      `json:"id"`
	Weight int      `json:"weight"`
	Types  []string `json:"types"`
	Name   string   `json:"name"`
	Desc   string   `json:"desc"`
}

type BusinessAttributeUpdate struct {
	ID             int      `json:"id"`
	Attributename  string   `json:"attributename"`
	Businessname   string   `json:"businessname"`
	Businessguid   string   `json:"businessguid"`
	Userid         int      `json:"userid"`
	Username       string   `json:"username"`
	Avatar         string   `json:"avatar"`
	GUID           string   `json:"guid"`
	Description    string   `json:"description"`
	Createtime     string   `json:"createtime"`
	Attributeguid  string   `json:"attributeguid"`
	Updatetime     string   `json:"updatetime"`
	Updateuserid   int      `json:"updateuserid"`
	UpdateUsername string   `json:"update_username"`
	UpdateAvatar   string   `json:"update_avatar"`
	Weight         int      `json:"weight"`
	Types          []string `json:"types"`
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
		formattedTime := t.Format("2006-01-02 15:04:05")
		newbusiness := BusinessMetaInfo{
			Businessname: businessInfo.Name,
			Userid:       GetUserId(businessInfo.CreatedBy),
			Createtime:   formattedTime,
			Description:  businessInfo.Description,
			Guid:         businessInfo.GUID,
		}
		err = db.Create(&newbusiness).Error
		for _, attribute := range businessInfo.AttributeDefs {
			typeNames := []string{}
			json.Unmarshal([]byte(attribute.TypeName), &typeNames)
			for _, name := range typeNames {
				timestamp2 := int64(businessInfo.CreateTime) / 1000 // 转换为以秒为单位的时间戳
				t2 := time.Unix(timestamp2, 0)
				// 将time.Time对象格式化为指定格式的字符串
				formattedTime2 := t2.Format("2006-01-02 15:04:05")
				newbusinessAttribute := BusinessMetaAttributeInfo{
					Attributename: attribute.Name,
					Businessname:  businessInfo.Name,
					Businessguid:  businessInfo.GUID,
					Userid:        GetUserId(businessInfo.CreatedBy),
					Description:   businessInfo.Description,
					Createtime:    formattedTime2,
				}
				err = db.Create(&newbusinessAttribute).Error
				newbusinessAttributeType := BusinessMetaAttributeTypeInfo{
					Attributename: attribute.Name,
					Businessname:  businessInfo.Name,
					Userid:        GetUserId(businessInfo.CreatedBy),
					Createtime:    formattedTime2,
					Typename:      name,
				}
				err = db.Create(&newbusinessAttributeType).Error
			}

		}
	}
}

func AddBusinessMetaInfo(businessname string, username string, version int, desc string, guid string) {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	newbusiness := BusinessMetaInfo{
		Businessname: businessname,
		Userid:       GetUserId(username),
		Createtime:   timeString,
		Description:  desc,
		Guid:         guid,
		Version:      version,
	}
	err = db.Create(&newbusiness).Error
}

func AddBusinessAttributeInfo(businessname string, attribute string, username string, weight int, desc string, guid string) {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	newbusiness := BusinessMetaAttributeInfo{
		Businessname:  businessname,
		Weight:        weight,
		Userid:        GetUserId(username),
		Createtime:    timeString,
		Description:   desc,
		Attributename: attribute,
		Guid:          guid,
	}
	err = db.Create(&newbusiness).Error
}

func UpdateBusinessAttributeInfo(businessname string, attribute string, userid int, weight int, desc string) {
	businessAttr := BusinessMetaAttributeInfo{}
	_ = db.Model(&businessAttr).Where("businessname=?", businessname).Where("attributename=?", attribute).Update("weight", weight).Error
	_ = db.Model(&businessAttr).Where("businessname=?", businessname).Where("attributename=?", attribute).Update("description", desc)
	_ = db.Model(&businessAttr).Where("businessname=?", businessname).Where("attributename=?", attribute).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
	_ = db.Model(&businessAttr).Where("businessname=?", businessname).Where("attributename=?", attribute).Update("updateuserid", userid)
}

func UpdateBusinessAttributeTypes(businessname string, attribute string, types []string, userId int, guid string) {
	for _, typeInfo := range types {
		attributeType := BusinessMetaAttributeTypeInfo{}
		_ = db.Where("businessname=?", businessname).Where("attributename=?", attribute).Where("typename=?", typeInfo).Find(&attributeType)
		if attributeType.Id == 0 {
			newAttributeType := BusinessMetaAttributeTypeInfo{
				Businessname:  businessname,
				Attributename: attribute,
				Typename:      typeInfo,
				Userid:        userId,
				Guid:          guid,
				Createtime:    time.Now().Format("2006-01-02 15:04:05"),
			}
			db.Create(&newAttributeType)
		}
	}
}

func AddBusinessAttributeTypeInfo(typename string, businessname string, attribute string, username string, guid string) {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	newbusiness := BusinessMetaAttributeTypeInfo{
		Typename:      typename,
		Businessname:  businessname,
		Userid:        GetUserId(username),
		Createtime:    timeString,
		Attributename: attribute,
		Guid:          guid,
	}
	err = db.Create(&newbusiness).Error
}

func GetBusinessInfoByType(typename string, guid string) []guidAttributeLists {
	businessTypeInfos := []BusinessMetaAttributeTypeInfo{}
	businessInfo2 := []guidAttributeLists{}
	_ = db.Where("typename=?", typename).Find(&businessTypeInfos)
	for _, businessTypeInfo := range businessTypeInfos {
		businessInfo := BusinessMetaInfo{}
		_ = db.Where("businessname=?", businessTypeInfo.Businessname).Find(&businessInfo)

		attrbiteInfos := []BusinessMetaAttributeInfo{}
		_ = db.Where("attributename=?", businessTypeInfo.Attributename).Where("businessname=?", businessTypeInfo.Businessname).Find(&attrbiteInfos)

		for _, attributeInfo := range attrbiteInfos {
			isExitst := false
			entityInfo := BusinessMetaEntityInfo{}
			_ = db.Where("attributename=?", attributeInfo.Attributename).Where("businessname=?", attributeInfo.Businessname).Where("entityguid=?", guid).Find(&entityInfo)
			if entityInfo.Id != 0 {
				isExitst = true
			}
			businessInfo2 = append(businessInfo2, guidAttributeLists{
				Guid:          guid,
				BusinessName:  businessInfo.Businessname,
				BusinessDesc:  businessInfo.Description,
				AttributeName: attributeInfo.Attributename,
				AttributeDesc: attributeInfo.Description,
				IsExist:       isExitst,
			})
		}

	}
	return businessInfo2
}

func GetBusinessMeta(guid string) BusinessMetaInfo {
	businessInfo := BusinessMetaInfo{}
	_ = db.Where("guid=?", guid).Find(&businessInfo)
	username, avatar := GetUserInfo(businessInfo.Userid)
	businessInfo.Username = username
	businessInfo.Avatar = avatar
	return businessInfo
}

func GetBusinessEntityIds(businessName string) []BusinessMetaEntityInfo {
	businessEntitys := []BusinessMetaEntityInfo{}
	_ = db.Where("businessname=?", businessName).Find(&businessEntitys)
	return businessEntitys
}

func GetBusinessMetaAttribute(guid string) []BusinessMetaAttributeInfo {
	businessInfoAttribute := []BusinessMetaAttributeInfo{}
	businessInfoAttributes := []BusinessMetaAttributeInfo{}
	_ = db.Where("guid=?", guid).Find(&businessInfoAttribute)
	for _, info := range businessInfoAttribute {
		username, avatar := GetUserInfo(info.Userid)
		info.Username = username
		info.Avatar = avatar
		businessInfoAttributes = append(businessInfoAttributes, info)
	}
	return businessInfoAttributes
}

func GetBusinessInfo(name string) BusinessMetaInfoType {
	BusinessInfo := BusinessMetaInfo{}
	_ = db.Where("businessname=?", name).Find(&BusinessInfo)
	attributeInfo2 := []BusinessMetaAttributeInfo{}
	attributeInfo := []BusinessMetaAttributeInfo{}
	_ = db.Where("businessname=?", name).Find(&attributeInfo)
	for _, info := range attributeInfo {
		typeInfos := []BusinessMetaAttributeTypeInfo{}
		typeNames := []string{}
		_ = db.Where("attributename=?", info.Attributename).Where("businessname=?", info.Businessname).Find(&typeInfos)
		for _, typeInfo := range typeInfos {
			typeNames = append(typeNames, typeInfo.Typename)
		}
		info.Username, info.Avatar = GetUserInfo(info.Userid)
		info.UpdateUsername, info.UpdateAvatar = GetUserInfo(info.Updateuserid)
		info.Types = typeNames
		attributeInfo2 = append(attributeInfo2, info)
	}
	username, avatar := GetUserInfo(BusinessInfo.Userid)
	updateusername, updateavatar := GetUserInfo(BusinessInfo.Updateuserid)
	businessInfos := BusinessMetaInfoType{
		Id:             BusinessInfo.Id,
		Businessname:   BusinessInfo.Businessname,
		Userid:         BusinessInfo.Userid,
		Username:       username,
		Avatar:         avatar,
		Createtime:     BusinessInfo.Createtime,
		Description:    BusinessInfo.Description,
		Guid:           BusinessInfo.Guid,
		Updatetime:     BusinessInfo.Updatetime,
		Updateuserid:   BusinessInfo.Updateuserid,
		Version:        BusinessInfo.Version,
		AttributeInfos: attributeInfo2,
		UpdateAvatar:   updateavatar,
		UpdateUsername: updateusername,
	}
	return businessInfos
}

func GetAttributeInfoWithTypes(businessname string, attributename string) BusinessMetaAttributeInfo {
	attributeInfo := BusinessMetaAttributeInfo{}
	_ = db.Where("businessname=?", businessname).Where("attributename=?", attributename).Find(&attributeInfo)
	typeInfos := []BusinessMetaAttributeTypeInfo{}
	typeNames := []string{}
	_ = db.Where("attributename=?", attributename).Where("businessname=?", businessname).Find(&typeInfos)
	for _, typeInfo := range typeInfos {
		typeNames = append(typeNames, typeInfo.Typename)
	}
	attributeInfo.Types = typeNames
	return attributeInfo
}

func GetAddBusinessAttributesAtlasReq(atts BusinessAttributeAdd, username string, businessname string) BusinessAddAtlas {
	attributes := []AttributeDef{}
	oriBusiness := GetBusinessInfo(businessname)
	for _, att := range oriBusiness.AttributeInfos {
		typestring, _ := json.Marshal(att.Types)
		attributes = append(attributes, AttributeDef{
			Name:                  att.Attributename,
			TypeName:              "string",
			IsOptional:            true,
			Cardinality:           "SINGLE",
			ValuesMinCount:        0,
			ValuesMaxCount:        1,
			IsUnique:              false,
			IsIndexable:           true,
			IncludeInNotification: false,
			SearchWeight:          att.Weight,
			Options: struct {
				ApplicableEntityTypes string `json:"applicableEntityTypes"`
				MaxStrLength          string `json:"maxStrLength"`
			}{string(typestring), "50"},
		})
	}
	for _, att := range atts {
		typestring, _ := json.Marshal(att.Types)
		attributes = append(attributes, AttributeDef{
			Name:                  att.Name,
			TypeName:              "string",
			IsOptional:            true,
			Cardinality:           "SINGLE",
			ValuesMinCount:        0,
			ValuesMaxCount:        1,
			IsUnique:              false,
			IsIndexable:           true,
			IncludeInNotification: false,
			SearchWeight:          att.Weight,
			Options: struct {
				ApplicableEntityTypes string `json:"applicableEntityTypes"`
				MaxStrLength          string `json:"maxStrLength"`
			}{string(typestring), "50"},
		})
	}
	fmt.Println(oriBusiness.Createtime)
	fmt.Println(utils.TimeStringToUnix(oriBusiness.Createtime))
	BusinessInfo := BusinessDef{
		Category:      "BUSINESS_METADATA",
		GUID:          oriBusiness.Guid,
		CreatedBy:     oriBusiness.Username,
		UpdatedBy:     username,
		CreateTime:    utils.TimeStringToUnix(oriBusiness.Createtime),
		UpdateTime:    time.Now().Unix(),
		Version:       oriBusiness.Version,
		Name:          oriBusiness.Businessname,
		Description:   oriBusiness.Description,
		TypeVersion:   "1.1",
		AttributeDefs: attributes,
	}
	businessAddInfos := BusinessAddAtlas{
		EnumDefs:             nil,
		StructDefs:           nil,
		ClassificationDefs:   nil,
		EntityDefs:           nil,
		RelationshipDefs:     nil,
		BusinessMetadataDefs: []BusinessDef{BusinessInfo},
	}
	return businessAddInfos
}

func GetEntityBusinessInfos(guid string) []guidBusinessInfo {
	businessResInfos := []guidBusinessInfo{}
	businessEntityInfo := []BusinessMetaEntityInfo{}
	_ = db.Where("entityguid=?", guid).Find(&businessEntityInfo)
	businessNameMap := make(map[string]bool)
	for _, info := range businessEntityInfo {
		businessNameMap[info.Businessname] = true
	}
	for businessname, _ := range businessNameMap {
		businessEntityInfo2 := []BusinessMetaEntityInfo{}
		_ = db.Where("entityguid=?", guid).Where("businessname=?", businessname).Find(&businessEntityInfo2)
		businessInfo := BusinessMetaInfo{}
		_ = db.Where("businessname=?", businessname).Find(&businessInfo)
		attributeResInfos := []guidAttributeInfo{}
		for _, entityinfo := range businessEntityInfo2 {
			attributeInfo := BusinessMetaAttributeInfo{}
			_ = db.Where("businessname=?", businessname).Where("attributename=?", entityinfo.Attributename).Find(&attributeInfo)
			attributeResInfos = append(attributeResInfos, guidAttributeInfo{
				AttributeName:  attributeInfo.Attributename,
				AttributeValue: entityinfo.Attributevalue,
				AttributeDesc:  attributeInfo.Description,
				UserId:         attributeInfo.Userid,
				CreateTime:     attributeInfo.Createtime,
			})
		}
		businessResInfos = append(businessResInfos, guidBusinessInfo{
			Guid:           guid,
			UserId:         businessInfo.Userid,
			BusinessName:   businessname,
			BusinessDesc:   businessInfo.Description,
			CreateTime:     businessInfo.Createtime,
			AttributeInfos: attributeResInfos,
		})
	}
	return businessResInfos
}

func DeleteEntityBusinessInfo(guid, businessname string) {
	businessEntityInfo := BusinessMetaEntityInfo{}
	db.Where("entityguid= ?", guid).Where("businessname=?", businessname).Delete(&businessEntityInfo)
}

func DeleteEntityAttributeInfo(guid, businessname, attributename string) {
	businessEntityInfo := BusinessMetaEntityInfo{}
	db.Where("entityguid= ?", guid).Where("businessname=?", businessname).Where("attributename=?", attributename).Delete(&businessEntityInfo)
}

func UpdateEntityAttributeInfo(userid int, guid, businessname, attributename, attributevalue string) {
	businessEntityInfo := BusinessMetaEntityInfo{}
	_ = db.Model(&businessEntityInfo).Where("entityguid=?", guid).Where("businessname=?", businessname).Where("attributename=?", attributename).Update("attributevalue", attributevalue).Update("userid", userid)
}

func AddEntityBusinessInfo(userid int, guid, businessname, attributename, attributevalue string) {
	businessEntityInfo := BusinessMetaEntityInfo{
		Entityguid:     guid,
		Userid:         userid,
		Attributename:  attributename,
		Attributevalue: attributevalue,
		Createtime:     time.Now().Format("2006-01-02 15:04:05"),
		Businessname:   businessname,
	}
	_ = db.Create(&businessEntityInfo)
}

func UpdateBusinessInfo(name string, desc string, userid int) {
	businessInfo := BusinessMetaInfo{}
	_ = db.Model(&businessInfo).Where("businessname=?", name).Update("description", desc)
	_ = db.Model(&businessInfo).Where("businessname=?", name).Update("updateuserid", userid)
	_ = db.Model(&businessInfo).Where("businessname=?", name).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))

}

func UpdateBusinessVersion(businessname string, userid int, version int) int {
	newBusiness := BusinessMetaInfo{}
	_ = db.Model(&newBusiness).Where("businessname=?", businessname).Update("version", version).Error
	_ = db.Model(&newBusiness).Where("businessname=?", businessname).Update("updateuserid", userid)
	_ = db.Model(&newBusiness).Where("businessname=?", businessname).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
	//err = db.Debug().Find(&newBusiness).Update("version", version).Error
	//if err != nil {
	//	return utils.ERROR_CREAT_WRONG
	//}
	//_ = db.Debug().Find(&newBusiness).Update("updateuserid", userid)
	//_ = db.Debug().Find(&newBusiness).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))

	return utils.SUCCESS
}
