package main

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Mes       string    `gorm:"not null" json:"mes"`             // 任务内容
	Completed bool      `gorm:"default:false" json:"completed"`  // 完成状态
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // 创建时间
}
