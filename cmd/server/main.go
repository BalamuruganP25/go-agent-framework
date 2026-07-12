package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/BalamuruganP25/go-agent-framework/internal/api"
	"github.com/BalamuruganP25/go-agent-framework/internal/config"
	"github.com/BalamuruganP25/go-agent-framework/internal/db"
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	database, err := db.New(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(
		&models.Message{},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatal(err)
	}

	router := api.NewRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("server started on :%s", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
