package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/config"
)

type ClaimToken struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func CreateToken(user *domain.Users, config *config.Configurations) (string, error) {
	// Membuat token
	token := jwt.New(jwt.SigningMethodHS256)
	var secretKey = []byte(config.JWT_SECRET_KEY)

	// Menambahkan klaim ke token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.BaseModel.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Menandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJwt(tokenString string, config *config.Configurations) (*ClaimToken, error) {
	var secretKey = []byte(config.JWT_SECRET_KEY)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	tokenClaim := ClaimToken{
		UserId:   token.Claims.(jwt.MapClaims)["id"].(string),
		Username: token.Claims.(jwt.MapClaims)["username"].(string),
		Role:     token.Claims.(jwt.MapClaims)["role"].(string),
	}
	return &tokenClaim, nil
}
