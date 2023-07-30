package model

import (
	"encoding/json"
	"time"
)

type Blog struct {
	Id        int
	Title     string      `json:"title" binding:"required"`
	Author    string      `json:"author" binding:"required"`
	Text      string      `json:"body" binding:"required"`
	Stock     json.Number `json:"stock" binding:"required,number"`
	CreatedAt time.Time   `json:"createdAt"`
}
