package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/helpers"
	"finalproject/app/handler"
	"strings"
	"fmt"
)



func CheckAuth(c *gin.Context){
	tokenHeader := c.Request.Header.Get("Authorization")
	tokenArr := strings.Split(tokenHeader, "Bearer ")
	if len(tokenArr) !=2{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"Unauthorized",
		})
		return
	}
	
	tokenStr := tokenArr[1]
	
	payload, err := helpers.ValidateToken(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"Unauthorized",
		})
		return
	}

	TokenHandler := handler.NewTokenHandler()
	email := fmt.Sprint(payload["email"])
	checkToken := TokenHandler.CheckToken(email, tokenStr)
	if checkToken == false{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":"Unauthorized",
		})
		return
	}

	c.Set("username", payload["username"])
	c.Next()
}