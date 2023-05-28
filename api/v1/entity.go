package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/utils"
)

func SearchNodeResult(c *gin.Context) {
	guid := c.Query("guid")
	queryParams := make(map[string]string)
	queryParams["guid"] = guid
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)

	entity := map[string]interface{}{}
	entityJson, _ := utils.Call("atlas/v2/entity/guid/"+guid, username, password, "GET", nil, nil)
	_ = json.Unmarshal(entityJson, &entity)
	entity["status"] = utils.SUCCESS
	entity["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, entity)
}
