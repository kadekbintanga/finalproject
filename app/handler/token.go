package handler

import(
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "finalproject/app/models"
	"finalproject/app/repository"
)




type TokenHandler struct{
	repo repository.UserRepository
}

func NewTokenHandler() *TokenHandler{
	return &TokenHandler{
		repository.NewUserRepository(),
	}
}

func(h *TokenHandler) CheckToken(email string, token string) bool{
	repo := h.repo
	getUser, err := repo.GetUserByEmail(email)
	if err != nil {
		return false
	}
	getToken, err := repo.GetToken(getUser.ID)
	if err != nil {
		return false
	}
	if getToken.Token != token{
		return false
	}
	return true

}