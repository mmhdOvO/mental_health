// middleware/auth.go
package middleware

import (
	"emo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "未提供Token")
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Token无效")
			c.Abort()
			return
		}

		// 将用户ID存入上下文
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
