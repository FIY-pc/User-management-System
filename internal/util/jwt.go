package utils

import (
	"User-management-System/internal/config"
	"errors"
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserId int64 `json:"userId"`
	Exp    int64 `json:"exp"`
}

func (c JwtClaims) Valid() error {
	if jwt.TimeFunc().Unix() > c.Exp {
		return errors.New("token expired")
	}
	return nil
}

func GenerateToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.Jwt.Secret))
}

func ParseToken(tokenString string) (*JwtClaims, error) {
	if len(tokenString) > 7 && tokenString[0:7] == "Bearer" {
		tokenString = tokenString[7:]
	} else {
		return &JwtClaims{}, jwt.NewValidationError("token is not a bearer token", jwt.ValidationErrorMalformed)
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		return &JwtClaims{}, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return &JwtClaims{}, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}
