// controllers/auth.go
package controllers

import (
	"emo/internal/models"
	"emo/pkg/database"
	"emo/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// 注册请求体
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"` // 用户名必填，长度3-50
	Phone    string `json:"phone" binding:"required,len=11"`          // 手机号必填，11位
	Password string `json:"password" binding:"required,min=6"`        // 密码必填，至少6位
}

// 注册接口
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 检查手机号或用户名是否已存在
	if database.DB.Where("phone = ? OR username = ?", req.Phone, req.Username).First(&models.User{}).Error == nil {
		utils.ErrorResponse(c, http.StatusConflict, "手机号或用户名已存在")
		return
	}

	// 密码哈希加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser := models.User{
		Username:     req.Username,
		Phone:        req.Phone,
		PasswordHash: string(hashedPassword),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "注册失败")
		return
	}

	utils.SuccessResponse(c, "注册成功", nil)
}

// 登录请求体
type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // 支持用户名或手机号登录
	Password   string `json:"password" binding:"required"`
}

// 登录接口
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 查询用户（支持用户名或手机号）
	var user models.User
	query := database.DB.Where("username = ? OR phone = ?", req.Identifier, req.Identifier)
	if err := query.First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户不存在")
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "密码错误")
		return
	}

	// 生成JWT Token
	token, _ := utils.GenerateJWT(user.UserID)
	utils.SuccessResponse(c, "登录成功", gin.H{"token": token, "user_id": user.UserID})
}
