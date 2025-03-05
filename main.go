package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化数据库
	InitDB()

	r := gin.Default()

	// CORS配置
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// 路由配置
	// 添加 /api 前缀
	api := r.Group("/api")
	{
		api.GET("/todos", GetTodos)
		api.POST("/todos", CreateTodo)
		api.PUT("/todos/:id", UpdateTodo)
		api.DELETE("/todos/:id", DeleteTodo)
		api.PATCH("/todos/:id", UpdateTodo) // 更新任务（PATCH）
		// 新增清除已完成任务的路由
		api.DELETE("/todos/clear-completed", ClearCompletedTodos)
	}

	r.Run(":8080")
}
