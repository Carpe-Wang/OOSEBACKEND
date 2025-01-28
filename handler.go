package handler

import (
	"log"
	"markdown-notes-backend/config"
	"markdown-notes-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// 初始化数据库
	config.ConnectDatabase()

	// 初始化路由
	router = gin.New()

	// 测试根路径
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running on Vercel!")
	})

	// 注册其他路由
	routes.AuthRoutes(router)
	routes.NotesRoutes(router)

	log.Println("Routes initialized successfully")
}

// Vercel 的入口函数
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
