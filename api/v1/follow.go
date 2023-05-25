package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
)

func AddFollow(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	touserid := c.Query("touserid")
	touseridInt, _ := strconv.Atoi(touserid)
	model.AddFollowRecord(useridInt, touseridInt)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func DelFollow(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	touserid := c.Query("touserid")
	touseridInt, _ := strconv.Atoi(touserid)
	model.DelFollowRecord(useridInt, touseridInt)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func CheckFollow(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	touserid := c.Query("touserid")
	touseridInt, _ := strconv.Atoi(touserid)
	check := model.CheckFollow(useridInt, touseridInt)
	fmt.Println("checkinfo:", check)
	c.JSON(http.StatusOK, map[string]interface{}{
		"checked": check,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}

func GetAllFollow(c *gin.Context) {
	userid := c.Query("userid")
	useridInt, _ := strconv.Atoi(userid)
	myfollow, followmy := model.GetAllFollowInfos(useridInt)
	c.JSON(
		http.StatusOK, map[string]interface{}{
			"followmy": followmy,
			"myfollow": myfollow,
			"status":   utils.SUCCESS,
			"message":  utils.GetErrMsg(utils.SUCCESS),
		})
}
