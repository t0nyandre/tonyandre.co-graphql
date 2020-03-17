package main

import (
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
	"github.com/t0nyandre/go-graphql/internal/resolver"
)

func init() {
	if err := godotenv.Load("internal/config/.env"); err != nil {
		log.Fatalf("Could not find any .env files: %v", err)
	}
}

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(*gql.Merge(" ", "internal/graphql"), &resolver.Resolvers{}, opts...)

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	r.Use(cors.Handler)
	r.Handle("/query", &relay.Handler{Schema: schema})

	fmt.Printf("GraphQL is listening for queries using: %s", os.Getenv("APP_URL"))
	if err := http.ListenAndServe(fmt.Sprintf("%s", os.Getenv("APP_URL")), r); err != nil {
		log.Fatalf("Could not start GraphQL server: %v", err)
	}
}
