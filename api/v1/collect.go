package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
)

func AddCollect(c *gin.Context) {
	userid := c.GetHeader("user_id")
	touserName := c.Query("tousername")
	useridInt, _ := strconv.Atoi(userid)
	touseridInt := model.GetUserId(touserName)
	guid := c.Query("guid")
	name := c.Query("name")
	desc := c.Query("desc")
	typeInfo := c.Query("type") //entity or type

	typename := c.Query("typename")
	findname := c.Query("findname")
	model.AddCollect(useridInt, touseridInt, guid, name, typename, findname, desc, typeInfo)
	collectResp := make(map[string]interface{})
	collectResp["status"] = utils.SUCCESS
	collectResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, collectResp)
}

func CheckCollect(c *gin.Context) {
	collected := false
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	switch typeInfo {
	case "entity":
		guid := c.Query("guid")
		collected = model.CheckEntityCollect(useridInt, guid)
	case "type":
		typename := c.Query("typename")
		findname := c.Query("findname")
		collected = model.CheckTypeCollect(useridInt, typename, findname)
	}
	checkcollectResp := make(map[string]interface{})
	checkcollectResp["check"] = collected
	checkcollectResp["status"] = utils.SUCCESS
	checkcollectResp["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, checkcollectResp)

}
func DeleteSingleCollect(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	typeInfo := c.Query("type")
	switch typeInfo {
	case "entity":
		guid := c.Query("guid")
		model.DeleteEntityCollect(useridInt, guid)
	case "type":
		typename := c.Query("typename")
		findname := c.Query("findname")
		model.DeleteTypeCollect(useridInt, typename, findname)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}
