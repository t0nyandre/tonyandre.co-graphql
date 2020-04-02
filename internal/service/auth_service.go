package service

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/lithammer/shortuuid"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type AuthService struct {
	store               *redis.Client
	appName             string
	accessSignedSecret  string
	refreshSignedSecret string
}

func NewAuthService(store *redis.Client) *AuthService {
	appName := os.Getenv("APP_NAME")
	accessSignedSecret := os.Getenv("JWT_ACCESS_SECRET")
	refreshSignedSecret := os.Getenv("JWT_REFRESH_SECRET")
	return &AuthService{
		store:               store,
		appName:             appName,
		accessSignedSecret:  accessSignedSecret,
		refreshSignedSecret: refreshSignedSecret,
	}
}

func (s *AuthService) CreateTokens(user *model.User) (*model.TokenDetails, error) {
	var err error

	td := &model.TokenDetails{
		AccessID:  shortuuid.New(),
		RefreshID: shortuuid.New(),
	}
	td.AtExpires, _ = time.ParseDuration(os.Getenv("JWT_ACCESS_DURATION"))
	td.RtExpires, _ = time.ParseDuration(os.Getenv("JWT_REFRESH_DURATION"))

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"access_id": td.AccessID,
		"exp":       time.Now().Add(td.AtExpires).Unix(),
		"iss":       s.appName,
	})

	td.AccessToken, err = at.SignedString([]byte(s.accessSignedSecret))
	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"refresh_id": td.RefreshID,
		"exp":        time.Now().Add(td.RtExpires).Unix(),
		"iss":        s.appName,
	})

	td.RefreshToken, err = rt.SignedString([]byte(s.refreshSignedSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) CreateAuth(userID string, td *model.TokenDetails) error {
	err := s.store.Set(fmt.Sprintf("%s:%s", "AccessToken", td.AccessID), userID, td.AtExpires).Err()
	if err != nil {
		return err
	}

	err = s.store.Set(fmt.Sprintf("%s:%s", "RefreshToken", td.RefreshID), userID, td.RtExpires).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ValidateAccessToken(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessSignedSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *AuthService) ValidateRefreshToken(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.refreshSignedSecret), nil
	})
	return token, err
}

func (s *AuthService) ValidateAuth(prefix, id string) (*string, error) {
	userid, err := s.store.Get(fmt.Sprintf("%s:%s", prefix, id)).Result()
	if err != nil {
		return nil, err
	}

	return &userid, nil
}

func (s *AuthService) InvalidateAuth(prefix string, id string) (bool, error) {
	err := s.store.Del(fmt.Sprintf("%s:%s", prefix, id)).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}
