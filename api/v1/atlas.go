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

func SearchPre(c *gin.Context) {

	username := c.GetHeader("username")
	password := c.GetHeader("password")
	query := c.Query("query")
	queryParams := make(map[string]string)
	queryParams["query"] = query
	queryParams["limit"] = "5"
	queryParams["offset"] = "0"
	s, _ := utils.Call("atlas/v2/search/quick", username, password, "GET", queryParams, nil)
	var atlasSearchPre model.AtlasSearchPreType
	_ = json.Unmarshal(s, &atlasSearchPre)
	atlasSearchPreMap := utils.JSONMethod(atlasSearchPre)
	atlasSearchPreMap["status"] = utils.SUCCESS
	atlasSearchPreMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, atlasSearchPreMap)
}

func SearchResult(c *gin.Context) {
	guid := c.Query("guid")
	queryParams := make(map[string]string)
	queryParams["guid"] = guid

	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)

	username := c.GetHeader("username")
	password := c.GetHeader("password")

	entity := make(map[string]interface{})
	entityJson, _ := utils.Call("atlas/v2/entity/guid/"+guid, username, password, "GET", nil, nil)
	_ = json.Unmarshal(entityJson, &entity)

	lineage := make(map[string]interface{})
	lineageJson, _ := utils.Call("atlas/v2/lineage/"+guid, username, password, "GET", nil, nil)
	infos := analysisLineage(lineageJson)
	_ = json.Unmarshal(lineageJson, &lineage)
	entity["lineageSite"] = infos
	entity["lineage"] = lineage
	classifications := make(map[string]interface{})
	classificationsJson, _ := utils.Call("atlas/v2/entity/guid/"+guid+"/classifications", username, password, "GET", nil, nil)
	_ = json.Unmarshal(classificationsJson, &classifications)
	entity["classifications"] = classifications

	entity["comments"] = model.GetCommentInfo(guid, "", useridInt)
	otherInfos := []map[string]interface{}{}
	otherInfoJson, _ := utils.Call("atlas/v2/entity/"+guid+"/audit", username, password, "GET", nil, nil)
	_ = json.Unmarshal(otherInfoJson, &otherInfos)
	for _, info := range otherInfos {
		details := info["details"].(string)
		detailsName := strings.Split(details, ":")[0]
		if !strings.Contains(details, "{") {
			detailsArray := strings.Split(details, ":")
			actionDetailsMap := make(map[string]interface{})
			if len(detailsArray) > 1 {
				actionDetailsMap[detailsArray[0]] = detailsArray[1]
			}
			info[detailsName] = actionDetailsMap
			continue
		}
		actionDetails := strings.Join(strings.Split(details, "{")[1:], "{")
		actionDetails = "{" + actionDetails
		actionDetailsMap := make(map[string]interface{})
		json.Unmarshal([]byte(actionDetails), &actionDetailsMap)
		if actionDetailsMap == nil {
			fmt.Println("actionDetailsMap is null", details)
		}
		info[detailsName] = actionDetailsMap
	}
	entity["otherInfos"] = otherInfos

	entity["status"] = utils.SUCCESS
	entity["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, entity)
}

func TypeEntity(c *gin.Context) {
	type TypeEntitys struct {
		Name        string      `json:"name"`
		Count       interface{} `json:"count"`
		Description string      `json:"description"`
		UserName    string      `json:"user_name"`
		CreateTime  string      `json:"create_time"`
		Avatar      string      `json:"avatar"`
	}
	typeEntitys := []TypeEntitys{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	entritySting, _ := utils.Call("atlas/admin/metrics", username, password, "GET", nil, nil)
	entityMap := make(map[string]interface{})
	json.Unmarshal(entritySting, &entityMap)
	entityTypes := entityMap["data"].(map[string]interface{})["entity"].(map[string]interface{})["entityActive"].(map[string]interface{})
	for typeName, cnt := range entityTypes {
		//model.AddEntityType(typeName)
		info := model.GetEntityType(typeName)
		typeEntitys = append(typeEntitys, TypeEntitys{
			Name:        typeName,
			Count:       cnt,
			Description: info.Description,
			UserName:    info.Username,
			CreateTime:  info.Createtime,
			Avatar:      info.Avatar,
		})
	}
	entityMaps := map[string]interface{}{
		"entity":  typeEntitys,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	}
	c.JSON(http.StatusOK, entityMaps)
}
func TypeClassification(c *gin.Context) {
	type Attribute struct {
		Name string `json:"name"`
	}
	type ClassificationEntitys struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		UserName    string      `json:"user_name"`
		CreateTime  string      `json:"create_time"`
		Avatar      string      `json:"avatar"`
		Guid        string      `json:"guid"`
		Attributes  []Attribute `json:"attributes"`
	}
	classificationEntitys := []ClassificationEntitys{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	query := map[string]string{
		"type": "classification",
	}
	classificatonSting, _ := utils.Call("atlas/v2/types/typedefs", username, password, "GET", query, nil)
	classificatonType := model.AtlasClassification{}
	_ = json.Unmarshal(classificatonSting, &classificatonType)
	for _, info := range classificatonType.ClassificationDefs {
		classificationName := info.Name
		Attributes := []Attribute{}
		info2 := model.GetClassificatioInfo(classificationName)
		attributeInfos := model.GetClassificationAttribute(info2.Guid)
		for _, attributeInfo := range attributeInfos {
			Attributes = append(Attributes, Attribute{Name: attributeInfo.Attributename})
		}
		classificationEntitys = append(classificationEntitys, ClassificationEntitys{
			Name:        classificationName,
			Description: info2.Description,
			UserName:    info2.Username,
			CreateTime:  info2.Createtime,
			Avatar:      info2.Avatar,
			Guid:        info2.Guid,
			Attributes:  Attributes,
		})
	}
	classificatonMaps := map[string]interface{}{
		"classificaton": classificationEntitys,
		"status":        utils.SUCCESS,
		"message":       utils.GetErrMsg(utils.SUCCESS),
	}
	c.JSON(http.StatusOK, classificatonMaps)
}
func TypeGlossary2(c *gin.Context) {
	type Term struct {
		Name string `json:"name"`
	}
	type GlossaryEntitys struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		DescriptionLong string `json:"description_long"`
		UserName        string `json:"user_name"`
		CreateTime      string `json:"create_time"`
		Avatar          string `json:"avatar"`
		Guid            string `json:"guid"`
		Terms           []Term `json:"terms"`
	}
	glossaryEntitys := []GlossaryEntitys{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	glossarySting, _ := utils.Call("atlas/v2/glossary", username, password, "GET", nil, nil)

	glossaryType := model.AtlasGlossary{}
	_ = json.Unmarshal(glossarySting, &glossaryType)
	for _, info := range glossaryType {
		info2 := model.GetGlossary(info.GUID)
		terms := []Term{}
		termsInfo := model.GetTerms(info.GUID)
		for _, glossaryTermsInfo := range termsInfo {
			terms = append(terms, Term{Name: glossaryTermsInfo.Termname})
		}
		glossaryEntitys = append(glossaryEntitys, GlossaryEntitys{
			Name:            info.QualifiedName,
			Description:     info.ShortDescription,
			DescriptionLong: info.LongDescription,
			UserName:        info2.Username,
			CreateTime:      info2.Createtime,
			Avatar:          info2.Avatar,
			Guid:            info2.Guid,
			Terms:           terms,
		})
	}
	glossaryMaps := map[string]interface{}{
		"glossary": glossaryEntitys,
		"status":   utils.SUCCESS,
		"message":  utils.GetErrMsg(utils.SUCCESS),
	}
	c.JSON(http.StatusOK, glossaryMaps)
}
func TypeGlossary(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	glossarySting, _ := utils.Call("atlas/v2/glossary", username, password, "GET", nil, nil)
	glossaryType := make(map[string]interface{})
	glossaryTypes := []map[string]interface{}{}
	_ = json.Unmarshal(glossarySting, &glossaryTypes)
	glossaryType["glossary"] = glossaryTypes
	glossaryType["status"] = utils.SUCCESS
	glossaryType["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, glossaryType)
}
func TypeBussinessMetadataGlossary(c *gin.Context) {
	type Attribute struct {
		Name string `json:"name"`
	}
	type BussinessMetaEntitys struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		UserName    string      `json:"user_name"`
		CreateTime  string      `json:"create_time"`
		Avatar      string      `json:"avatar"`
		Guid        string      `json:"guid"`
		Attributes  []Attribute `json:"attributes"`
	}
	bussinessMetaEntitys := []BussinessMetaEntitys{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	query := map[string]string{
		"type": "business_metadata",
	}
	bussinessMetadataSting, _ := utils.Call("atlas/v2/types/typedefs", username, password, "GET", query, nil)
	//AddBusiness(bussinessMetadataSting)
	bussinessMetadataType := model.AtlasBusinessMeta{}
	_ = json.Unmarshal(bussinessMetadataSting, &bussinessMetadataType)
	for _, info := range bussinessMetadataType.BusinessMetadataDefs {
		bussinessMetaName := info.Name
		Attributes := []Attribute{}
		info2 := model.GetBusinessMeta(info.GUID)
		attributeInfos := model.GetBusinessMetaAttribute(info2.Guid)
		for _, attributeInfo := range attributeInfos {
			Attributes = append(Attributes, Attribute{Name: attributeInfo.Attributename})
		}
		bussinessMetaEntitys = append(bussinessMetaEntitys, BussinessMetaEntitys{
			Name:        bussinessMetaName,
			Description: info2.Description,
			UserName:    info2.Username,
			CreateTime:  info2.Createtime,
			Avatar:      info2.Avatar,
			Guid:        info2.Guid,
			Attributes:  Attributes,
		})
	}
	bussinessMaps := map[string]interface{}{
		"bussinessMeta": bussinessMetaEntitys,
		"status":        utils.SUCCESS,
		"message":       utils.GetErrMsg(utils.SUCCESS),
	}
	c.JSON(http.StatusOK, bussinessMaps)
}

func FindTypeDetails(c *gin.Context) {
	find := c.Query("find")
	typeName := strings.ReplaceAll(c.Query("typename"), " ", "")
	pageLimit := c.Query("pageLimit")
	pageOffset := c.Query("pageOffset")
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	body := map[string]interface{}{
		"excludeDeletedEntities":          true,
		"includeSubClassifications":       true,
		"includeSubTypes":                 true,
		"includeClassificationAttributes": true,
		"entityFilters":                   nil,
		"tagFilters":                      nil,
		"attributes":                      nil,
		"offset":                          pageOffset,
		"limit":                           pageLimit,
	}

	typeDetailsMap := make(map[string]interface{})
	switch find {
	case "entity":
		body["typeName"] = typeName
		body["classification"] = nil
		body["termName"] = nil
		info := model.GetEntityType(typeName)
		typeDetailsMap["info"] = info
		typeDetails, _ := utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetails, &typeDetailsMap)
	case "classification":
		body["classification"] = typeName
		body["typeName"] = nil
		body["termName"] = nil
		info := model.GetClassificatioInfo(typeName)
		typeDetailsMap["info"] = info
		typeDetails, _ := utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetails, &typeDetailsMap)
	case "term":
		body["termName"] = typeName
		body["typeName"] = nil
		body["classification"] = nil
		if strings.Contains(typeName, "@") {
			termName := strings.Split(typeName, "@")[0]
			glossaryName := strings.Split(typeName, "@")[1]
			info := model.GetTermInfo(termName, glossaryName)
			typeDetailsMap["info"] = info
		}
		typeDetails, _ := utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetails, &typeDetailsMap)
	case "glossary":
		typeDetailsMap = GetGlossaryInfo2(typeName)
	}
	typeDetailsMap["status"] = utils.SUCCESS
	typeDetailsMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, typeDetailsMap)
}
