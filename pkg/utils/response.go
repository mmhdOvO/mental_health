package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一成功响应
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}

// 统一错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
