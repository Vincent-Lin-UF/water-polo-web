package main

import (
	"log"

	"github.com/vincent-lin-uf/water-polo-web/backend/configs"
	"github.com/vincent-lin-uf/water-polo-web/backend/pkg/db"
	"github.com/vincent-lin-uf/water-polo-web/backend/pkg/firebase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()

	pool := db.Connect(cfg.DatabaseURL)
	defer pool.Close()

	firebase.InitAuth(cfg.FirebaseCredFile)

	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/me", firebase.AuthMiddleware(), func(c *gin.Context) {
		uid := c.GetString("uid")
		c.JSON(200, gin.H{"user_id": uid})
	})

	log.Printf("Starting server on port %s\n", cfg.PORT)
	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
