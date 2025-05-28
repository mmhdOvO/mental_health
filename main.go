// main.go
package main

import (
	"emo/config"
	"emo/internal/models"
	"emo/pkg/database"
	"emo/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 加载配置文件（假设配置文件路径为 ./config/config.yaml）
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 初始化数据库（使用配置中的 DSN）
	if err := database.InitDB(cfg.Database.DSN); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 执行数据库迁移（AutoMigrate）
	if err := database.DB.AutoMigrate(
		&models.EmotionDiary{},
		&models.Test{},
		&models.AnonymousComment{},
		&models.AnonymousPost{},
		&models.PsychologicalAnalysis{},
		&models.AdviceTemplate{},
		&models.EmotionCategory{},
		&models.User{},
	); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册路由（保持原有逻辑）
	routes.AuthRoutes(r)

	// 启动服务（使用配置中的端口）
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
