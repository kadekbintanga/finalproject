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

func HealthSocialMedia(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"SocialMedia Handler is ready!",
	})
}

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

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

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

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

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

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

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
	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}