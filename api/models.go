package api

// The data model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Url struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode" gorm:"uniqueIndex"`
	CreateAt  time.Time `json:"createAt" gorm:"autoCreateTime"`
	UpdateAt  time.Time `json:"updateAt" gorm:"autoUpdateTime"`
}

type Count struct {
	Url
	AccessCount int `json:"accessCount"`
}
type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, response)
}
