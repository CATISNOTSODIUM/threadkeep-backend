package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func loadJWTKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	} 
	return os.Getenv("JWT_SECRET_KEY"), nil
}

func CreateToken(username string) (string, error) {
	secretKeyString, err := loadJWTKey()
	if err != nil {
   		return "", err
    }

	var secretKey = []byte(secretKeyString)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
   		return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
	secretKeyString, err := loadJWTKey()
	if err != nil {
   		return err
    }
	var secretKey = []byte(secretKeyString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	
	if err != nil {
		return err
	}
	
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	
   	return nil
}