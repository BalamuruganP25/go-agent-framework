package agent

import (
	"context"

	"github.com/BalamuruganP25/go-agent-framework/internal/llm"
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
	"github.com/BalamuruganP25/go-agent-framework/internal/repository"
)

type Service struct {
	llm      llm.Client
	messages repository.MessageRepository
}

func New(
	llmClient llm.Client,
	messageRepo repository.MessageRepository,
) *Service {
	return &Service{
		llm:      llmClient,
		messages: messageRepo,
	}
}

func (s *Service) Chat(
	ctx context.Context,
	sessionID string,
	message string,
) (string, error) {

	history, err := s.messages.
		GetBySession(ctx, sessionID)
	if err != nil {
		return "", err
	}

	var msgs []llm.Message

	for _, m := range history {
		msgs = append(
			msgs,
			llm.Message{
				Role:    m.Role,
				Content: m.Content,
			},
		)
	}

	msgs = append(
		msgs,
		llm.Message{
			Role:    "user",
			Content: message,
		},
	)

	resp, err := s.llm.Chat(ctx, msgs)
	if err != nil {
		return "", err
	}

	err = s.messages.Save(
		ctx,
		&models.Message{
			SessionID: sessionID,
			Role:      "user",
			Content:   message,
		},
	)
	if err != nil {
		return "", err
	}

	err = s.messages.Save(
		ctx,
		&models.Message{
			SessionID: sessionID,
			Role:      "assistant",
			Content:   resp,
		},
	)
	if err != nil {
		return "", err
	}

	return resp, nil
}
