package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Email   string   `json:"email"`
	Meta    UserMeta `json:"meta"`
	jwt.RegisteredClaims
}
