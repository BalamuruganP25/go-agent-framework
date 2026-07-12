package api

import (
	"net/http"
	"time"

	"github.com/BalamuruganP25/go-agent-framework/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(chatHandler *handler.ChatHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Compress(5))

	r.Get("/health", health)
	r.Post("/api/v1/chat", chatHandler.Chat)

	return r
}
