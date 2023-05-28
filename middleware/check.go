package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		passwordCrypto := c.Query("crypto")
		userid := c.GetHeader("user_id")
		useridInt, _ := strconv.Atoi(userid)
		passwordReal := model.GetUserInfos(useridInt).Password
		if !utils.ComparePwd(passwordReal, passwordCrypto) && passwordCrypto != "" {
			c.JSON(http.StatusOK, gin.H{
				"status":  utils.ERROR_USER_PASSWORD_WRONG,
				"message": utils.GetErrMsg(utils.ERROR_USER_PASSWORD_WRONG),
			})
			c.Abort()
			return
		} else {
			c.Set("password", passwordReal)
		}
	}
}
