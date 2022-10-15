package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func HealthUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"User Handler is ready!",
	})
}