package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/BalamuruganP25/go-agent-framework/internal/agent"
	"github.com/BalamuruganP25/go-agent-framework/internal/api"
	"github.com/BalamuruganP25/go-agent-framework/internal/config"
	"github.com/BalamuruganP25/go-agent-framework/internal/db"
	"github.com/BalamuruganP25/go-agent-framework/internal/handler"
	"github.com/BalamuruganP25/go-agent-framework/internal/llm"
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
	"github.com/BalamuruganP25/go-agent-framework/internal/repository"
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

	llmClient := llm.NewOllamaClient(
		cfg.OllamaURL,
		cfg.OllamaModel,
	)

	messageRepo := repository.NewMessageRepository(database)

	agentService := agent.New(
		llmClient,
		messageRepo,
	)

	chatHandler := handler.NewChatHandler(agentService)

	router := api.NewRouter(chatHandler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("server started on :%s", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
