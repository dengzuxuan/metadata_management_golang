package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type auditType struct {
	Times     string      `json:"times"`
	Timestamp int64       `json:"timestamp"`
	Action    string      `json:"action"`
	Details   string      `json:"details"`
	Entity    interface{} `json:"entity"`
	Type      interface{} `json:"type"`
}
type auditTypes []auditType

var guids []string
var wg sync.WaitGroup
var auidtTypes auditTypes

func (s auditTypes) Len() int {
	return len(s)
}
func (s auditTypes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s auditTypes) Less(i, j int) bool {
	return s[i].Timestamp < s[j].Timestamp
}
func LoginUser(c *gin.Context) {
	var userLoginInfo model.LoginUser
	_ = c.ShouldBindJSON(&userLoginInfo)
	username := userLoginInfo.Username
	password := userLoginInfo.Password
	fmt.Println("userLoginInfo", userLoginInfo)
	token := utils.RandString(20)
	if len(username) < 4 || len(username) > 12 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_USERNAME_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_USERNAME_WRONG),
		})
		return
	}
	if len(password) < 5 || len(password) > 20 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_PASSWORD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_PASSWORD_WRONG),
		})
		return
	}
	code, id := model.CheckLogin(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"id":      id,
		"token":   token,
	})
}
func GetUser(c *gin.Context) {
	userid := c.GetHeader("user_id")
	userIdInt, _ := strconv.Atoi(userid)
	user := model.GetUserInfos(userIdInt)
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func GetSingleUser(c *gin.Context) {
	userid := c.Query("userid")
	userIdInt, _ := strconv.Atoi(userid)
	user := model.GetUserInfos(userIdInt)
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}
func GetSingleUserId(c *gin.Context) {
	username := c.Query("username")
	user := model.GetUserId(username)
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}
func GetUserAudit(c *gin.Context) {
	guids = []string{}
	auidtTypes = auditTypes{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	userid := c.Query("userid")
	useridInt, _ := strconv.Atoi(userid)
	entritySting, _ := utils.Call("atlas/admin/metrics", username, password, "GET", nil, nil)
	entityMap := make(map[string]interface{})
	json.Unmarshal(entritySting, &entityMap)
	entity := entityMap["data"].(map[string]interface{})["entity"].(map[string]interface{})["entityActive"].(map[string]interface{})
	wg.Add(len(entity))
	for typeName, _ := range entity {
		go getAllGuids(username, password, typeName)
	}
	wg.Wait()
	guidsMap := make(map[string]bool)
	guidsArray := []string{}
	for _, guid := range guids {
		if _, ok := guidsMap[guid]; !ok {
			guidsArray = append(guidsArray, guid)
		}
		guidsMap[guid] = true
	}

	wg.Add(len(guidsArray))
	userFindName, _ := model.GetUserInfo(useridInt)
	for _, guid := range guidsArray {
		go getUserAudit(username, password, userFindName, guid)
	}
	wg.Wait()
	auditsMap := make(map[string]interface{})
	auditsMap["status"] = utils.SUCCESS
	auditsMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	sort.Sort(sort.Reverse(auidtTypes))
	auditsMap["auditInfos"] = auidtTypes
	auditsMap["num"] = len(auidtTypes)
	c.JSON(http.StatusOK, auditsMap)
}

func GetUserCollect(c *gin.Context) {
	CollectInfos := []model.CollectInfoType{}
	username := c.GetHeader("username")
	password := c.GetHeader("password")
	userid := c.Query("userid")
	useridInt, _ := strconv.Atoi(userid)
	collects := model.GetCollect(useridInt)
	for _, collect := range collects {
		switch collect.Type {
		case "entity":
			entity := map[string]interface{}{}
			//entity:=model.AtlasEntityInfo{}
			entityJson, _ := utils.Call("atlas/v2/entity/guid/"+collect.Collectguid, username, password, "GET", nil, nil)
			_ = json.Unmarshal(entityJson, &entity)
			entityInfo := entity["entity"].(map[string]interface{})
			createUserName := entityInfo["createdBy"].(string)
			updateUserName := entityInfo["updatedBy"].(string)
			createUserInfo := model.GetUserInfos(model.GetUserId(createUserName))
			updateUserInfo := model.GetUserInfos(model.GetUserId(updateUserName))
			CollectInfos = append(CollectInfos, model.CollectInfoType{
				Id:               collect.Id,
				Collectguid:      collect.Collectguid,
				Createtime:       collect.Createtime,
				EntityName:       collect.Collectname,
				Typename:         "",
				Type:             "entity",
				Created:          entityInfo["createTime"].(float64),
				Updated:          entityInfo["updateTime"].(float64),
				CreateUserId:     createUserInfo.Id,
				CreateUserName:   createUserInfo.Username,
				CreateUserAvatar: createUserInfo.Avatar,
				CreateUserRole:   createUserInfo.RoleInfo,
				UpdateUserId:     updateUserInfo.Id,
				UpdateUserName:   updateUserInfo.Username,
				UpdateUserAvatar: updateUserInfo.Avatar,
				UpdateUserRole:   updateUserInfo.RoleInfo,
				Desc:             "暂无",
			})
		case "type":
			typeName := collect.Typename
			findName := collect.Findname
			switch findName {
			case "classification":
				typeInfos := model.GetClassificatioInfo(typeName)
				createUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.Username))
				updateUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.UpdateUsername))
				CollectInfos = append(CollectInfos, model.CollectInfoType{
					Id:               collect.Id,
					Collectguid:      collect.Collectguid,
					Createtime:       collect.Createtime,
					EntityName:       collect.Collectname,
					Typename:         findName,
					Type:             "type",
					Created2:         utils.ChangeShowType(typeInfos.Createtime),
					Updated2:         utils.ChangeShowType(typeInfos.Updatetime),
					CreateUserId:     createUserInfo.Id,
					CreateUserName:   createUserInfo.Username,
					CreateUserAvatar: createUserInfo.Avatar,
					CreateUserRole:   createUserInfo.RoleInfo,
					UpdateUserId:     updateUserInfo.Id,
					UpdateUserName:   updateUserInfo.Username,
					UpdateUserAvatar: updateUserInfo.Avatar,
					UpdateUserRole:   updateUserInfo.RoleInfo,
					Desc:             typeInfos.Description,
				})
			case "term":
				if strings.Contains(typeName, "@") {
					termName, glossaryName := strings.Split(typeName, "@")[0], strings.Split(typeName, "@")[1]
					typeInfos := model.GetTermInfo(termName, glossaryName)
					createUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.Username))
					updateUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.UpdateUsername))
					CollectInfos = append(CollectInfos, model.CollectInfoType{
						Id:               collect.Id,
						Collectguid:      collect.Collectguid,
						Createtime:       collect.Createtime,
						EntityName:       collect.Collectname,
						Typename:         findName,
						Type:             "type",
						Created2:         utils.ChangeShowType(typeInfos.Createtime),
						Updated2:         utils.ChangeShowType(typeInfos.Updatetime),
						CreateUserId:     createUserInfo.Id,
						CreateUserName:   createUserInfo.Username,
						CreateUserAvatar: createUserInfo.Avatar,
						CreateUserRole:   createUserInfo.RoleInfo,
						UpdateUserId:     updateUserInfo.Id,
						UpdateUserName:   updateUserInfo.Username,
						UpdateUserAvatar: updateUserInfo.Avatar,
						UpdateUserRole:   updateUserInfo.RoleInfo,
						Desc:             typeInfos.Longdescription,
					})
				}
			case "glossary":
				typeInfos := model.GetGlossaryInfo(typeName)
				createUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.Username))
				updateUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.UpdateUsername))
				CollectInfos = append(CollectInfos, model.CollectInfoType{
					Id:               collect.Id,
					Collectguid:      collect.Collectguid,
					Createtime:       collect.Createtime,
					EntityName:       collect.Collectname,
					Typename:         findName,
					Type:             "type",
					Created2:         utils.ChangeShowType(typeInfos.Createtime),
					Updated2:         utils.ChangeShowType(typeInfos.Updatetime),
					CreateUserId:     createUserInfo.Id,
					CreateUserName:   createUserInfo.Username,
					CreateUserAvatar: createUserInfo.Avatar,
					CreateUserRole:   createUserInfo.RoleInfo,
					UpdateUserId:     updateUserInfo.Id,
					UpdateUserName:   updateUserInfo.Username,
					UpdateUserAvatar: updateUserInfo.Avatar,
					UpdateUserRole:   updateUserInfo.RoleInfo,
					Desc:             typeInfos.Longdescription,
				})
			case "business":
				typeInfos := model.GetClassificatioInfo(typeName)
				createUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.Username))
				updateUserInfo := model.GetUserInfos(model.GetUserId(typeInfos.UpdateUsername))
				CollectInfos = append(CollectInfos, model.CollectInfoType{
					Id:               collect.Id,
					Collectguid:      collect.Collectguid,
					Createtime:       collect.Createtime,
					EntityName:       collect.Collectname,
					Typename:         findName,
					Type:             "type",
					Created2:         utils.ChangeShowType(typeInfos.Createtime),
					Updated2:         utils.ChangeShowType(typeInfos.Updatetime),
					CreateUserId:     createUserInfo.Id,
					CreateUserName:   createUserInfo.Username,
					CreateUserAvatar: createUserInfo.Avatar,
					CreateUserRole:   createUserInfo.RoleInfo,
					UpdateUserId:     updateUserInfo.Id,
					UpdateUserName:   updateUserInfo.Username,
					UpdateUserAvatar: updateUserInfo.Avatar,
					UpdateUserRole:   updateUserInfo.RoleInfo,
					Desc:             typeInfos.Description,
				})
			}
		}
	}
	collectsMap := make(map[string]interface{})
	collectsMap["collectInfos"] = CollectInfos
	collectsMap["status"] = utils.SUCCESS
	collectsMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, collectsMap)
}

func UploadAvatar(c *gin.Context) {
	userid := c.GetHeader("user_id")
	uploadType := c.GetHeader("upload_info")
	file, err := c.FormFile("file") // 获取上传的文件
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  utils.ERROR_UPLOAD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_UPLOAD_WRONG),
		})
		return
	}
	switch uploadType {
	case "avatar":
		err = utils.UploadOss("avatar/"+userid+"_"+file.Filename, file)
	case "bg":
		err = utils.UploadOss("bg/"+userid+"_"+file.Filename, file)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_UPLOAD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_UPLOAD_WRONG),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func UpdateUserInfo(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	userUpdateJson := c.Query("content")
	userInfo := model.UpdateUserInfo{}
	json.Unmarshal([]byte(userUpdateJson), &userInfo)
	oriavatar := userInfo.Avatar
	if len(userInfo.Newavatar) != 0 {
		oriavatar = "http://metadata-manager.oss-cn-beijing.aliyuncs.com/avatar/" + userid + "_" + userInfo.Newavatar
	}
	oribg := userInfo.Background
	if len(userInfo.Newbg) != 0 {
		oribg = "http://metadata-manager.oss-cn-beijing.aliyuncs.com/bg/" + userid + "_" + userInfo.Newbg
	}
	updateUser := model.User{
		Avatar:     oriavatar,
		Department: userInfo.Department,
		Background: oribg,
		Telephone:  userInfo.Telephone,
		Email:      userInfo.Email,
		Address:    userInfo.Address,
		Place:      userInfo.Place,
		Statement:  userInfo.Statement,
		Male:       userInfo.Male,
	}
	model.UpdateUser(useridInt, updateUser)
	fmt.Println(userid, userUpdateJson)
}

func DeleteCollect(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	model.DeleteCollect(idInt)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func getAllGuids(username, password, typeName string) {
	body := map[string]interface{}{
		"excludeDeletedEntities":          true,
		"includeSubClassifications":       true,
		"includeSubTypes":                 true,
		"includeClassificationAttributes": true,
		"entityFilters":                   nil,
		"tagFilters":                      nil,
		"attributes":                      nil,
		"offset":                          0,
		"limit":                           999,
		"classification":                  nil,
		"termName":                        nil,
		"typeName":                        typeName,
	}
	typeDetail := model.AtlasFindType{}
	//typeDetailsMap := make(map[string]interface{})
	typeDetails, _ := utils.Call("atlas/v2/search/basic", username, password, "POST", nil, body)
	_ = json.Unmarshal(typeDetails, &typeDetail)
	for _, detail := range typeDetail.Entities {
		guids = append(guids, detail.GUID)
	}
	wg.Done()
}
func getUserAudit(username, password, userFindName, guid string) {
	otherInfos := model.AtlasAudit{}
	otherInfoJson, _ := utils.Call("atlas/v2/entity/"+guid+"/audit", username, password, "GET", nil, nil)
	_ = json.Unmarshal(otherInfoJson, &otherInfos)
	for _, info := range otherInfos {
		if info.User == userFindName {
			timestamp := int64(info.Timestamp) / 1000 // 转换为以秒为单位的时间戳
			t := time.Unix(timestamp, 0)

			// 将time.Time对象格式化为指定格式的字符串
			formattedTime := t.Format("2006-01-02 15:04")
			auidtTypes = append(auidtTypes, auditType{
				Times:     formattedTime,
				Timestamp: info.Timestamp,
				Action:    info.Action,
				Details:   info.Details,
				//todo:补充enetity的信息
				Entity: info.EntityID,
				Type:   info.Type,
			})
		}
	}
	wg.Done()
}
