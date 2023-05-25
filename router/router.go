package router

import (
	"github.com/gin-gonic/gin"
	v1 "others-part/api/v1"
	"others-part/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	user := r.Group("api/v1/user")
	{

		user.POST("/login", v1.LoginUser)
		user.GET("/info", v1.GetUser)
		user.GET("/userinfo", v1.GetSingleUser)
		user.GET("/getuserid", v1.GetSingleUserId)
		user.GET("/useraudits", v1.GetUserAudit)
		user.GET("/usercollect", v1.GetUserCollect)
		user.POST("/deletecollect", v1.DeleteCollect)
		user.POST("/uploadavatar", v1.UploadAvatar)
		user.POST("/updateinfo", v1.UpdateUserInfo)
	}
	atlas := r.Group("/api/v1/atlas")
	{
		atlas.GET("/searchpre", v1.SearchPre)
		atlas.GET("/searchresult", v1.SearchResult)
		atlas.GET("/types/entity", v1.TypeEntity)
		atlas.GET("/types/classification", v1.TypeClassification)
		atlas.GET("/types/glossary", v1.TypeGlossary2)
		atlas.GET("/types/businessmetadata", v1.TypeBussinessMetadataGlossary)

		atlas.POST("/types/find", v1.FindTypeDetails)

		atlas.GET("/nodedetails", v1.SearchNodeResult)
		atlas.POST("/types/updatetitleinfo", v1.UpdateTitleInfo)
		atlas.POST("/types/addinfo", v1.AddInfos)
		atlas.POST("/types/updateattribute", v1.UpdateAttrInfo)
		classification := atlas.Group("/classification")
		{

			classification.POST("/create", v1.AddClassificationInfo)
			classification.GET("/allclassification", v1.GetClassificationInfo)
			classification.GET("/classificationattribute", v1.GetClassificationAttributeInfo)
			classification.GET("/classificationname", v1.GetClassificationName)
			classification.POST("/addclassification", v1.EntityAddClassification)
			classification.POST("/updateclassification", v1.UpdateClassificatioInfo)
			classification.POST("/deleteclassification", v1.DeleteClassification)
		}
		glossary := atlas.Group("/glossary")
		{

			glossary.POST("/createglossary", v1.AddGlossaryInfo)
			glossary.POST("/createterm", v1.AddTermInfo)
			glossary.GET("/glossaryinfos", v1.GetGlossaryInfo)
			glossary.GET("/termname", v1.GetTermTotalName)
			glossary.GET("/totalname", v1.GetTermTotalName2)
		}
		label := atlas.Group("/label")
		{
			label.POST("/addlabel", v1.Addlabel)
		}
		userlabel := atlas.Group("/userlabel")
		{
			userlabel.POST("/addlabel", v1.AddUserLabel)

		}
		businessMetadata := atlas.Group("/businessmeta")
		{
			businessMetadata.POST("/createbusinessmeta", v1.CreateBusinessInfo)
			businessMetadata.GET("/getlists", v1.GetBusinessTypeInfos)
			businessMetadata.GET("/bmlists", v1.GetGuidBusiness)
			businessMetadata.GET("/getbusinessinfo", v1.GetBusinessInfos)
			businessMetadata.POST("/addbusinessmeta", v1.AddBusinessInfo)
		}
	}
	comment := r.Group("/api/v1/comment")
	{
		comment.GET("/allcoments", v1.GetAllComments)
		comment.POST("/addcomment", v1.AddComment)
	}
	collect := r.Group("/api/v1/collect")
	{
		collect.POST("/addcollect", v1.AddCollect)
		collect.POST("/deletesinglecollect", v1.DeleteSingleCollect)
		collect.GET("/checkcollected", v1.CheckCollect)
	}
	follow := r.Group("/api/v1/follow")
	{
		follow.POST("/addfollow", v1.AddFollow)
		follow.POST("/deletefollow", v1.DelFollow)
		follow.GET("/checkfollow", v1.CheckFollow)
		follow.GET("/getfollowinfo", v1.GetAllFollow)
	}
	return r
}
