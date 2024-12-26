package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/BorzooMV/xpensely/services"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "xpensely",
		"sub": username,
		"aud": "xpensely",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateRefreshToken(username string) (string, error) {
	var ctx = context.Background()
	expirationTime := time.Now().Add(time.Hour * 24)
	duration := time.Until(expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "xpensely",
		"sub": username,
		"aud": "xpensely",
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("couldn't create token string: %v", err)
	}

	redisClient := services.ConnectRedis()
	defer redisClient.Close()

	_, err = redisClient.Set(ctx, tokenString, username, duration).Result()
	if err != nil {
		return "", fmt.Errorf("couldn't store token in the redis: %v", err)
	}

	return tokenString, err
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", errors.New("couldn't parse the token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}

	return "", errors.New("invalid token")

}
