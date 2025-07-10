package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userID string, email string, role string) (string, string, error) {
	jti := uuid.New().String()

	claims := jwt.MapClaims{
		"sub": userID,
		"email": email,
		"jti": jti,
		"role": role,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SESSION_SECRET")

	signedToken, err := token.SignedString([]byte(secret))
	return signedToken, jti, err
}

func ParseJWT(tokenStr string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET")), nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
