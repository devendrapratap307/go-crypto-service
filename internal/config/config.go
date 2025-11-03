// package config

// import (
// 	"os"
// )

// type Config struct {
// 	AESKeyPath string
// 	RSAKeyDir  string
// }

// func Load() *Config {
// 	// simple env-based loader
// 	return &Config{
// 		AESKeyPath: getenv("AES_KEY_PATH", "./keys/aes.key"),
// 		RSAKeyDir:  getenv("RSA_KEY_DIR", "./keys"),
// 	}
// }
// func getenv(k, d string) string {
// 	if v := os.Getenv(k); v != "" {
// 		return v
// 	}
// 	return d
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AESKeyPath string
	RSADir     string
	RestPort   string
	GrpcPort   string
}

func Load() (*Config, error) {
	// Load .env file (optional, skip if running with environment vars)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, falling back to system env")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		AESKeyPath: os.Getenv("AES_KEY_PATH"),
		RSADir:     os.Getenv("RSA_KEY_DIR"),
		RestPort:   os.Getenv("REST_PORT"),
		GrpcPort:   os.Getenv("GRPC_PORT"),
	}, nil
}
