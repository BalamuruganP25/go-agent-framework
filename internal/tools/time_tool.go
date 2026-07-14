package tools

import (
	"context"
	"time"
)

type TimeTool struct{}

func NewTimeTool() *TimeTool {
	return &TimeTool{}
}

func (t *TimeTool) Name() string {
	return "time"
}

func (t *TimeTool) Description() string {
	return "Get the current date and time."
}

func (t *TimeTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {
	return time.Now().
		Format(time.RFC3339), nil
}