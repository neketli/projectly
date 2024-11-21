package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Email   string   `json:"email"`
	Meta    UserMeta `json:"meta"`
	jwt.RegisteredClaims
}

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
