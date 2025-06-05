package firebase

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var client *auth.Client

func InitAuth(credFile string) {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro getting Auth Client: %v\n", err)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			return
		}

		idToken := header[len("Bearer "):]

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}

		c.Set("uid", token.UID)
		c.Next()
	}
}
