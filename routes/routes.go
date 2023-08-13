package routes

import (
	"github.com/gin-gonic/gin"
	"jdyd-douyin/controller/basicApi"
)

func routes() {
	router := gin.Default()

	firstR := router.Group("/douyin")
	{

		firstR.GET("/feed/", basicApi.FeedController)
		firstR.POST("/user/register/", basicApi.UserRegisterController)
		firstR.POST("/user/login/", basicApi.UserLoginController)

		// 验证是否登录
		firstR.GET("/user/", basicApi.UserController)
		firstR.POST("/public/action/", basicApi.PublicActionController)
		firstR.GET("/public/list/", basicApi.PublicListController)

	}

}
