package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/utils"
)

func Addlabel(c *gin.Context) {
	//http://hadoop102:21000/api/atlas/v2/entity/guid/cc26704e-097b-46f0-8b60-693c2c09aa74/labels
	//[
	//  "111"
	//]
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	guid := c.Query("guid")
	label := c.Query("label")
	labels := []string{}
	_ = json.Unmarshal([]byte(label), &labels)
	labelAddRespString, _ := utils.Call("atlas/v2/entity/guid/"+guid+"/labels", username, password, "POST", nil, labels)
	labelAddResp := make(map[string]interface{})
	labelAddResp["label"] = string(labelAddRespString)
	labelAddResp["status"] = utils.SUCCESS
	labelAddResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, labelAddResp)

}
