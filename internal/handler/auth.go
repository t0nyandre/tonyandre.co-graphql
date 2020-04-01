package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/t0nyandre/go-graphql/internal/service"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			isAuthorized = false
			userID       string
		)
		ctx := r.Context()
		token, err := validateBearerAuthHeader(ctx, r)
		if err == nil {
			isAuthorized = true
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userIDByte, _ := base64.StdEncoding.DecodeString(claims["user_id"].(string))
				userID = string(userIDByte[:])
			} else {
				log.Println(err)
			}
		}

		ctx = context.WithValue(ctx, "user_id", &userID)
		ctx = context.WithValue(ctx, "is_authorized", isAuthorized)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateBearerAuthHeader(ctx context.Context, r *http.Request) (*jwt.Token, error) {
	authorization := r.Header.Get("Authorization")

	tokenString := strings.Split(fmt.Sprintf("%s", authorization), " ")
	if len(tokenString) < 2 {
		return nil, fmt.Errorf("Bearer token is not valid")
	}

	token, err := ctx.Value("authService").(*service.AuthService).ValidateAccessJWT(&tokenString[1])
	return token, err
}
