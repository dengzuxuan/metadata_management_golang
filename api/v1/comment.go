package v1

import (
	"fmt"
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
	commentname := c.Query("commentname")
	fmt.Println("commentname", commentname)
	commentinfos := model.GetCommentInfo(guid, commentname, useridInt)
	commentinfosMap := make(map[string]interface{})
	commentinfosMap["comments"] = commentinfos
	commentinfosMap["status"] = utils.SUCCESS
	commentinfosMap["message"] = utils.GetErrMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, commentinfosMap)
}

func AddComment(c *gin.Context) {
	username := c.GetHeader("username")
	password1, _ := c.Get("password")
	password := password1.(string)
	userid := c.GetHeader("user_id")
	var commentReq model.CommentInfoReq
	_ = c.ShouldBindJSON(&commentReq)
	retCode := model.AddCommentInfo(userid, commentReq, username, password)
	commentResp := make(map[string]interface{})
	commentResp["status"] = retCode
	commentResp["message"] = utils.GetErrMsg(retCode)
	c.JSON(http.StatusOK, commentResp)
}

func AddCommentLike(c *gin.Context) {
	userid := c.GetHeader("user_id")
	commentid := c.Query("comment_id")
	retCode := model.AddCommentLike(userid, commentid)
	commentResp := make(map[string]interface{})
	commentResp["status"] = retCode
	commentResp["message"] = utils.GetErrMsg(retCode)
	c.JSON(http.StatusOK, commentResp)
}

func DelCommentLike(c *gin.Context) {
	userid := c.GetHeader("user_id")
	commentid := c.Query("comment_id")
	retCode := model.DelCommentLike(userid, commentid)
	commentResp := make(map[string]interface{})
	commentResp["status"] = retCode
	commentResp["message"] = utils.GetErrMsg(retCode)
	c.JSON(http.StatusOK, commentResp)
}

func DelComment(c *gin.Context) {
	userid := c.GetHeader("user_id")
	commentid := c.Query("comment_id")
	retCode := model.DelComment(userid, commentid)
	commentResp := make(map[string]interface{})
	commentResp["status"] = retCode
	commentResp["message"] = utils.GetErrMsg(retCode)
	c.JSON(http.StatusOK, commentResp)
}
