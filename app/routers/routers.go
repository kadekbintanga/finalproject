package routers

import(
	_ "finalproject/docs"
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/handler"
	"finalproject/app/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



func InitRouter(){
	UserHandler := handler.NewUserHandler()
	PhotoHandler := handler.NewPhotoHandler()
	CommentHandler := handler.NewCommentHandler()
	SocialMediaHandler := handler.NewSocialMediaHandler()
	r := gin.Default()
	api := r.Group("/api/v1")

	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"I am ready!",
		})
	})
	api.GET("/user/health", handler.HealthUser)
	api.POST("/user/register", UserHandler.CreateUser)
	api.POST("/user/login", UserHandler.LoginUser)
	api.PUT("/user", middleware.CheckAuth, UserHandler.UpdateUser)
	api.DELETE("/user", middleware.CheckAuth, UserHandler.DeleteUser)

	api.GET("/photo/health", handler.HealthPhoto)
	api.POST("/photo", middleware.CheckAuth, PhotoHandler.CreatePhoto)
	api.GET("/photo", middleware.CheckAuth, PhotoHandler.GetPhotoUser)
	api.GET("/allphoto", middleware.CheckAuth, PhotoHandler.GetAllPhoto)
	api.PUT("/photo/:photo_id", middleware.CheckAuth, PhotoHandler.UpdatePhoto)
	api.DELETE("/photo/:photo_id", middleware.CheckAuth, PhotoHandler.DeletePhoto)

	api.GET("/comment/health", handler.HealthComment)
	api.POST("/comment", middleware.CheckAuth, CommentHandler.CreateComment)
	api.GET("/comment/:photo_id", middleware.CheckAuth, CommentHandler.GetComment)
	api.PUT("/comment/:comment_id", middleware.CheckAuth, CommentHandler.UpdateComment)
	api.DELETE("/comment/:comment_id", middleware.CheckAuth, CommentHandler.DeleteComment)

	api.GET("/socialmedia/health", handler.HealthSocialMedia)
	api.POST("/socialmedia", middleware.CheckAuth, SocialMediaHandler.CreateSocialMedia)
	api.GET("/socialmedia", middleware.CheckAuth, SocialMediaHandler.GetSocialMedia)
	api.PUT("/socialmedia/:socialmedia_id", middleware.CheckAuth, SocialMediaHandler.UpdateSocialMedia)
	api.DELETE("/socialmedia/:socialmedia_id", middleware.CheckAuth, SocialMediaHandler.DeleteSocialMedia)

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}