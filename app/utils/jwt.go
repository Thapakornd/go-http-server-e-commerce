package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/thapakornd/fiber-go/app/models"
)

type AuthorizationRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func GenerateJWT(u *models.GenerateToken, expHour int64) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
	secretString := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	if strconv.FormatInt(u.IDS, 10)[0] == '1' {
		claims["role"] = os.Getenv("USER_ROLE")
	} else {
		claims["role"] = os.Getenv("ADMIN_ROLE")
	}

	claims["id"] = u.IDS
	claims["username"] = u.Username
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
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
