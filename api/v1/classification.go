package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
)

func AddClassificationInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	avatar := model.GetUserAvatar(useridInt)
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
			model.AddClassification(classification.Name, useridInt, username, avatar, len(classification.AttributeDefs), classification.Description, classificationRespAtlas.ClassificationDefs[0].GUID)
			for _, attribute := range classification.AttributeDefs {
				model.AddClassificationAttribute(classification.Name, attribute.Name, useridInt, username, avatar, attribute.Description, classificationRespAtlas.ClassificationDefs[0].GUID)
			}
		}
	}
	addAtlasClassificationResp["status"] = utils.SUCCESS
	addAtlasClassificationResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addAtlasClassificationResp)
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
