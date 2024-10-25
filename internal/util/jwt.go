package utils

import (
	"User-management-System/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type JwtClaims struct {
	UserId uint   `json:"userId"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

// Valid 方法用于验证 JwtClaims 结构体中的 token 是否有效。
func (c JwtClaims) Valid() error {
	if jwt.TimeFunc().Unix() > c.Exp {
		return jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
	}
	return nil
}

// GenerateToken 用于根据传入的JwtClaim结构体生成token
func GenerateToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.Jwt.Secret))
}

// ParseToken 用于解析token
func ParseToken(tokenString string) (*JwtClaims, error) {
	// 去掉前缀Bearer，验证长度
	if len(tokenString) > 7 && tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		return &JwtClaims{}, jwt.NewValidationError("token is not a bearer token", jwt.ValidationErrorMalformed)
	}
	// 解析token成JwtClaim结构体
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

// JWTAuthMiddleware 是一个JWT中间件，用于进行基本的token有效性验证
func JWTAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从请求中获取Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header missing")
			}
			// 解析并验证JWT令牌
			claims, err := ParseToken(authHeader)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			// 将解析后的claims存入上下文，供后续处理器使用
			c.Set("claims", claims)
			// 调用下一个处理器
			return next(c)
		}
	}
}

// UserRoleMiddleware 是一个中间件，用于检查用户的角色是否为 "user"。
func UserRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("claims").(*JwtClaims)
		if claims.Role != "user" {
			return echo.NewHTTPError(http.StatusForbidden, "Access denied")
		}
		return next(c)
	}
}

// AdminRoleMiddleware 是一个中间件，用于检查用户的角色是否为 "admin"
func AdminRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("claims").(*JwtClaims)
		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "Access denied")
		}
		return next(c)
	}
}
