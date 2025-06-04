package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT               string
	DatabaseURL        string
	FirebaseCredFile   string
	AWSRegion          string
	AWSBucket          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found -> reading from environemnt")
	}

	return Config{
		PORT:               getEnv("PORT", "8080"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		FirebaseCredFile:   getEnv("FIREBASE_CREDENTIAL_FILE", ""),
		AWSRegion:          getEnv("AWS_REGION", ""),
		AWSBucket:          getEnv("AWS_Bucket", ""),
		AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
