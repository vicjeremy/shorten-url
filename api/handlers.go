package api

// The application business logic
import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migrate the schema
	if err := DB.AutoMigrate(&Url{}, &Count{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}

func generateShortURL() string {
	b := make([]byte, 6) // 6 bytes = 48 bits
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func CreateUrl(c *gin.Context) {
	var url Url

	// Bind the request body
	if err := c.ShouldBindJSON(&url); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	if err := DB.Where("url = ?", url.Url).First(&url).Error; err == nil {
		ResponseJSON(c, http.StatusConflict, "Url already exists", url)
		return
	}
	url.ShortCode = generateShortURL()
	DB.Create(&url)

	count := Count{Url: url, AccessCount: 0}
	DB.Create(&count)

	ResponseJSON(c, http.StatusCreated, "Url created successfully", url)
}

func GetUrl(c *gin.Context) {
	var url Url
	if err := DB.Where("short_code = ?", c.Param("code")).First(&url).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Url not found", nil)
		return
	}

	// Find or create Count record
	var urlCount Count
	result := DB.Where("id = ?", url.ID).First(&urlCount)
	if result.Error != nil {
		// Create new Count record if not found
		urlCount = Count{Url: url, AccessCount: 1}
	} else {
		urlCount.AccessCount++
	}
	DB.Save(&urlCount)

	ResponseJSON(c, http.StatusOK, "Url retrieved successfully", url)
}

func UpdateUrl(c *gin.Context) {
	var url Url
	if err := DB.Where("short_code = ?", c.Param("code")).First(&url).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Url not found", nil)
		return
	}

	// bind the request body
	if err := c.ShouldBindJSON(&url); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	DB.Save(&url)
	count := Count{Url: url}
	DB.Save(&count)
	ResponseJSON(c, http.StatusOK, "Url updated successfully", url)
}

func DeleteUrl(c *gin.Context) {
	var url Url
	if err := DB.Where("short_code = ?", c.Param("code")).First(&url).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Url not found", nil)
		return
	}
	// Delete associated Count record
	DB.Where("id = ?", url.ID).Delete(&Count{})

	// Delete the URL record
	DB.Delete(&url)

	ResponseJSON(c, http.StatusNoContent, "Url deleted successfully", nil)
}

func StatsUrl(c *gin.Context) {
	var url Url
	if err := DB.Where("short_code = ?", c.Param("code")).First(&url).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Url not found", nil)
		return
	}

	// Find or create Count record
	var urlCount Count
	result := DB.Where("id = ?", url.ID).First(&urlCount)
	if result.Error != nil {
		// Create new Count record if not found
		urlCount = Count{Url: url, AccessCount: 0}
	}

	ResponseJSON(c, http.StatusOK, "Url stats retrieved successfully", urlCount)
}
