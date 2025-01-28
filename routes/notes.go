package routes

import (
	"markdown-notes-backend/controllers"
	"markdown-notes-backend/middleware"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.Use(middleware.AuthMiddleware())
	{
		notes.POST("/create", controllers.CreateNote) // 创建笔记
		notes.GET("", controllers.GetNotesPaginated)  // 倒序获取笔记
		notes.GET("/search", controllers.SearchNotes) // 搜索笔记
		notes.PUT("/:id", controllers.UpdateNote)     // 更新笔记
		notes.DELETE("/:id", controllers.DeleteNote)  // 删除笔记
		notes.GET("/:id", controllers.GetNoteByID)
	}
}
