package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes/models"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func (j *JWT) InitJWT(secretString string) {
	j.secret = secretString
}

func (j *JWT) ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	})
}

func (j *JWT) CreateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiresAt": 15000,
		"ID":        user.ID,
		"username":  user.Username,
	})

	return token.SignedString([]byte(j.secret))
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
