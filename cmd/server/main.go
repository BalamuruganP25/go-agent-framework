package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/BalamuruganP25/go-agent-framework/internal/agent"
	"github.com/BalamuruganP25/go-agent-framework/internal/api"
	"github.com/BalamuruganP25/go-agent-framework/internal/clients"
	"github.com/BalamuruganP25/go-agent-framework/internal/config"
	"github.com/BalamuruganP25/go-agent-framework/internal/db"
	"github.com/BalamuruganP25/go-agent-framework/internal/handler"
	"github.com/BalamuruganP25/go-agent-framework/internal/llm"
	internalmcp "github.com/BalamuruganP25/go-agent-framework/internal/mcp"
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
	"github.com/BalamuruganP25/go-agent-framework/internal/repository"
	"github.com/BalamuruganP25/go-agent-framework/internal/tools"
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

	messageRepo := repository.NewMessageRepository(
		database,
	)

	llmClient := llm.NewOllamaClient(
		cfg.OllamaURL,
		cfg.OllamaModel,
	)

	// ---------------------------------
	// Tool Registry
	// ---------------------------------

	toolRegistry := tools.NewRegistry()

	toolRegistry.Register(
		tools.NewTimeTool(),
	)

	toolRegistry.Register(
		tools.NewCalculatorTool(),
	)

	weatherClient := clients.NewWeatherClient()

	toolRegistry.Register(
		tools.NewWeatherTool(
			weatherClient,
		),
	)

	// ---------------------------------
	// MCP Integration
	// ---------------------------------

	mcpClient, err := internalmcp.NewFilesystemClient(
		context.Background(),
		"/workspace",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer mcpClient.Close()

	mcpTools, err := mcpClient.ListTools(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, tool := range mcpTools {
		log.Printf(
			"Register MCP Tool: %s",
			tool.Name,
		)

		toolRegistry.Register(
			tools.NewMCPTool(
				mcpClient,
				tool,
			),
		)
	}

	// ---------------------------------
	// Agent
	// ---------------------------------

	planner := agent.NewFallbackPlanner(
		agent.NewPhi3Planner(
			llmClient,
			toolRegistry,
		),
		agent.NewKeywordPlanner(),
	)

	executor := agent.NewExecutor(
		toolRegistry,
	)

	agentService := agent.New(
		llmClient,
		messageRepo,
		planner,
		executor,
	)

	chatHandler := handler.NewChatHandler(
		agentService,
	)

	router := api.NewRouter(
		chatHandler,
	)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf(
		"server started on :%s",
		cfg.Port,
	)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
