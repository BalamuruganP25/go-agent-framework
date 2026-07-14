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

	msg := strings.ToLower(message)

	// Time
	if strings.Contains(msg, "time") ||
		strings.Contains(msg, "date") {
		return &Plan{
			ToolName: "time",
		}, nil
	}

	// Calculator
	if strings.Contains(msg, "+") ||
		strings.Contains(msg, "-") ||
		strings.Contains(msg, "*") ||
		strings.Contains(msg, "/") {
		return &Plan{
			ToolName: "calculator",
			Input:    message,
		}, nil
	}

	// Weather
	if strings.Contains(msg, "weather") {
		return &Plan{
			ToolName: "weather",
			Input:    message,
		}, nil
	}

	// MCP - List Directory
	if strings.Contains(msg, "list files") ||
		strings.Contains(msg, "show files") ||
		strings.Contains(msg, "current directory") ||
		strings.Contains(msg, "list directory") {
		return &Plan{
			ToolName: "list_directory",
			Input:    ".",
		}, nil
	}

	// MCP - Read README
	if strings.Contains(msg, "read readme") ||
		strings.Contains(msg, "open readme") {
		return &Plan{
			ToolName: "read_text_file",
			Input:    "README.md",
		}, nil
	}

	// MCP - Search Go files
	if strings.Contains(msg, "go files") ||
		strings.Contains(msg, "*.go") ||
		strings.Contains(msg, "search files") {
		return &Plan{
			ToolName: "search_files",
			Input:    "**/*.go",
		}, nil
	}

	return nil, nil
}
