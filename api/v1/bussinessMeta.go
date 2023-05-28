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

func CreateBusinessInfo(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	//avatar := model.GetUserAvatar(useridInt)
	var businessReq model.AtlasBusinessMeta
	_ = c.ShouldBindJSON(&businessReq)
	//atlas
	//http://hadoop102:21000/api/atlas/v2/types/typedefs?type=business_metadata
	query := map[string]string{"type": "business_metadata"}
	addAtlasBusiness, _ := utils.Call("atlas/v2/types/typedefs", username, password, "POST", query, businessReq)
	addAtlasBusinessResp := make(map[string]interface{})
	_ = json.Unmarshal(addAtlasBusiness, &addAtlasBusinessResp)

	//mysql
	var businessRespAtlas model.BMRespAtlas
	_ = json.Unmarshal(addAtlasBusiness, &businessRespAtlas)
	if len(businessRespAtlas.BusinessMetadataDefs) != 0 {
		for _, business := range businessReq.BusinessMetadataDefs {
			model.AddBusinessMetaInfo(business.Name, business.CreatedBy, business.Version, business.Description, businessRespAtlas.BusinessMetadataDefs[0].GUID)
			for _, attribute := range business.AttributeDefs {
				model.AddBusinessAttributeInfo(business.Name, attribute.Name, business.CreatedBy, attribute.SearchWeight, attribute.Desc, businessRespAtlas.BusinessMetadataDefs[0].GUID)
				types := []string{}
				json.Unmarshal([]byte(attribute.Options.ApplicableEntityTypes), &types)
				for _, typeName := range types {
					model.AddBusinessAttributeTypeInfo(typeName, business.Name, attribute.Name, business.CreatedBy, businessRespAtlas.BusinessMetadataDefs[0].GUID)
				}
			}
			businessJson, _ := json.Marshal(business)
			model.AddTypeRecord(useridInt, "Business Meatedata Create", string(businessJson), business.Name)
		}
	}

	addAtlasBusinessResp["status"] = utils.SUCCESS
	addAtlasBusinessResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addAtlasBusinessResp)
}

func EntityAddBusiness(c *gin.Context) {
	//http://hadoop102:21000/api/atlas/v2/entity/guid/0efaaec3-cf2a-44a8-a9cc-c7e5c67a4125/businessmetadata?isOverwrite=true
	businessReq := make(map[string]interface{})
	businessReqJson, err := c.GetRawData()
	err = json.Unmarshal(businessReqJson, &businessReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	classificationSting, _ := utils.Call("atlas/v2/entity/guid/classification", username, password, "POST", nil, businessReq)
	classificationType := make(map[string]interface{})
	_ = json.Unmarshal(classificationSting, &classificationType)
	classificationType["status"] = utils.SUCCESS
	classificationType["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, classificationType)
}

func GetBusinessTypeInfos(c *gin.Context) {
	typename := c.Query("typename")
	guid := c.Query("guid")
	typeInfos := model.GetBusinessInfoByType(typename, guid)
	typeInfosMap := make(map[string]interface{})
	typeInfosMap["typeInfos"] = typeInfos
	typeInfosMap["status"] = utils.SUCCESS
	typeInfosMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, typeInfosMap)
}

func GetGuidBusiness(c *gin.Context) {
	guid := c.Query("guid")
	businessInfo := model.GetEntityBusinessInfos(guid)
	businessInfoMap := make(map[string]interface{})
	businessInfoMap["businessInfos"] = businessInfo
	businessInfoMap["status"] = utils.SUCCESS
	businessInfoMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, businessInfoMap)
}

func GetBusinessInfos(c *gin.Context) {
	guid := c.Query("guid")

	businessInfos := model.GetEntityBusinessInfos(guid)
	businessMap := make(map[string]interface{})
	businessMap["businessInfos"] = businessInfos
	businessMap["status"] = utils.SUCCESS
	businessMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, businessMap)
}

func GetBusinessNameInfos(username, password, businessName string, pageLimit, pageOffset int) map[string]interface{} {
	entityIds := make(map[string]bool)
	limitFlag := 0
	findInfos := make(map[string]interface{})
	findEntityInfos := []map[string]interface{}{}
	entityGuids := model.GetBusinessEntityIds(businessName)
	for i, entityInfo := range entityGuids {
		if i >= pageOffset {
			limitFlag++
			entity := make(map[string]interface{})
			if _, ok := entityIds[entityInfo.Entityguid]; ok {
				continue
			}
			entityIds[entityInfo.Entityguid] = true
			entityJson, _ := utils.Call("atlas/v2/entity/guid/"+entityInfo.Entityguid, username, password, "GET", nil, nil)
			_ = json.Unmarshal(entityJson, &entity)
			if entity["entity"] != nil {
				findEntityInfos = append(findEntityInfos, entity["entity"].(map[string]interface{}))
			}
			if limitFlag > pageLimit {
				break
			}
		}
	}
	findInfos["entities"] = findEntityInfos
	return findInfos
}

func AddBusinessInfo(c *gin.Context) {

	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	guid := c.Query("guid")
	oriBmReqJson := c.Query("oribm")
	addBmReqJson := c.Query("addbm")
	var oriBmReq model.OriBmReq
	var addBmReq model.AddBmReq
	json.Unmarshal([]byte(oriBmReqJson), &oriBmReq)
	json.Unmarshal([]byte(addBmReqJson), &addBmReq)
	//{
	//	"bm002": {
	//	"a1": "test值",
	//	"a10": "a10v"
	//},
	//	"bm008": {
	//	"a2": "asa"
	//}
	//}

	addBm := make(map[string]map[string]string)
	for _, ori := range oriBmReq {
		if addBm[ori.BusinessName] == nil {
			addBm[ori.BusinessName] = make(map[string]string)
		}
		for _, attribute := range ori.AttributeInfos {
			addBm[ori.BusinessName][attribute.AttributeName] = attribute.AttributeValue
		}
	}
	for _, add := range addBmReq {
		if addBm[add.AttributeName.BusinessName] == nil {
			addBm[add.AttributeName.BusinessName] = make(map[string]string)
		}
		addBm[add.AttributeName.BusinessName][add.AttributeName.AttributeName] = add.AttributeValue
	}
	oriEntityInfo := model.GetEntityBusinessInfos(guid)
	if len(oriBmReq) == 0 {
		for _, info := range oriEntityInfo {
			addBm[info.BusinessName] = map[string]string{}
		}
	}
	//mysql 更新 删除
	for _, oriInfo := range oriEntityInfo {
		if _, ok := addBm[oriInfo.BusinessName]; !ok {
			//删除该guid下的business
			model.DeleteEntityBusinessInfo(guid, oriInfo.BusinessName)
			continue
		}
		for _, oriAttributeinfo := range oriInfo.AttributeInfos {
			if _, ok := addBm[oriInfo.BusinessName][oriAttributeinfo.AttributeName]; !ok {
				//删除该guid下的attribute
				model.DeleteEntityAttributeInfo(guid, oriInfo.BusinessName, oriAttributeinfo.AttributeName)
				continue
			}
			if oriAttributeinfo.AttributeValue != addBm[oriInfo.BusinessName][oriAttributeinfo.AttributeName] {
				//更新该guid下的attribute value
				model.UpdateEntityAttributeInfo(useridInt, guid, oriInfo.BusinessName, oriAttributeinfo.AttributeName, addBm[oriInfo.BusinessName][oriAttributeinfo.AttributeName])
				continue
			}
		}
	}
	//mysql 新建
	for _, addInfo := range addBmReq {
		model.AddEntityBusinessInfo(useridInt, guid, addInfo.AttributeName.BusinessName, addInfo.AttributeName.AttributeName, addInfo.AttributeValue)
	}

	//atlas更新
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	//http://hadoop102:21000/api/atlas/v2/entity/guid/0efaaec3-cf2a-44a8-a9cc-c7e5c67a4125/businessmetadata?isOverwrite=true
	addBusinessSting, _ := utils.Call("atlas/v2/entity/guid/"+guid+"/businessmetadata", username, password, "POST", map[string]string{"isOverwrite": "true"}, addBm)
	addBusinessResp := make(map[string]interface{})
	_ = json.Unmarshal(addBusinessSting, &addBusinessResp)
	addBusinessResp["status"] = utils.SUCCESS
	addBusinessResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, addBusinessResp)
}
