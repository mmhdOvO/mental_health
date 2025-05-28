// routes/auth.go
package routes

import (
	"emo/internal/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", controllers.Register) // 注册接口
		authGroup.POST("/login", controllers.Login)       // 登录接口
	}
}
