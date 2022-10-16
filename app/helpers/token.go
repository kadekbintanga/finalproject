package helpers

import(
	"github.com/golang-jwt/jwt/v4"
	"os"
	"fmt"
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


func ValidateToken(tokenString string)(map[string]interface{}, error){
	errResp := fmt.Errorf("Unauthorized")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token)(interface{}, error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errResp
		}
		
		return []byte(os.Getenv("JWT_KEY")),nil
	})

	if err != nil {
		return nil, err
	}

	if _,ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid{
		fmt.Println("invalid")
		return nil, errResp
	}

	var payload = map[string]interface{}{}
	claims := token.Claims.(jwt.MapClaims)
	payload["username"] = claims["username"]
	payload["email"] = claims["email"]

	return payload, nil
}