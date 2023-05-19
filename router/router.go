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
		user.GET("/useraudits", v1.GetUserAudit)
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

		classification := atlas.Group("/classification")
		{

			classification.POST("/create", v1.AddClassificationInfo)
			classification.GET("/allclassification", v1.GetClassificationInfo)
			classification.GET("/classificationattribute", v1.GetClassificationAttributeInfo)
			classification.GET("/classificationname", v1.GetClassificationName)
			classification.POST("/addclassification", v1.EntityAddClassification)
		}
		glossary := atlas.Group("/glossary")
		{

			glossary.POST("/createglossary", v1.AddGlossaryInfo)
			glossary.POST("/createterm", v1.AddTermInfo)
			glossary.GET("/glossaryinfos", v1.GetGlossaryInfo)
			glossary.GET("/termname", v1.GetTermTotalName)
		}
		label := atlas.Group("/label")
		{
			label.POST("/addlabel", v1.Addlabel)
		}
		userlabel := atlas.Group("/userlabel")
		{
			userlabel.POST("/addlabel", v1.AddUserLabel)

		}
	}
	comment := r.Group("/api/v1/comment")
	{
		comment.GET("/allcoments", v1.GetAllComments)
		comment.POST("/addcomment", v1.AddComment)
	}
	return r
}
