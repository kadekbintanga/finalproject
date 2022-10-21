package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/models"
	"finalproject/app/repository"
	"finalproject/app/resource"
	"finalproject/app/helpers"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"strconv"
)

type UserHandler struct{
	repo repository.UserRepository
}

func NewUserHandler() *UserHandler{
	return	&UserHandler{
		repository.NewUserRepository(),
	}
}

// Test Health User godoc
// @Summary Test health user handler
// @Description Test health without any input
// @Tags User
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/user/health [get]
func HealthUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"User Handler is ready!",
	})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags User
// @Param data body resource.InputUser true "body data"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/user/register [post]
func (h *UserHandler) CreateUser(c *gin.Context){
	repo := h.repo

	var req resource.InputUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := helpers.FormatValidationErrorPlayground(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	encrptPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errors := "Someting went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	req.Password = string(encrptPass)
	User := models.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
		Age: req.Age,
	}

	res, err := repo.CreateUser(User)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id":res.ID,
		"username":res.Username,
		"email":res.Email,
		"age":res.Age,
	}

	response := helpers.APIResponse("Success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)

}

// LoginUser godoc
// @Summary Login a user
// @Description Login User  with the input payload
// @Tags User
// @Param data body resource.LoginUser true "body data"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/user/login [post]
func (h *UserHandler) LoginUser(c *gin.Context){
	repo :=h.repo

	var req resource.LoginUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := helpers.FormatValidationErrorPlayground(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	email := req.Email
	dataUser, err := repo.GetUserByEmail(email)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if dataUser.ID == nil{
		errors := "Email not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	checkPass := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(req.Password))
	if checkPass != nil {
		errors := "Password is not correct"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	generateToken, err := helpers.GenerateToken(dataUser.Username, dataUser.Email, dataUser.Age)
	UserToken := models.UserToken{
		UserID: dataUser.ID,
		Token: generateToken,
	}
	_, err = repo.AddToken(UserToken)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"token":generateToken}
	c.JSON(http.StatusOK, response)
}


// UpdateUser godoc
// @Summary Update a user
// @Description Update a user with the input payload
// @Tags User
// @Param user_id query string true "User ID"
// @Param data body resource.UpdateUser true "body data"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/user [put]
func(h *UserHandler) UpdateUser(c *gin.Context){
	repo :=h.repo
	user_id,_ := strconv.ParseUint(c.DefaultQuery("user_id","0"), 10, 64)
	fmt.Println(user_id)
	var req resource.UpdateUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := helpers.FormatValidationErrorBinding(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := helpers.FormatValidationErrorPlayground(err)
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
	dataUser,err := repo.GetUserByEmail(email)
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
	num := uint(user_id) 
	if *dataUser.ID != num{
		errors := "Invalid parameter user_id"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	User := models.User{
		Email: req.Email,
		Username: req.Username,
	}
	update, err := repo.UpdateUser(dataUser.ID, User)
	if err != nil {
		errors := helpers.FormatValidationErrorSQL(err)
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id":update.ID,
		"email":update.Email,
		"username":update.Username,
		"age":dataUser.Age,
		"updated_at":update.UpdatedAt,
	}
	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}
// DeleteUser godoc
// @Summary Delete a user
// @Description delete a user with the token
// @Tags User
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /api/v1/user [delete]
func(h *UserHandler) DeleteUser(c *gin.Context){
	repo := h.repo
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
	userData, err := repo.GetUserByEmail(email)
	if err != nil {
		errors := "User not found"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = repo.DeleteUser(email)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = repo.DeleteToken(userData.ID)
	if err != nil {
		errors := "Something went wrong"
		errorMessage := gin.H{"message":errors}
		response := helpers.APIResponseFailed("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"message":"Your account has been successfuly deleted",
	}
	response := helpers.APIResponse("Success", http.StatusOK,data)
	c.JSON(http.StatusOK, response)
}