package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
)

func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) < 4 || len(username) > 12 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_USERNAME_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_USERNAME_WRONG),
		})
		return
	}
	if len(password) < 6 || len(password) > 20 {
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
	})
}
