// pkg/utils/jwt.go
package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTSecret = []byte("your_secret_key") // 从环境变量读取更安全

// 生成JWT Token（有效期7天）
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// 解析JWT Token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
