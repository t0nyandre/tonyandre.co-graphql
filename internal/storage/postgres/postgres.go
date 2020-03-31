package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/t0nyandre/go-graphql/internal/model"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_SSL")))
	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Printf("Retry database connection in 5 seconds ...")
		time.Sleep(time.Duration(5) * time.Second)
		return ConnectDB()
	}

	db.AutoMigrate(&model.User{}, &model.Post{})

	log.Printf("Database successfully connected")
	return db, nil
}
