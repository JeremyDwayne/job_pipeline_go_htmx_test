package auth

import (
	"job_pipeline/internal/store"

	"github.com/golang-jwt/jwt"
)

type TokenAuth interface {
	GenerateToken(user store.User) (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}
