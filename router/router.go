package router

import (
	"final-project-mygram/controllers"
	"final-project-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	routers := gin.Default()

	userRouter := routers.Group("/user")
	{
		userRouter.POST("/register", controllers.RegisterUser)

		userRouter.POST("/login", controllers.LoginUser)
	}

	routers.Static("/img", "./assets")
	photoRouter := routers.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/post", controllers.CreatePhoto)

		photoRouter.GET("/getAll", controllers.GetAllPhoto)

		photoRouter.GET("/getOne/:photoID", controllers.GetOnePhoto)

		photoRouter.PUT("/update/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		photoRouter.DELETE("/delete/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := routers.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())

		commentRouter.POST("/create", controllers.CreateComment)

		commentRouter.GET("/getAll", controllers.GetAllComent)

		commentRouter.GET("/getOne/:commentID", controllers.GetOneComment)

		commentRouter.PUT("/update/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)

		commentRouter.DELETE("/delete/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComent)
	}

	socialMediaRouter := routers.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())

		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)

		socialMediaRouter.GET("/getAll", controllers.GetAllSocialMedia)

		socialMediaRouter.GET("/getOne/:socialMediaID", controllers.GetOneSocialMedia)

		socialMediaRouter.PUT("/update/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

		socialMediaRouter.DELETE("/delete/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return routers
}
