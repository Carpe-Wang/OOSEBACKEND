package controllers

import (
	"markdown-notes-backend/config"
	"markdown-notes-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	userID := c.GetUint("userID")

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{
		UserID:  userID,
		Title:   input.Title,
		Content: input.Content,
	}

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"note": note})
}

func GetNotes(c *gin.Context) {
	userID := c.GetUint("userID")

	var notes []models.Note
	if err := config.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func UpdateNote(c *gin.Context) {
	userID := c.GetUint("userID")
	noteID := c.Param("id")

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Title = input.Title
	note.Content = input.Content

	if err := config.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"note": note})
}

func DeleteNote(c *gin.Context) {
	userID := c.GetUint("userID")
	noteID := c.Param("id")

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := config.DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func GetNotesPaginated(c *gin.Context) {
	userID := c.GetUint("userID")

	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// 查询用户的笔记，按创建时间倒序排列
	var notes []models.Note
	result := config.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&notes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	// 获取总数
	var total int64
	config.DB.Model(&models.Note{}).Where("user_id = ?", userID).Count(&total)

	// 返回分页结果
	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
		},
	})
}

func SearchNotes(c *gin.Context) {
	userID := c.GetUint("userID")

	// 获取查询参数
	query := c.DefaultQuery("q", "")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// 查询用户的笔记，根据标题或内容搜索
	var notes []models.Note
	result := config.DB.Where("user_id = ?", userID).
		Where("title ILIKE ? OR content ILIKE ?", "%"+query+"%", "%"+query+"%").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&notes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search notes"})
		return
	}

	// 获取总数
	var total int64
	config.DB.Model(&models.Note{}).
		Where("user_id = ?", userID).
		Where("title ILIKE ? OR content ILIKE ?", "%"+query+"%", "%"+query+"%").
		Count(&total)

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
		},
	})
}
