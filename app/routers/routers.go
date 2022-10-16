package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/handler"
	"finalproject/app/middleware"
)



func InitRouter(){
	UserHandler := handler.NewUserHandler()
	PhotoHandler := handler.NewPhotoHandler()
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
	r.Run()
}