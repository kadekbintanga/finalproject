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

type CommentHandler struct{
	repoC repository.CommentRepository
	repoP repository.PhotoRepository
	repo repository.UserRepository
}

func NewCommentHandler() *CommentHandler{
	return &CommentHandler{
		repository.NewCommentRepository(),
		repository.NewPhotoRepository(),
		repository.NewUserRepository(),
	}
}

func HealthComment(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Comment Handler is ready!",
	})
}

func (h *CommentHandler) CreateComment(c *gin.Context){
	repoUser := h.repo
	repoPhoto := h.repoP
	repoComment := h.repoC

	var req resource.InputComment
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
	num := uint(req.PhotoID)
	fmt.Println(num)
	dataPhoto, err := repoPhoto.GetPhotobyId(num)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if dataPhoto.ID == nil{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var num2 *uint = &num
	Comment := models.Comment{
		Message: req.Message,
		UserID: dataUser.ID,
		PhotoID: num2,
	}

	res, err := repoComment.CreateComment(Comment)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id": res.ID,
		"message": res.Message,
		"photo_id": res.PhotoID,
		"user_id": res.UserID,
		"created_at": res.CreatedAt,
	}

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

func(h *CommentHandler) GetComment(c *gin.Context){
	repoComment := h.repoC
	photoId,_ := strconv.ParseUint(c.Param("photo_id"),10,64)
	photo_id := uint(photoId)

	res, err := repoComment.GetCommentbyPhotoId(photo_id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var data []map[string]interface{}
	for _, value := range res{
		d := gin.H{
			"id": value.ID,
			"message": value.Message,
			"photo_id": value.PhotoID,
			"user_id": value.UserID,
			"updated_at": value.UpdatedAt,
			"created_at": value.CreatedAt,
			"User": gin.H{
				"id": value.User.ID,
				"email": value.User.Email,
				"username": value.User.Username,
			},
			"Photo": gin.H{
				"id": value.Photo.ID,
				"title": value.Photo.Title,
				"caption": value.Photo.Caption,
				"photo_url": value.Photo.PhotoUrl,
				"user_id": value.Photo.UserID,
			},
		}
		data = append(data,d)
	}


	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) UpdateComment(c *gin.Context){
	repoComment:= h.repoC
	comentId,_ := strconv.ParseUint(c.Param("comment_id"),10,64)
	var req resource.UpdateComment
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	id := uint(comentId)
	checkComment, err := repoComment.GetCommentbyId(id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkComment.ID == nil{
		errors := "Photo not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Comment := models.Comment{
		Message: req.Message,
	}

	update, err := repoComment.UpdateComment(id, Comment)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id": id,
		"title": checkComment.Photo.Title,  
		"caption": checkComment.Photo.Caption,
		"message": update.Message,
		"photo_url": checkComment.Photo.PhotoUrl,
		"user_id": checkComment.UserID,
		"update_at": update.UpdatedAt,
	}

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}

func(h *CommentHandler) DeleteComment(c *gin.Context){
	repoComment := h.repoC
	commentId,_ := strconv.ParseUint(c.Param("comment_id"),10,64)
	comment_id := uint(commentId)

	checkComment, err := repoComment.GetCommentbyId(comment_id)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if checkComment.ID == nil{
		errors := "Comment not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = repoComment.DeleteComment(comment_id)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"message":"Your comment has been successfuly deleted",
	}
	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)
}