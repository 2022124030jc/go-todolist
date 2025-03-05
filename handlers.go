package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取所有Todo
func GetTodos(c *gin.Context) {
	var todos []Todo
	if result := DB.Find(&todos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// 创建Todo
func CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if result := DB.Create(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// 更新Todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo

	// 查找任务
	if result := DB.First(&todo, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 绑定更新数据
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 更新任务字段
	if mes, ok := updateData["mes"]; ok {
		todo.Mes = mes.(string)
	}
	if completed, ok := updateData["completed"]; ok {
		todo.Completed = completed.(bool)
	}

	// 保存更新
	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// 删除Todo（软删除）
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo

	if result := DB.First(&todo, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

// ClearCompletedTodos 删除所有已完成的任务
func ClearCompletedTodos(c *gin.Context) {
	// 查找所有已完成的任务
	var todos []Todo
	if result := DB.Where("completed = ?", true).Find(&todos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch completed todos"})
		return
	}

	// 如果没有找到已完成的任务，直接返回成功消息
	if len(todos) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No completed todos to clear"})
		return
	}

	// 删除所有已完成的任务
	if err := DB.Delete(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete completed todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Completed todos cleared successfully"})
}
