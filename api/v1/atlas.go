package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"regexp"
	"strconv"
	"strings"
)

func SearchPre(c *gin.Context) {

	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)

	entity := make(map[string]interface{})
	entityJson, _ := utils.Call("atlas/v2/entity/guid/"+guid, username, password, "GET", nil, nil)
	_ = json.Unmarshal(entityJson, &entity)

	lineage := make(map[string]interface{})
	lineageJson, _ := utils.Call("atlas/v2/lineage/"+guid, username, password, "GET", nil, nil)
	infos := analysisLineage(lineageJson)
	_ = json.Unmarshal(lineageJson, &lineage)
	entity["lineageSite"] = infos
	entity["lineage"] = lineage
	entity["businessmetadata"] = model.GetEntityBusinessInfos(guid)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	password1, _ := c.Get("password")
	password := password1.(string)
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
	typeDetailJson := []byte{}
	switch find {
	case "entity":
		body["typeName"] = typeName
		body["classification"] = nil
		body["termName"] = nil
		info := model.GetEntityType(typeName)
		typeDetailsMap["info"] = info
		typeDetailJson, _ = utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetailJson, &typeDetailsMap)
		if typeDetailsMap["entities"] != nil {
			for _, entityInfo := range typeDetailsMap["entities"].([]interface{}) {
				entityInfoMap := entityInfo.(map[string]interface{})
				guid := entityInfoMap["guid"]
				businessInfo := model.GetEntityBusinessInfos(guid.(string))
				entityInfoMap["businessInfo"] = businessInfo
			}
		}
	case "classification":
		body["classification"] = typeName
		body["typeName"] = nil
		body["termName"] = nil
		info := model.GetClassificatioInfo(typeName)
		typeDetailsMap["info"] = info
		typeDetailJson, _ = utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetailJson, &typeDetailsMap)
		if typeDetailsMap["entities"] != nil {
			for _, entityInfo := range typeDetailsMap["entities"].([]interface{}) {
				entityInfoMap := entityInfo.(map[string]interface{})
				guid := entityInfoMap["guid"]
				businessInfo := model.GetEntityBusinessInfos(guid.(string))
				entityInfoMap["businessInfo"] = businessInfo
			}
		}
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
		typeDetailJson, _ = utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
		_ = json.Unmarshal(typeDetailJson, &typeDetailsMap)
		if typeDetailsMap["entities"] != nil {
			for _, entityInfo := range typeDetailsMap["entities"].([]interface{}) {
				entityInfoMap := entityInfo.(map[string]interface{})
				guid := entityInfoMap["guid"]
				businessInfo := model.GetEntityBusinessInfos(guid.(string))
				entityInfoMap["businessInfo"] = businessInfo
			}
		}
	case "glossary":
		typeDetailsMap = GetGlossaryInfo2(typeName)
	case "business":
		pageLimitInt, _ := strconv.Atoi(pageLimit)
		pageOffsetInt, _ := strconv.Atoi(pageOffset)
		typeDetailsMap = GetBusinessNameInfos(username, password, typeName, pageLimitInt, pageOffsetInt)

		info := model.GetBusinessInfo(typeName)
		typeDetailsMap["info"] = info
		if typeDetailsMap["entities"] != nil {
			for _, entityInfoMap := range typeDetailsMap["entities"].([]map[string]interface{}) {
				//entityInfoMap := entityInfo.(map[string]interface{})
				guid := entityInfoMap["guid"]
				businessInfo := model.GetEntityBusinessInfos(guid.(string))
				entityInfoMap["businessInfo"] = businessInfo
			}
		}
	}
	typeDetailsMap["status"] = utils.SUCCESS
	typeDetailsMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, typeDetailsMap)
}

func UpdateTitleInfo(c *gin.Context) {

	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	updateInfo := c.Query("content")
	//username := c.GetHeader("username")
	switch typeInfo {
	case "entity":

	case "classification":
		classificationUpdate := model.ClassificationInfoType{}
		json.Unmarshal([]byte(updateInfo), &classificationUpdate)
		classificationOri := model.GetClassificatioInfo(classificationUpdate.Classificationname)
		if classificationOri.Description != classificationUpdate.Description {
			model.AddTypeRecord(useridInt, "Update Classification Description", updateInfo, classificationOri.Classificationname)
			model.UpdateClassification(classificationOri.Classificationname, classificationUpdate.Description, useridInt)
		}
	case "business":
		businessUpdate := model.BusinessMetaInfoType{}
		json.Unmarshal([]byte(updateInfo), &businessUpdate)
		businessOri := model.GetBusinessInfo(businessUpdate.Businessname)
		if businessUpdate.Description != businessOri.Description {
			model.AddTypeRecord(useridInt, "Update Business Metadata Description", updateInfo, businessOri.Businessname)
			model.UpdateBusinessInfo(businessOri.Businessname, businessUpdate.Description, useridInt)
		}
	case "glossary":
		glossaryUpdate := model.GlossaryInfo{}
		json.Unmarshal([]byte(updateInfo), &glossaryUpdate)
		glossaryOri := model.GetGlossaryInfo(glossaryUpdate.Glossaryname)
		if glossaryOri.Shortdescription != glossaryUpdate.Shortdescription || glossaryOri.Longdescription != glossaryUpdate.Longdescription {
			model.AddTypeRecord(useridInt, "Update Glossary Description", updateInfo, glossaryOri.Glossaryname)
			model.UpdateGlossary(glossaryOri.Glossaryname, glossaryUpdate.Longdescription, glossaryUpdate.Shortdescription, useridInt)
		}
	case "term":
		termUpdate := model.GlossaryTermsInfo{}
		json.Unmarshal([]byte(updateInfo), &termUpdate)
		termOri := model.GetTermInfo(termUpdate.Termname, termUpdate.Glossaryname)
		if termOri.Shortdescription != termUpdate.Shortdescription || termOri.Longdescription != termUpdate.Longdescription {
			model.AddTypeRecord(useridInt, "Update Term Description", updateInfo, termOri.Termname+"@"+termOri.Glossaryname)
			model.UpdateTerm(termOri.Glossaryname, termOri.Termname, termUpdate.Longdescription, termUpdate.Shortdescription, useridInt)
		}
	}
	resMap := map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	}
	c.JSON(http.StatusOK, resMap)
}

func AddInfos(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	typeName := c.Query("typename")
	content := c.Query("content")
	switch typeInfo {
	case "glossary":
		guid := model.GetGlossaryGuid(typeName)
		termInfoReq := model.AddTermReq{}
		json.Unmarshal([]byte(content), &termInfoReq)
		termInfoAtlasReq := model.AddTermInfo{
			Name:             termInfoReq.Termname,
			ShortDescription: termInfoReq.Shortdesc,
			LongDescription:  termInfoReq.Longdesc,
			Anchor: struct {
				GlossaryGuid string `json:"glossaryGuid"`
				DisplayText  string `json:"displayText"`
			}{guid, typeName},
		}
		//http://hadoop102:21000/api/atlas/v2/glossary/term
		addAtlasTerm, _ := utils.Call("atlas/v2/glossary/term", username, password, "POST", nil, termInfoAtlasReq)
		addAtlasTermResp := make(map[string]interface{})
		_ = json.Unmarshal(addAtlasTerm, &addAtlasTermResp)

		//mysql
		var termRespAtlas model.TermRespAtlas
		var glossaryName string
		_ = json.Unmarshal(addAtlasTerm, &termRespAtlas)
		if strings.Contains(termRespAtlas.QualifiedName, "@") {
			glossaryName = strings.Split(termRespAtlas.QualifiedName, "@")[1]
		}
		model.AddTerm(termRespAtlas.Name, glossaryName, termRespAtlas.ShortDescription, termRespAtlas.LongDescription, useridInt, termRespAtlas.Anchor.RelationGUID, termRespAtlas.GUID, guid)

		model.AddTypeRecord(useridInt, "Add Term", content, termRespAtlas.QualifiedName)
		addTermResp := make(map[string]interface{})
		addTermResp["status"] = utils.SUCCESS
		addTermResp["message"] = utils.GetErrMsg(utils.SUCCESS)
		c.JSON(http.StatusOK, addTermResp)

	case "classification":
		attributeInfoReq := model.AddClassificationReq{}
		_ = json.Unmarshal([]byte(content), &attributeInfoReq)
		addAttributeAtlasReq := model.GetAddAttributesAtlasReq(attributeInfoReq, username, typeName)
		//http://hadoop102:21000/api/atlas/v2/types/typedefs?type=classification
		addAtlasAttr, _ := utils.Call("atlas/v2/types/typedefs", username, password, "PUT", map[string]string{"type": "classification"}, addAttributeAtlasReq)
		addAtlasAttrResp := make(map[string]interface{})
		_ = json.Unmarshal(addAtlasAttr, &addAtlasAttrResp)
		//mysql
		//1.更新classification version
		//2.更新attribute info
		var classificationRespAtlas model.ClassificationRespAtlas
		_ = json.Unmarshal(addAtlasAttr, &classificationRespAtlas)
		if len(classificationRespAtlas.ClassificationDefs) != 0 {
			for _, classification := range classificationRespAtlas.ClassificationDefs {
				model.UpdateClassificationVersion(classification.Name, useridInt, classificationRespAtlas.ClassificationDefs[0].Version)
			}
		}
		for _, attribute := range attributeInfoReq {
			model.AddClassificationAttribute(typeName, attribute.Name, useridInt, attribute.Description, classificationRespAtlas.ClassificationDefs[0].GUID)
		}

		model.AddTypeRecord(useridInt, "Add Classification Attribute", content, classificationRespAtlas.ClassificationDefs[0].Name)

	case "business":
		attributeInfoReq := model.BusinessAttributeAdd{}
		_ = json.Unmarshal([]byte(content), &attributeInfoReq)
		attributeInfoAtlasReq := model.GetAddBusinessAttributesAtlasReq(attributeInfoReq, username, typeName)
		addAtlasAttr, _ := utils.Call("atlas/v2/types/typedefs", username, password, "PUT", map[string]string{"type": "business_metadata"}, attributeInfoAtlasReq)
		//addAtlasAttrResp := make(map[string]interface{})
		//mysql
		var businessRespAtlas model.BMRespAtlas
		_ = json.Unmarshal(addAtlasAttr, &businessRespAtlas)
		if len(businessRespAtlas.BusinessMetadataDefs) != 0 {
			for _, business := range businessRespAtlas.BusinessMetadataDefs {
				model.UpdateBusinessVersion(business.Name, useridInt, business.Version)
			}
		}
		for _, attributeInfo := range attributeInfoReq {
			model.AddBusinessAttributeInfo(typeName, attributeInfo.Name, username, attributeInfo.Weight, attributeInfo.Desc, businessRespAtlas.BusinessMetadataDefs[0].GUID)
			for _, typename := range attributeInfo.Types {
				model.AddBusinessAttributeTypeInfo(typename, typeName, attributeInfo.Name, username, businessRespAtlas.BusinessMetadataDefs[0].GUID)
			}
		}

		model.AddTypeRecord(useridInt, "Add Business Metadata Attribute", content, businessRespAtlas.BusinessMetadataDefs[0].Name)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func UpdateAttrInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	typeName := c.Query("typename")
	content := c.Query("content")
	switch typeInfo {
	case "classification":
		attribute := model.ClassificationAttributeUpdateInfo{}
		_ = json.Unmarshal([]byte(content), &attribute)
		ori := model.GetAttributeInfo(attribute.Classificationname, attribute.Attributename)
		if ori.Description != attribute.Description {
			model.UpdateAttribute(ori.Classificationname, attribute.Attributename, attribute.Description, useridInt)
			model.AddTypeRecord(useridInt, "Update Attribute Info Description", content, ori.Classificationname)
		}
	case "business":
		attribute := model.BusinessAttributeUpdate{}
		_ = json.Unmarshal([]byte(content), &attribute)
		oriAttribute := model.GetAttributeInfoWithTypes(typeName, attribute.Attributename)
		if len(attribute.Types) < len(oriAttribute.Types) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  utils.ERROR_WRONG_CHANGE,
				"message": utils.GetErrMsg(utils.ERROR_WRONG_CHANGE),
			})
			return
		}
		for _, oriType := range oriAttribute.Types {
			flag := false
			for _, newType := range attribute.Types {
				if newType != oriType {
					continue
				} else {
					flag = true
					break
				}
			}
			if flag == false {
				c.JSON(http.StatusOK, map[string]interface{}{
					"status":  utils.ERROR_WRONG_CHANGE,
					"message": utils.GetErrMsg(utils.ERROR_WRONG_CHANGE),
				})
				return
			}
		}

		if attribute.Weight == oriAttribute.Weight && len(attribute.Types) == len(oriAttribute.Types) {
			if attribute.Description == oriAttribute.Description {
				c.JSON(http.StatusOK, map[string]interface{}{
					"status":  utils.SUCCESS,
					"message": utils.GetErrMsg(utils.SUCCESS),
				})
				return
			} else {
				model.UpdateBusinessAttributeInfo(typeName, attribute.Attributename, useridInt, attribute.Weight, attribute.Description)
				model.AddTypeRecord(useridInt, "Update Business Metadata Attribute Description", content, oriAttribute.Businessname)
				c.JSON(http.StatusOK, map[string]interface{}{
					"status":  utils.SUCCESS,
					"message": utils.GetErrMsg(utils.SUCCESS),
				})
				return
			}
		}
		//atlas
		attributeInfoReq := model.BusinessAttributeAdd{}
		attributeInfoReq = append(attributeInfoReq, model.BusinessAttribueAddItem{
			ID:     0,
			Weight: attribute.Weight,
			Types:  attribute.Types,
			Name:   attribute.Attributename,
			Desc:   attribute.Description,
		})
		attributeInfoAtlasReq := model.GetAddBusinessAttributesAtlasReq(attributeInfoReq, username, typeName)
		addAtlasAttr, _ := utils.Call("atlas/v2/types/typedefs", username, password, "PUT", map[string]string{"type": "business_metadata"}, attributeInfoAtlasReq)
		var businessRespAtlas model.BMRespAtlas
		fmt.Println(string(addAtlasAttr))
		_ = json.Unmarshal(addAtlasAttr, &businessRespAtlas)
		//mysql
		if len(businessRespAtlas.BusinessMetadataDefs) != 0 {
			for _, business := range businessRespAtlas.BusinessMetadataDefs {
				model.UpdateBusinessVersion(business.Name, useridInt, business.Version)
			}
		}
		for _, attributeInfo := range attributeInfoReq {
			model.UpdateBusinessAttributeInfo(typeName, attributeInfo.Name, useridInt, attributeInfo.Weight, attributeInfo.Desc)
			model.UpdateBusinessAttributeTypes(typeName, attributeInfo.Name, attributeInfo.Types, useridInt, businessRespAtlas.BusinessMetadataDefs[0].GUID)
		}
		if attribute.Weight != oriAttribute.Weight {
			model.AddTypeRecord(useridInt, "Update Business Metadata Attribute Weight", content, oriAttribute.Businessname)
		}
		if len(attribute.Types) != len(oriAttribute.Types) {
			model.AddTypeRecord(useridInt, "Update Business Metadata Attribute Types", content, oriAttribute.Businessname)
		}

	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func DeleteAttr(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	typeName := c.Query("typename")
	guid := c.Query("guid")
	relateid := c.Query("relateid")
	//qualifiedName := c.Query("typeName")
	delAtlasResp := make(map[string]interface{})
	switch typeInfo {
	case "term":
		if strings.Contains(typeName, "@") {
			delGlossaryTerm := []model.DelGlossaryTerm{}
			termName, glossaryName := strings.Split(typeName, "@")[0], strings.Split(typeName, "@")[1]
			termInfo := model.GetTermInfo(termName, glossaryName)
			delGlossaryTerm = append(delGlossaryTerm, model.DelGlossaryTerm{
				GUID:             guid,
				RelationshipGUID: relateid,
			})

			delAtlasTerm, _ := utils.Call("atlas/v2/glossary/terms/"+termInfo.Guid+"/assignedEntities", username, password, "PUT", nil, delGlossaryTerm)
			_ = json.Unmarshal(delAtlasTerm, &delAtlasResp)
		}

	case "classification":
		//http://hadoop102:21000/api/atlas/v2/entity/guid/652e068a-946c-494d-bda8-364f56e38396/classification/%E6%B5%8B%E8%AF%95%E7%B1%BB%E5%9E%8B
		delAtlasClassification, _ := utils.Call("atlas/v2/entity/guid/"+guid+"/classification/"+typeName, username, password, "DELETE", nil, nil)
		_ = json.Unmarshal(delAtlasClassification, &delAtlasResp)
	case "glossary":
		//http://hadoop102:21000/api/atlas/v2/glossary/terms/94d5a2fc-7eb8-480f-87ef-6b50788dd63e/assignedEntities
		glossaryId := model.GetTermGlossaryName(guid, relateid)
		delGlossaryTerm := []model.DelGlossaryTerm{}
		delGlossaryTerm = append(delGlossaryTerm, model.DelGlossaryTerm{
			GUID:             glossaryId,
			RelationshipGUID: relateid,
		})
		delAtlasGlossary, _ := utils.Call("atlas/v2/glossary/term/"+guid, username, password, "DELETE", nil, nil)
		_ = json.Unmarshal(delAtlasGlossary, &delAtlasResp)
		if delAtlasResp["errorMessage"] != "" {
			msg := delAtlasResp["errorMessage"]

			re := regexp.MustCompile(`\{(\d+)\}`)
			match := re.FindStringSubmatch(msg.(string))

			if len(match) > 1 {
				number := match[1]
				delAtlasResp["number"] = number
			}
		} else {
			//mysql删除
			model.DelTerm(guid, relateid)
			model.AddTypeRecord(useridInt, "Delete All Term", "", typeName)
		}
	}
	typeDelMap := make(map[string]interface{})
	typeDelMap["info"] = delAtlasResp
	typeDelMap["status"] = utils.SUCCESS
	typeDelMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, typeDelMap)
}

func DeleteType(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typename := c.Query("typename")
	typeInfo := c.Query("type")
	switch typeInfo {
	case "classification":
		classificationInfo := model.GetClassificatioInfo(typename)
		if classificationInfo.Userid != useridInt && model.GetUserRole(useridInt) == 3 {
			c.JSON(
				http.StatusOK, map[string]interface{}{
					"status":  utils.ERROR_USER_AUTH_NOT_ENOUGH,
					"message": utils.GetErrMsg(utils.ERROR_USER_AUTH_NOT_ENOUGH),
				})
		} else {
			//http://hadoop102:21000/api/atlas/v2/types/typedef/name/分类4
			_, err := utils.Call("atlas/v2/types/typedef/name/"+typename, username, password, "DELETE", nil, nil)
			if err != nil {
				c.JSON(
					http.StatusOK, map[string]interface{}{
						"status":  utils.ERROR,
						"message": utils.GetErrMsg(utils.ERROR),
					})
			} else {
				c.JSON(
					http.StatusOK, map[string]interface{}{
						"status":  utils.SUCCESS,
						"message": utils.GetErrMsg(utils.SUCCESS),
					})
				model.AddTypeRecord(useridInt, "Remove Type", "Remove classification name:"+typename, classificationInfo.Classificationname)
			}
		}
	}

}
