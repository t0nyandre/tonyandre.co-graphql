package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	gql "github.com/mattdamon108/gqlmerge/lib"
	"github.com/t0nyandre/go-graphql/internal/handler"
	"github.com/t0nyandre/go-graphql/internal/resolver"
	"github.com/t0nyandre/go-graphql/internal/service"
	"github.com/t0nyandre/go-graphql/internal/storage/postgres"
	"github.com/t0nyandre/go-graphql/internal/storage/redis"
)

func init() {
	if err := godotenv.Load("config/env/.env"); err != nil {
		log.Fatalf("Could not find any .env files: %v", err)
	}
}

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(*gql.Merge(" ", "internal/graphql"), &resolver.Resolver{}, opts...)

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	store, err := redis.OpenRedis()
	if err != nil {
		log.Fatalf("Unable to connect to redis: %s", err)
	}

	ctx := context.Background()
	postService := service.NewPostService(db, store)
	userService := service.NewUserService(db, store)

	ctx = context.WithValue(ctx, "postService", postService)
	ctx = context.WithValue(ctx, "userService", userService)

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	r.Use(cors.Handler)
	r.Handle("/query", handler.AddContext(ctx, &relay.Handler{Schema: schema}))

	fmt.Printf("GraphQL is listening for queries using: %s", os.Getenv("APP_URL"))
	if err := http.ListenAndServe(fmt.Sprintf("%s", os.Getenv("APP_URL")), r); err != nil {
		log.Fatalf("Could not start GraphQL server: %v", err)
	}
}
