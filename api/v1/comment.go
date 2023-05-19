package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"others-part/model"
	"others-part/utils"
	"strconv"
)

func GetAllComments(c *gin.Context) {
	userid := c.GetHeader("user_id")
	useridInt, _ := strconv.Atoi(userid)
	guid := c.Query("guid")
	otherid := c.Query("otherid")
	commentinfos := model.GetCommentInfo(guid, otherid, useridInt)
	commentinfosMap := make(map[string]interface{})
	commentinfosMap["comments"] = commentinfos
	commentinfosMap["status"] = utils.SUCCESS
	commentinfosMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, commentinfosMap)
}

func AddComment(c *gin.Context) {
	userid := c.GetHeader("user_id")
	var commentReq model.CommentInfoReq
	_ = c.ShouldBindJSON(&commentReq)
	retCode := model.AddCommentInfo(userid, commentReq)
	commentResp := make(map[string]interface{})
	commentResp["status"] = retCode
	commentResp["message"] = utils.GetErrMsg(retCode)
	c.JSON(http.StatusOK, commentResp)
}
