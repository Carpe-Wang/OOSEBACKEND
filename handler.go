package main

import (
	"log"
	"markdown-notes-backend/config"
	"markdown-notes-backend/models"
	"markdown-notes-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局路由
var router *gin.Engine

// 初始化函数
func init() {
	// 连接数据库
	config.ConnectDatabase()

	// 初始化 OAuth
	config.InitOAuth()

	// 自动迁移数据库模型
	if err := config.DB.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")

	// 初始化 Gin 路由
	router = gin.New()
	router.RedirectTrailingSlash = false

	// 根路径
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running on Vercel!")
	})

	// 注册路由
	routes.AuthRoutes(router)
	routes.NotesRoutes(router)

	for _, route := range router.Routes() {
		log.Printf("Registered Route: %s %s\n", route.Method, route.Path)
	}
}

// Vercel 的入口函数
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
