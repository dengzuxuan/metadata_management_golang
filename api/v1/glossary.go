package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
	"strings"
)

func AddGlossaryInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	avatar := model.GetUserAvatar(useridInt)
	var GlossaryReq model.GlossaryReqAtlas
	_ = c.ShouldBindJSON(&GlossaryReq)
	//atlas
	addAtlasGlossary, _ := utils.Call("atlas/v2/glossary", username, password, "POST", nil, GlossaryReq)
	addAtlasGlossaryResp := make(map[string]interface{})
	_ = json.Unmarshal(addAtlasGlossary, &addAtlasGlossaryResp)

	//mysql
	var glossaryRespAtlas model.GlossaryRespAtlas
	_ = json.Unmarshal(addAtlasGlossary, &glossaryRespAtlas)
	model.AddGlossary(glossaryRespAtlas.Name, glossaryRespAtlas.ShortDescription, glossaryRespAtlas.LongDescription, useridInt, username, avatar, glossaryRespAtlas.GUID, 0)

	model.AddTypeRecord(useridInt, "Glossary Create", string(addAtlasGlossary), GlossaryReq.Name)
	addAtlasglossaryResp := make(map[string]interface{})
	addAtlasglossaryResp["status"] = utils.SUCCESS
	addAtlasglossaryResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addAtlasglossaryResp)
}

func AddTermInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	var termReq model.TermReqAtlas
	_ = c.ShouldBindJSON(&termReq)
	guid := model.GetGlossaryGuid(termReq.Anchor.DisplayText)
	termReq.Anchor.GlossaryGUID = guid
	//atlas
	addAtlasTerm, _ := utils.Call("atlas/v2/glossary/term", username, password, "POST", nil, termReq)
	addAtlasTermResp := make(map[string]interface{})
	_ = json.Unmarshal(addAtlasTerm, &addAtlasTermResp)

	//mysql
	var termRespAtlas model.TermRespAtlas
	var glossaryName string
	_ = json.Unmarshal(addAtlasTerm, &termRespAtlas)
	if strings.Contains(termRespAtlas.QualifiedName, "@") {
		glossaryName = strings.Split(termRespAtlas.QualifiedName, "@")[1]
	}
	model.AddTerm(termRespAtlas.Name, glossaryName, termRespAtlas.ShortDescription, termRespAtlas.LongDescription, useridInt, termRespAtlas.GUID, termRespAtlas.Anchor.RelationGUID, termReq.Anchor.GlossaryGUID)

	model.AddTypeRecord(useridInt, "Term Create", string(addAtlasTerm), termReq.Name)
	addTermResp := make(map[string]interface{})
	addTermResp["status"] = utils.SUCCESS
	addTermResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addTermResp)
}

func GetGlossaryInfo(c *gin.Context) {
	guid := c.Query("guid")
	glossaryInfo := model.GetGlossary(guid)
	termsInfo := model.GetTerms(guid)
	termsInfosMap := []map[string]interface{}{}
	for _, term := range termsInfo {
		termInfoMap := utils.JSONMethod(term)
		termClassifications := model.GetTermClassifications(term.Guid)
		classificationMap := make(map[string]interface{})
		for _, termClassification := range termClassifications {
			termClassificationAttributeMap := make(map[string]string)
			termClassificationAttributes := model.GetTermClassificationAttributes(term.Guid, termClassification.Classificationname)
			for _, attribute := range termClassificationAttributes {
				if attribute.Attributevalue != "" {
					termClassificationAttributeMap[attribute.Attributename] = attribute.Attributevalue
				}
			}
			classificationMap[termClassification.Classificationname] = termClassificationAttributeMap
		}
		termInfoMap["attributes"] = classificationMap
		termsInfosMap = append(termsInfosMap, termInfoMap)
	}
	glossaryMap := make(map[string]interface{})
	glossaryMap["glossary"] = glossaryInfo
	glossaryMap["terms"] = termsInfosMap
	glossaryMap["status"] = utils.SUCCESS
	glossaryMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, glossaryMap)
}

func GetGlossaryInfo2(glossaryName string) map[string]interface{} {
	guid := model.GetGlossaryGuid(glossaryName)
	glossaryInfo := model.GetGlossary(guid)
	termsInfo := model.GetTerms(guid)
	termEntitys := []map[string]interface{}{}
	for _, term := range termsInfo {
		termEntity := make(map[string]interface{})
		termEntity["createtime"] = term.Createtime
		termEntity["shortdescription"] = term.Shortdescription
		termEntity["relationshipGuid"] = term.Relationshipguid
		termEntity["guid"] = term.Guid
		termEntity["attributes"] = map[string]string{"qualifiedName": term.Termname, "owner": term.Username, "avatar": term.Avatar}
		termClassifications := model.GetTermClassifications(term.Guid)
		classificationMaps := []map[string]interface{}{}
		for _, termClassification := range termClassifications {
			classificationMap := make(map[string]interface{})
			classificationMap["typeName"] = termClassification.Classificationname
			termClassificationAttributeMap := make(map[string]string)
			termClassificationAttributes := model.GetTermClassificationAttributes(term.Guid, termClassification.Classificationname)
			for _, attribute := range termClassificationAttributes {
				if attribute.Attributevalue != "" {
					termClassificationAttributeMap[attribute.Attributename] = attribute.Attributevalue
				}
			}
			classificationMap["attributes"] = termClassificationAttributeMap
			classificationMaps = append(classificationMaps, classificationMap)
		}
		termEntity["classifications"] = classificationMaps
		termEntitys = append(termEntitys, termEntity)
	}
	glossaryMap := make(map[string]interface{})
	glossaryMap["info"] = glossaryInfo
	glossaryMap["entities"] = termEntitys
	return glossaryMap
}
func GetTermTotalName(c *gin.Context) {
	guid := c.Query("guid")
	termMap := make(map[string]interface{})
	termName := model.GetTermTotalName(guid)
	termMap["termName"] = termName
	termMap["status"] = utils.SUCCESS
	termMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, termMap)
}
func GetTermTotalName2(c *gin.Context) {
	termname := c.Query("termname")
	termMap := make(map[string]interface{})
	termName := model.GetTermTotalName2(termname)
	termMap["termName"] = termName
	termMap["status"] = utils.SUCCESS
	termMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, termMap)
}
func AddEntityTermInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	//userid := c.GetHeader("user_id")
	guid := c.Query("guid")
	Termname := c.Query("termname")
	glossaryname := c.Query("glossaryname")
	termInfo := model.GetTermInfo(Termname, glossaryname)
	termGuid := termInfo.Guid
	//useridInt, _ := strconv.Atoi(userid)
	/*
		shortdesc := c.Query("shortdesc")
			longdesc := c.Query("longdesc")
			termName := c.Query("termname")
			glossaryNameText := c.Query("glossaryname")
			glossaryGuid := model.GetGlossaryGuid(glossaryNameText)
			termReq := model.TermReqAtlas{
				Name:             termName,
				ShortDescription: shortdesc,
				LongDescription:  longdesc,
				Anchor: struct {
					GlossaryGUID string `json:"glossaryGuid"`
					DisplayText  string `json:"displayText"`
				}(struct {
					GlossaryGUID string
					DisplayText  string
				}{GlossaryGUID: glossaryGuid, DisplayText: glossaryNameText}),
			}*/
	var termReq []model.TermAtlasAdd
	termReq = append(termReq, model.TermAtlasAdd{guid})
	// http://hadoop102:21000/api/atlas/v2/glossary/terms/f13ee363-926a-4623-bf64-e435dc405296/assignedEntities
	//http://hadoop102:21000/api/atlas/v2/glossary/terms/b8fef1c0-8592-4be1-9c21-9270bb1fbc20/assignedEntities
	addAtlasTerm, _ := utils.Call("atlas/v2/glossary/terms/"+termGuid+"/assignedEntities", username, password, "POST", nil, termReq)
	fmt.Println(string(addAtlasTerm))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
		"infos":   addAtlasTerm,
	})
}

func GetGlossaryName(c *gin.Context) {
	guid := c.Query("guid")
	name := model.GetGlossaryName(guid)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
		"name":    name,
	})
}

func GetAllGlossaryInfos(c *gin.Context) {
	infos := model.GetAllGlossaryInfos()
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
		"infos":   infos,
	})
}
func GetAllTermInfos(c *gin.Context) {
	name := c.Query("glossaryname")
	infos := model.GetAllTerms(name)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
		"terms":   infos,
	})
}
