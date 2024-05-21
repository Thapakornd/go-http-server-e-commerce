package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/thapakornd/fiber-go/app/models"
)

func GenerateJWT(u *models.User) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
	secretString := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.IDS
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(secretString)
	return t
}
