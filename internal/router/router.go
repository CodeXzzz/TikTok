package router

import (
	"Gin_GoBlog/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter(r *gin.Engine) {
	//use two middlewares:logger and recovery
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	apiRouter := r.Group("/douyin")
	{
		//basic apis
		apiRouter.GET("/feed/", controller.Feed)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", controller.Register)
		apiRouter.POST("/user/login/", controller.Login)
		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.GetFavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)
		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.GetFollowing)
		apiRouter.GET("/relation/follower/list", controller.GetFollowers)
	}
	err := r.Run("localhost:8319")
	if err != nil {
		log.Println("In func InitRouter(): ", err)
	}
}
