package service

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type AuthService struct {
	appName             *string
	accessSignedSecret  *string
	refreshSignedSecret *string
	accessExpiredTime   *time.Duration
	refreshExpiredTime  *time.Duration
}

func NewAuthService() *AuthService {
	appName := os.Getenv("APP_NAME")
	accessSignedSecret := os.Getenv("JWT_ACCESS_SECRET")
	accessExpiredTime, _ := time.ParseDuration(os.Getenv("JWT_ACCESS_DURATION"))
	refreshSignedSecret := os.Getenv("JWT_REFRESH_SECRET")
	refreshExpiredTime, _ := time.ParseDuration(os.Getenv("JWT_REFRESH_DURATION"))
	return &AuthService{
		appName:             &appName,
		accessSignedSecret:  &accessSignedSecret,
		accessExpiredTime:   &accessExpiredTime,
		refreshSignedSecret: &refreshSignedSecret,
		refreshExpiredTime:  &refreshExpiredTime,
	}
}

func (s *AuthService) SignAccessJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"exp":     time.Now().Add(*s.accessExpiredTime),
		"iss":     *s.appName,
	})

	tokenString, err := token.SignedString([]byte(*s.accessSignedSecret))
	return &tokenString, err
}

func (s *AuthService) ValidateAccessJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*s.accessSignedSecret), nil
	})
	return token, err
}

func (s *AuthService) SignRefreshJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"exp":     time.Now().Add(*s.refreshExpiredTime),
		"iss":     *s.appName,
	})

	tokenString, err := token.SignedString([]byte(*s.refreshSignedSecret))
	return &tokenString, err
}

func (s *AuthService) ValidateRefreshJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*s.refreshSignedSecret), nil
	})
	return token, err
}
