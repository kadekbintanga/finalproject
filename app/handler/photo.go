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
	"strconv"
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

// Test Health Photo godoc
// @Summary Test health photo handler
// @Description Test health without any input
// @Tags Photo
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/photo/health [get]
func HealthPhoto(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Photo Handler is ready!",
	})
}

// Create Photo godoc
// @Summary Create a photo
// @Description Create a photo with the input payload
// @Tags Photo
// @Param data body resource.InputPhoto true "body data"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/photo [post]
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
	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}

// Get User Photo godoc
// @Summary Get user photo
// @Description Get photo with bearer token
// @Tags Photo
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/photo [get]
func (h *PhotoHandler) GetPhotoUser(c *gin.Context){
	repoUser := h.repo
	repoPhoto := h.repoP

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

	result, err := repoPhoto.GetPhotobyUserId(dataUser.ID)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var data []map[string]interface{}
	for _, value := range result{
		d := gin.H{
			"id": value.ID,
			"title": value.Title,
			"caption": value.Caption,
			"photo_url": value.PhotoUrl,
			"user_id": value.UserID,
			"created_at": value.CreatedAt,
			"updated_at": value.UpdatedAt,
			"User": gin.H{
				"email": value.User.Email,
				"username": value.User.Username,
			},
		}
		data = append(data, d)
	}


	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}

// Get All  Photo godoc
// @Summary Get all  photo handler
// @Description Get all photo  without any input
// @Tags Photo
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/allphoto [get]
func (h *PhotoHandler) GetAllPhoto(c *gin.Context){
	repoPhoto := h.repoP
	result, err := repoPhoto.GetAllPhoto()
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var data []map[string]interface{}
	for _, value := range result{
		d := gin.H{
			"id": value.ID,
			"title": value.Title,
			"caption": value.Caption,
			"photo_url": value.PhotoUrl,
			"user_id": value.UserID,
			"created_at": value.CreatedAt,
			"updated_at": value.UpdatedAt,
			"User": gin.H{
				"email": value.User.Email,
				"username": value.User.Username,
			},
		}
		data = append(data, d)
	}

	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}

// UpdatePhoto godoc
// @Summary Update a photo
// @Description Update a photo with the input payload
// @Tags Photo
// @Param photo_id path string true "Photo ID"
// @Param data body resource.InputPhoto true "body data"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/photo/{photo_id} [put]
func (h *PhotoHandler) UpdatePhoto(c *gin.Context){
	repoUser := h.repo
	repoPhoto := h.repoP
	photoId,_ := strconv.ParseUint(c.Param("photo_id"),10,64)
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
	if dataUser.ID == nil{
		errors := "Unauthorized"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	id := uint(photoId)
	checkPhoto, err := repoPhoto.GetPhotobyId(id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkPhoto.ID == nil{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if *dataUser.ID != *checkPhoto.UserID{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Photo := models.Photo{
		Title: req.Title,
		Caption: req.Caption,
		PhotoUrl: req.PhotoUrl,
	}
	update, err := repoPhoto.UpdatePhoto(id, Photo)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id": id,
		"title": update.Title,
		"caption": update.Caption,
		"photo_url": update.PhotoUrl,
		"user_id": checkPhoto.UserID,
		"updated_at": update.UpdatedAt,
	}

	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}

// DeletePhoto godoc
// @Summary Delete a photo
// @Description delete a photo with the token
// @Tags Photo
// @Param photo_id path string true "Photo ID"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/photo/{photo_id} [delete]
func(h *PhotoHandler) DeletePhoto(c *gin.Context){
	repoUser := h.repo
	repoPhoto := h.repoP
	photoId,_ := strconv.ParseUint(c.Param("photo_id"),10,64)

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
	if dataUser.ID == nil{
		errors := "Unauthorized"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := uint(photoId)
	checkPhoto, err := repoPhoto.GetPhotobyId(id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkPhoto.ID == nil{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if *dataUser.ID != *checkPhoto.UserID{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	err = repoPhoto.DeletePhoto(id)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"message":"Your photo has been successfuly deleted",
	}
	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}