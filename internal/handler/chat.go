package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BalamuruganP25/go-agent-framework/internal/agent"
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
)

type ChatHandler struct {
	agent *agent.Service
}

func NewChatHandler(
	agent *agent.Service,
) *ChatHandler {
	return &ChatHandler{
		agent: agent,
	}
}

func (h *ChatHandler) Chat(w http.ResponseWriter, r *http.Request) {
	var req models.ChatRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	resp, err := h.agent.Chat(
		r.Context(),
		req.SessionID,
		req.Message,
	)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		models.ChatResponse{
			SessionID: req.SessionID,
			Response:  resp,
		},
	)
}
