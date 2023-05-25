package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
	"time"
)

func AddClassificationInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	//avatar := model.GetUserAvatar(useridInt)
	var classificationReq model.ClassificationReqAtlas
	_ = c.ShouldBindJSON(&classificationReq)
	//atlas
	query := map[string]string{"type": "classification"}
	addAtlasClassification, _ := utils.Call("atlas/v2/types/typedefs", username, password, "POST", query, classificationReq)
	addAtlasClassificationResp := make(map[string]interface{})
	_ = json.Unmarshal(addAtlasClassification, &addAtlasClassificationResp)

	//mysql
	var classificationRespAtlas model.ClassificationRespAtlas
	_ = json.Unmarshal(addAtlasClassification, &classificationRespAtlas)
	if len(classificationRespAtlas.ClassificationDefs) != 0 {
		for _, classification := range classificationReq.ClassificationDefs {
			model.AddClassification(classification.Name, useridInt, len(classification.AttributeDefs), classificationRespAtlas.ClassificationDefs[0].Version, classification.Description, classificationRespAtlas.ClassificationDefs[0].GUID)
			for _, attribute := range classification.AttributeDefs {
				model.AddClassificationAttribute(classification.Name, attribute.Name, useridInt, attribute.Description, classificationRespAtlas.ClassificationDefs[0].GUID)
			}
		}
	}
	addAtlasClassificationResp["status"] = utils.SUCCESS
	addAtlasClassificationResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addAtlasClassificationResp)
}

func UpdateClassificatioInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	name := c.Query("name")
	desc := c.Query("desc")
	attributeInfoJson := c.Query("attribute")
	updateAttribute := model.UpdateAttribue{}
	_ = json.Unmarshal([]byte(attributeInfoJson), &updateAttribute)
	oriClassification := model.GetClassificatioInfo(name)
	//atlas
	createname, _ := model.GetUserInfo(oriClassification.Userid)
	attributeUpdate := []model.AttributeDefType{}
	for _, attributeInfo := range updateAttribute {
		attributeUpdate = append(attributeUpdate, model.AttributeDefType{
			Name:           attributeInfo.Name,
			TypeName:       "string",
			IsOptional:     true,
			Cardinality:    "SINGLE",
			ValuesMinCount: 0,
			ValuesMaxCount: 1,
			IsUnique:       false,
			IsIndexable:    true,
		})
	}
	classificationUpdate := model.ClassificationDef{
		Category:      "CLASSIFICATION",
		GUID:          oriClassification.Guid,
		CreatedBy:     createname,
		UpdatedBy:     username,
		CreateTime:    time.Now().Unix(),
		UpdateTime:    utils.TimeStringToUnix(oriClassification.Createtime),
		Version:       1,
		Name:          name,
		Description:   desc,
		TypeVersion:   "1.0",
		AttributeDefs: attributeUpdate,
		SuperTypes:    nil,
		EntityTypes:   nil,
		SubTypes:      nil,
	}
	updateClassification := map[string]interface{}{}
	updateClassificationJson, _ := utils.Call("atlas/v2/types/typedefs", username, password, "PUT", map[string]string{
		"type": "classification",
	}, classificationUpdate)
	_ = json.Unmarshal(updateClassificationJson, &updateClassification)
	updateClassificationString := string(updateClassificationJson)
	fmt.Println(updateClassificationString)
	//mysql
	if oriClassification.Description != desc {
		model.UpdateClassification(oriClassification.Classificationname, desc, useridInt)
		model.AddTypeRecord(useridInt, "Update Classification Description", oriClassification.Description+"->"+desc)
	}
	for _, attribute := range updateAttribute {
		attributeName := attribute.Name
		attributeDesc := attribute.Description
		action := model.CheckClassificationAttribute(attributeName, name, useridInt, attributeDesc, oriClassification.Guid)
		if action != "none" {
			model.AddTypeRecord(useridInt, action, attributeInfoJson)
		}
	}
	updateClassification["status"] = utils.SUCCESS
	updateClassification["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, updateClassification)
}

func DeleteClassification(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")
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
				model.AddTypeRecord(useridInt, "Remove Type", "Remove classification name:"+typename)
			}
		}
	}

}

func GetClassificationInfo(c *gin.Context) {
	classification := model.GetAllClassification()
	classificationMap := make(map[string]interface{})
	classificationMap["allClassification"] = classification
	classificationMap["status"] = utils.SUCCESS
	classificationMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, classificationMap)
}

func GetClassificationName(c *gin.Context) {
	guid := c.Query("guid")
	classificationInfo := model.GetClassificationName(guid)
	classificationNameMap := make(map[string]interface{})
	classificationNameMap["classificationName"] = classificationInfo.Classificationname
	classificationNameMap["status"] = utils.SUCCESS
	classificationNameMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, classificationNameMap)
}

func GetClassificationAttributeInfo(c *gin.Context) {
	guid := c.Query("guid")
	classificationAttribute := model.GetClassificationAttribute(guid)
	classificationAttributeMap := make(map[string]interface{})
	classificationAttributeMap["classificationAttribute"] = classificationAttribute
	classificationAttributeMap["status"] = utils.SUCCESS
	classificationAttributeMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, classificationAttributeMap)
}

// /v2/entity/bulk/classification
func EntityAddClassification(c *gin.Context) {
	classificationReq := make(map[string]interface{})
	classificationReqJson, err := c.GetRawData()
	err = json.Unmarshal(classificationReqJson, &classificationReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	classificationSting, _ := utils.Call("atlas/v2/entity/bulk/classification", username, password, "POST", nil, classificationReq)
	classificationType := make(map[string]interface{})
	_ = json.Unmarshal(classificationSting, &classificationType)
	classificationType["status"] = utils.SUCCESS
	classificationType["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, classificationType)
}
