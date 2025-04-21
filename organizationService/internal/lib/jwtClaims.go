package lib

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtChangeClaims(secret string, baseClaims map[string]interface{}, overrides map[string]interface{}) (string, error) {
	for k, v := range overrides {
		baseClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":     baseClaims["uid"],
		"role_id": baseClaims["role_id"],
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"email":   baseClaims["email"],
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
