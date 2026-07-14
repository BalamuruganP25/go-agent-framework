package agent

import (
	"context"
	"strings"
)

type KeywordPlanner struct{}

func NewKeywordPlanner() *KeywordPlanner {
	return &KeywordPlanner{}
}

func (p *KeywordPlanner) Plan(
	ctx context.Context,
	message string,
) (*Plan, error) {

	message = strings.ToLower(message)

	if strings.Contains(message, "time") {
		return &Plan{
			ToolName: "time",
		}, nil
	}

	return nil, nil
}
