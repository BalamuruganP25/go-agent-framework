package models

type ChatRequest struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}

type ChatResponse struct {
	SessionID string `json:"session_id"`
	Response  string `json:"response"`
}
