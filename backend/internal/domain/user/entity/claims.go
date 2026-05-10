package entity

import "github.com/golang-jwt/jwt/v5"

// JWTClaims represents JWT token claims for a user.
type JWTClaims struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Email   string   `json:"email"`
	Meta    UserMeta `json:"meta"`
	jwt.RegisteredClaims
}

// Tokens contains access and refresh tokens.
type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
