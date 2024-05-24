package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/thapakornd/fiber-go/app/models"
)

type AuthorizationRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func GenerateJWT(u *models.GenerateToken, expHour int64, role string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
	secretString := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = u.IDS
	claims["username"] = u.Username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	t, _ := token.SignedString(secretString)
	return t
}

func VerifyJWT(t string) (jwt.MapClaims, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretString := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(t, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})

	if err != nil {
		fmt.Printf("\n\nError: %s", err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return *claims, nil
	} else {
		fmt.Println(claims, ok, token.Valid)
	}

	return nil, fmt.Errorf("invalid token")
}
