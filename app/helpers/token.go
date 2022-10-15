package helpers

import(
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func GenerateToken(username string, email string, age uint)(string, error){
	payload := jwt.MapClaims{
		"username":username,
		"email":email,
		"age":age,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signed, err := jwtToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	
	return signed, nil
}