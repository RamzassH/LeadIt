package jwt

import (
	"github.com/RamzassH/LeadIt/authService/backend/internal/config"
	"github.com/RamzassH/LeadIt/authService/backend/internal/domain/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func NewToken(user models.User, duration time.Duration) (string, error) {
	cfg := config.MustLoadConfig()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(cfg.TokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
