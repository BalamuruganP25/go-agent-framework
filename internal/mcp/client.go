package mcp

import "context"

type Tool struct {
	Name        string
	Description string
}

type Client interface {
	ListTools(
		ctx context.Context,
	) ([]Tool, error)

	CallTool(
		ctx context.Context,
		name string,
		args map[string]any,
	) (string, error)

	Close() error
}