package middleware

import (
	"clinic-backend/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, role string) (string, error) {
	cfg := config.Load()
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.Load()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func extractToken(c *gin.Context) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, jwt.ErrSignatureInvalid
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, jwt.ErrSignatureInvalid
	}
	return parseToken(parts[1])
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := extractToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			c.Abort()
			return
		}

		role, _ := claims["role"].(string)
		if role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён"})
			c.Abort()
			return
		}

		id, _ := claims["id"].(float64)
		c.Set("userID", uint(id))
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := extractToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			c.Abort()
			return
		}

		role, _ := claims["role"].(string)
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Требуется доступ администратора"})
			c.Abort()
			return
		}

		id, _ := claims["id"].(float64)
		c.Set("adminID", uint(id))
		c.Next()
	}
}
