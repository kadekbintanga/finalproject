package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/handler"
)



func InitRouter(){
	UserHandler := handler.NewUserHandler()
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
	r.Run()
}