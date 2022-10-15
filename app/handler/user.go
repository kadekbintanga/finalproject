package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"finalproject/app/models"
	"finalproject/app/repository"
	"finalproject/app/resource"
	"finalproject/app/helpers"
	// "fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct{
	repo repository.UserRepository
}

func NewUserHandler() *UserHandler{
	return	&UserHandler{
		repository.NewUserRepository(),
	}
}

func HealthUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"User Handler is ready!",
	})
}

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

	response := helpers.APIResponse("Success", http.StatusOK,0,0,0, data)
	c.JSON(http.StatusOK, response)

}

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