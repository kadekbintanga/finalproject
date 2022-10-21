package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/models"
	"finalproject/app/repository"
	"finalproject/app/resource"
	"finalproject/app/helpers"
	"fmt"
	"strings"
	"strconv"
)

type SocialMediaHandler struct{
	repo repository.UserRepository
	repoS repository.SocialMediaRepository
}

func NewSocialMediaHandler() *SocialMediaHandler{
	return	&SocialMediaHandler{
		repository.NewUserRepository(),
		repository.NewSocialMediaRepository(),
	}
}

// Test Health Social Media godoc
// @Summary Test health social media handler
// @Description Test health without any input
// @Tags Social Media
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/socialmedia/health [get]
func HealthSocialMedia(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"SocialMedia Handler is ready!",
	})
}

// Create Social Media godoc
// @Summary Create a social media
// @Description Create a social media with the input payload
// @Tags Social Media
// @Param data body resource.InputSocialMedia true "body data"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/socialmedia [post]
func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context){
	repoUser := h.repo
	repoSocialMedia := h.repoS


	var req resource.InputSocialMedia
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
	SocialMedia := models.SosialMedia{
		Name: req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID: dataUser.ID,
	}

	res, err := repoSocialMedia.CreateSocialMedia(SocialMedia)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"id": res.ID,
		"name": res.Name,
		"social_media_url": res.SocialMediaUrl,
		"user_id": res.UserID,
		"created_at": res.CreatedAt,
	}

	response := helpers.APIResponse("Success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// Get User Social Media godoc
// @Summary Get user social media
// @Description Get social media with bearer token
// @Tags Social Media
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/socialmedia [get]
func (h *SocialMediaHandler) GetSocialMedia(c *gin.Context){
	repoUser := h.repo
	repoSocialMedia := h.repoS

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

	result, err := repoSocialMedia.GetSocialMediabyUserId(dataUser.ID)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var data []map[string]interface{}
	for _, value := range result{
		d := gin.H{
			"id": value.ID,
			"name": value.Name,
			"social_media_url": value.SocialMediaUrl,
			"user_id": value.UserID,
			"created_at": value.CreatedAt,
			"updated_at": value.UpdatedAt,
			"User": gin.H{
				"username": value.User.Username,
				"email": value.User.Email,
			},
		}
		data = append(data, d)
	}

	response := helpers.APIResponse("Success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// Update Social Media godoc
// @Summary Update a social media
// @Description Update a social media with the input payload
// @Tags Social Media
// @Param socialmedia_id path string true "SocialMedia ID"
// @Param data body resource.InputSocialMedia true "body data"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/socialmedia/{socialmedia_id} [put]
func(h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context){
	repoSocialMedia := h.repoS
	repoUser := h.repo
	socialMediaId,_ := strconv.ParseUint(c.Param("socialmedia_id"),10,64)
	var req resource.InputSocialMedia
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
	id := uint(socialMediaId)
	checkSocialMedia, err := repoSocialMedia.GetSocialMediabyId(id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkSocialMedia.ID == nil{
		errors := "Social media  not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if *dataUser.ID != *checkSocialMedia.UserID{
		errors := "Social Media not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	SocialMedia := models.SosialMedia{
		Name: req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}

	update, err := repoSocialMedia.UpdateSocialMedia(id, SocialMedia)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id": checkSocialMedia.ID,
		"name": update.Name,
		"social_media_url": update.SocialMediaUrl,
		"user_id": checkSocialMedia.UserID,
		"updated_at": update.UpdatedAt,
	}

	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}

// Delete Social Media godoc
// @Summary Delete a social media
// @Description delete a social media with the token
// @Tags Social Media
// @Param socialmedia_id path string true "SocialMedia ID"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/socialmedia/{socialmedia_id} [delete]
func(h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context){
	repoSocialMedia := h.repoS
	repoUser := h.repo
	socialMediaId,_ := strconv.ParseUint(c.Param("socialmedia_id"),10,64)
	id := uint(socialMediaId)

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
	checkSocialMedia, err := repoSocialMedia.GetSocialMediabyId(id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkSocialMedia.ID == nil{
		errors := "Social Media not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if *dataUser.ID != *checkSocialMedia.UserID{
		errors := "Social Media not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = repoSocialMedia.DeleteSocialMedia(id)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"message":"Your social media has been successfuly deleted",
	}
	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}