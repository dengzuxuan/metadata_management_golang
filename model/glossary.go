package model

import (
	"others-part/utils"
	"time"
)

type GlossaryReqAtlas struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
}
type GlossaryRespAtlas struct {
	GUID             string `json:"guid"`
	QualifiedName    string `json:"qualifiedName"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
}
type TermReqAtlas struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Anchor           struct {
		GlossaryGUID string `json:"glossaryGuid"`
		DisplayText  string `json:"displayText"`
	} `json:"anchor"`
}
type TermRespAtlas struct {
	GUID             string `json:"guid"`
	QualifiedName    string `json:"qualifiedName"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Anchor           struct {
		GlossaryGUID string `json:"glossaryGuid"`
		RelationGUID string `json:"relationGuid"`
	} `json:"anchor"`
}
type AddTermReq struct {
	Termname  string `json:"termname"`
	Shortdesc string `json:"shortdesc"`
	Longdesc  string `json:"longdesc"`
}

type AddTermInfo struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Anchor           struct {
		GlossaryGuid string `json:"glossaryGuid"`
		DisplayText  string `json:"displayText"`
	} `json:"anchor"`
}

type GlossaryInfo struct {
	Id               int    `gorm:"id" json:"id"`
	Glossaryname     string `gorm:"glossaryname" json:"glossaryname"`
	Shortdescription string `gorm:"shortdescription" json:"shortdescription"`
	Longdescription  string `gorm:"longdescription" json:"longdescription"`
	Userid           int    `gorm:"userid" json:"userid"`
	Username         string `gorm:"-" json:"username"`
	Avatar           string `gorm:"-" json:"avatar"`
	Termnumber       int    `gorm:"termnumber" json:"termnumber"`
	Createtime       string `gorm:"createtime" json:"createtime"`
	Guid             string `gorm:"guid" json:"guid"`

	Updatetime   string `gorm:"updatetime" json:"updatetime"`
	Updateuserid int    `gorm:"updateuserid" json:"updateuserid"`

	UpdateUsername string `gorm:"-" json:"update_username"`
	UpdateAvatar   string `gorm:"-" json:"update_avatar"`
}

type GlossaryTermsInfo struct {
	Id               int    `gorm:"id" json:"id"`
	Name             string `gorm:"-" json:"name"`
	Termname         string `gorm:"termname" json:"termname"`
	Userid           int    `gorm:"userid" json:"userid"`
	Username         string `gorm:"-" json:"username"`
	Avatar           string `gorm:"-" json:"avatar"`
	Createtime       string `gorm:"createtime" json:"createtime"`
	Shortdescription string `gorm:"shortdescription" json:"shortdescription"`
	Longdescription  string `gorm:"longdescription" json:"longdescription"`
	Guid             string `gorm:"guid" json:"guid"`
	Glossaryguid     string `gorm:"glossaryguid" json:"glossaryguid"`
	Glossaryname     string `gorm:"glossaryname" json:"glossaryname"`

	Updatetime   string `gorm:"updatetime" json:"updatetime"`
	Updateuserid int    `gorm:"updateuserid" json:"updateuserid"`

	UpdateUsername string `gorm:"-" json:"update_username"`
	UpdateAvatar   string `gorm:"-" json:"update_avatar"`
}
type GlossaryTermClassificationAttributeInfo struct {
	Id                 int    `gorm:"id" json:"id"`
	Termid             int    `gorm:"termid" json:"termid"`
	Guid               string `gorm:"guid" json:"guid"`
	TermName           string `gorm:"termname" json:"termName"`
	Classificationname string `gorm:"classificationname" json:"classificationname"`
	Attributename      string `gorm:"attributename" json:"attributename"`
	Attributevalue     string `gorm:"attributevalue" json:"attributevalue"`
	Userid             int    `gorm:"userid" json:"userid"`
	Createtime         string `gorm:"createtime" json:"createtime"`
}

func (this *GlossaryInfo) TableName() string {
	return "GlossaryInfo"
}
func (this *GlossaryTermsInfo) TableName() string {
	return "GlossaryTermsInfo"
}
func (this *GlossaryTermClassificationAttributeInfo) TableName() string {
	return "GlossaryTermClassificationAttributeInfo"
}
func AddGlossary(Glossaryname string, Shortdescription string, Longdescription string, userid int, username string, avatar string, guid string, Termnumber int) int {
	newGlossary := GlossaryInfo{
		Glossaryname:     Glossaryname,
		Shortdescription: Shortdescription,
		Longdescription:  Longdescription,
		Userid:           userid,
		Username:         username,
		Avatar:           avatar,
		Termnumber:       Termnumber,
		Createtime:       time.Now().Format("2006-01-02 15:04:05"),
		Guid:             guid,
	}
	err = db.Create(&newGlossary).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}
func GetGlossaryGuid(glossaryName string) string {
	glossary := GlossaryInfo{}
	_ = db.Where("Glossaryname=?", glossaryName).Find(&glossary)
	return glossary.Guid
}
func GetTermTotalName(guid string) string {
	term := GlossaryTermsInfo{}
	_ = db.Where("guid=?", guid).Find(&term)
	totalName := term.Termname + "@" + term.Glossaryname
	return totalName
}
func GetTermTotalName2(termname string) string {
	term := GlossaryTermsInfo{}
	_ = db.Where("termname=?", termname).Find(&term)
	totalName := term.Termname + "@" + term.Glossaryname
	return totalName
}
func GetGlossary(guid string) GlossaryInfo {
	glossary := GlossaryInfo{}
	_ = db.Where("guid=?", guid).Find(&glossary)
	username, avatar := GetUserInfo(glossary.Userid)
	upusername, upavatar := GetUserInfo(glossary.Updateuserid)
	glossary.Username = username
	glossary.Avatar = avatar
	glossary.UpdateUsername = upusername
	glossary.UpdateAvatar = upavatar
	return glossary
}
func UpdateGlossary(name string, longdesc string, shortdesc string, userid int) {
	glossaryInfo := GlossaryInfo{}
	_ = db.Model(&glossaryInfo).Where("glossaryname=?", name).Update("shortdescription", shortdesc)
	_ = db.Model(&glossaryInfo).Where("glossaryname=?", name).Update("longdescription", longdesc)
	_ = db.Model(&glossaryInfo).Where("glossaryname=?", name).Update("updateuserid", userid)
	_ = db.Model(&glossaryInfo).Where("glossaryname=?", name).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
}
func UpdateTerm(glossaryname string, termname string, longdesc string, shortdesc string, userid int) {
	termInfo := GlossaryTermsInfo{}
	_ = db.Model(&termInfo).Where("glossaryname=?", glossaryname).Where("termname=?", termname).Update("shortdescription", shortdesc)
	_ = db.Model(&termInfo).Where("glossaryname=?", glossaryname).Where("termname=?", termname).Update("longdescription", longdesc)
	_ = db.Model(&termInfo).Where("glossaryname=?", glossaryname).Where("termname=?", termname).Update("updateuserid", userid)
	_ = db.Model(&termInfo).Where("glossaryname=?", glossaryname).Where("termname=?", termname).Update("updatetime", time.Now().Format("2006-01-02 15:04:05"))
}
func GetGlossaryInfo(glossaryName string) GlossaryInfo {
	glossary := GlossaryInfo{}
	_ = db.Where("glossaryname=?", glossaryName).Find(&glossary)
	return glossary
}
func GetTermInfo(termName string, glossaryName string) GlossaryTermsInfo {
	termInfo := GlossaryTermsInfo{}
	_ = db.Where("Termname=?", termName).Where("Glossaryname=?", glossaryName).Find(&termInfo)
	username, avatar := GetUserInfo(termInfo.Userid)
	upusername, upavatar := GetUserInfo(termInfo.Updateuserid)
	termInfo.Username = username
	termInfo.Avatar = avatar
	termInfo.UpdateUsername = upusername
	termInfo.UpdateAvatar = upavatar
	return termInfo
}
func GetTerms(guid string) []GlossaryTermsInfo {
	termInfos := []GlossaryTermsInfo{}
	terms := []GlossaryTermsInfo{}
	_ = db.Where("glossaryguid=?", guid).Find(&terms)
	for _, term := range terms {
		username, avatar := GetUserInfo(term.Userid)
		term.Username = username
		term.Avatar = avatar
		termInfos = append(termInfos, term)
	}
	return termInfos
}
func GetTermClassifications(guid string) []GlossaryTermClassificationAttributeInfo {
	termClassifications := []GlossaryTermClassificationAttributeInfo{}
	_ = db.Where("Guid=?", guid).Find(&termClassifications)
	return termClassifications
}
func GetTermClassificationAttributes(guid string, classificationName string) []GlossaryTermClassificationAttributeInfo {
	termClassifications := []GlossaryTermClassificationAttributeInfo{}
	_ = db.Where("Classificationname=?", classificationName).Where("Guid=?", guid).Find(&termClassifications)
	return termClassifications
}
func AddTerm(Termname string, Glossaryname string, Shortdescription string, Longdescription string, userid int, guid string, Glossaryguid string) int {
	newTerm := GlossaryTermsInfo{
		Termname:         Termname,
		Userid:           userid,
		Createtime:       time.Now().Format("2006-01-02 15:04:05"),
		Shortdescription: Shortdescription,
		Longdescription:  Longdescription,
		Guid:             guid,
		Glossaryname:     Glossaryname,
		Glossaryguid:     Glossaryguid,
	}
	err = db.Create(&newTerm).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}
