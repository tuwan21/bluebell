package routes

import (
	"gin_bluebell/controller"
	"gin_bluebell/logger"
	"gin_bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", controller.LoginHandler)
		v1.POST("/signup", controller.SignUpHandler)
	}

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用 JWT 认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/post", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/post2", controller.GetPostListHandler2)
		//投票
		v1.POST("/vote", controller.PostVoteController)
	}
	//r.GET("/community", controller.CommunityHandler)
	//r.GET("/community/:id", controller.CommunityDetailHandler)
	//r.POST("/post", controller.CreatePostHandler)
	//r.GET("/post/:id", controller.GetPostDetailHandler)
	//r.GET("/post", controller.GetPostListHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "404",
		})

	})
	return r
}
