package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/utils"
)

func AddUserLabel(c *gin.Context) {
	userLabelReq := make(map[string]interface{})
	userLabelReqJson, err := c.GetRawData()
	err = json.Unmarshal(userLabelReqJson, &userLabelReq)
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
	//http://hadoop102:21000/api/atlas/v2/entity
	userLabelAddRespString, _ := utils.Call("atlas/v2/entity", username, password, "POST", nil, userLabelReq)
	userLabelAddResp := make(map[string]interface{})
	userLabelAddResp["label"] = string(userLabelAddRespString)
	userLabelAddResp["status"] = utils.SUCCESS
	userLabelAddResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, userLabelAddResp)

}
