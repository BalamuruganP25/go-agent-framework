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

	if strings.Contains(
		message,
		"calculate",
	) {
		return &Plan{
			ToolName: "calculator",
			Input: strings.TrimSpace(
				strings.TrimPrefix(
					message,
					"calculate",
				),
			),
		}, nil
	}

	return nil, nil
}
