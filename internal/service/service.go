package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
	"urlshortener/pkg/short"
)

var (
	store = make(map[string]string)
)

type Service struct {
	router *gin.Engine
	db     *pgxpool.Pool
}

func NewService(router *gin.Engine, db *pgxpool.Pool) *Service {
	return &Service{
		router: router,
		db:     db,
	}
}

func (s *Service) CreateShortUrl(c *gin.Context) {
	var json struct {
		URL         string `json:"url" binding:"required"`
		CustomAlias string `json:"custom_alias"`
	}

	if err := c.ShouldBindJSON(&json); err != nil { // привязываем данные
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url := json.URL
	customAlias := json.CustomAlias
	expireData := time.Now().Add(24 * time.Hour)

	var shortUrl string
	if customAlias != "" {
		shortUrl = customAlias
	} else {
		shortUrl = short.GenerateShortKey(10)
	}
	query := `
		INSERT INTO schema_name.urls (short_id, url, expire_at)
		VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(context.Background(), query, shortUrl, url, expireData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("DB insert error:", err)
		return
	}

	fmt.Printf("Created short URL %s -> %s\n", shortUrl, url)

	c.JSON(http.StatusOK, gin.H{
		"url":          url,
		"short_url":    shortUrl,
		"custom_alias": customAlias,
		"expireData":   expireData.Format(time.RFC3339),
	})
}

func (s *Service) GetShortUrl(c *gin.Context) {
	shortLink := c.Param("short_link")
	query := `
		SELECT url FROM schema_name.urls
		WHERE short_id = $1 AND expire_at > now()
	`

	var originalURL string
	err := s.db.QueryRow(context.Background(), query, shortLink).Scan(&originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	country := c.GetHeader("User-Country")
	deviceType := c.GetHeader("User-Device-Type")

	statInsert := `
		INSERT INTO schema_name.statistic (short_id, url, ip_address, user_agent, country, device_type)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, statErr := s.db.Exec(context.Background(), statInsert, shortLink, originalURL, ip, userAgent, country, deviceType)
	if statErr != nil {
		// Можно залогировать, но не ломать редирект
		fmt.Println("failed to insert statistic:", statErr)
	}

	c.Redirect(http.StatusTemporaryRedirect, originalURL)
}

func (s *Service) GetStats(c *gin.Context) {
	shortLink := c.Param("short_link")
	query := `
		SELECT clicked, ip_address, user_agent, country, device_type
		FROM schema_name.statistic
		WHERE short_id = $1
		ORDER BY clicked DESC
	`
	rows, err := s.db.Query(context.Background(), query, shortLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get statistics"})
		return
	}
	defer rows.Close()

	type Stat struct {
		Clicked    time.Time `json:"clicked"`
		IPAddress  string    `json:"ip_address"`
		UserAgent  string    `json:"user_agent"`
		Country    string    `json:"country"`
		DeviceType string    `json:"device_type"`
	}

	var stats []Stat

	for rows.Next() {
		var stat Stat
		err := rows.Scan(&stat.Clicked, &stat.IPAddress, &stat.UserAgent, &stat.Country, &stat.DeviceType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse statistics"})
			return
		}
		stats = append(stats, stat)
	}

	// 2. Проверяем: были ли вообще переходы
	if len(stats) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No statistics found for this short link"})
		return
	}

	// 3. Возвращаем результат
	c.JSON(http.StatusOK, gin.H{
		"short_id": shortLink,
		"clicks":   len(stats),
		"details":  stats,
	})
}

func (s *Service) RunService() {
	s.router.POST("/api/short", s.CreateShortUrl)
	s.router.GET("/:short_link", s.GetShortUrl)
	s.router.GET("/api/stat/:short_link", s.GetStats)
}
