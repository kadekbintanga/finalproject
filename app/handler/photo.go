package handler

import(
	"finalproject/app/repository"
	"finalproject/app/resource"
	"finalproject/app/helpers"
	"finalproject/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)


type PhotoHandler struct{
	repoP repository.PhotoRepository
	repo repository.UserRepository
}

func NewPhotoHandler() *PhotoHandler{
	return &PhotoHandler{
		repository.NewPhotoRepository(),
		repository.NewUserRepository(),
	}
}

func HealthPhoto(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Photo Handler is ready!",
	})
}

func (h *PhotoHandler) CreatePhoto(c *gin.Context){
	repoUser := h.repo
	repoPhoto := h.repoP

	var req resource.InputPhoto
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	tokenHeader := c.Request.Header.Get("Authorization")
	tokenArr := strings.Split(tokenHeader, "Bearer ")
	tokenStr := tokenArr[1]
	getEmailToken, err := helpers.ValidateToken(tokenStr)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	email := fmt.Sprint(getEmailToken["email"])
	dataUser,err := repoUser.GetUserByEmail(email)
	if err != nil {
		errors := "Unauthorized"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	Photo := models.Photo{
		Title: req.Title,
		Caption: req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID: dataUser.ID,
	}

	res, err := repoPhoto.CreatePhoto(Photo)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"id":res.ID,
		"title":res.Title,
		"caption":res.Caption,
		"photo_url":res.PhotoUrl,
		"user_id":res.UserID,
		"created_at":res.CreatedAt,
	}
	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}