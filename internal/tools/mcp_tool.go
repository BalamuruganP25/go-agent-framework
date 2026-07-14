package tools

import (
	"context"

	internalmcp "github.com/BalamuruganP25/go-agent-framework/internal/mcp"
)

type MCPTool struct {
	client internalmcp.Client
	tool   internalmcp.Tool
}

func NewMCPTool(
	client internalmcp.Client,
	tool internalmcp.Tool,
) *MCPTool {
	return &MCPTool{
		client: client,
		tool:   tool,
	}
}

func (t *MCPTool) Name() string {
	return t.tool.Name
}

func (t *MCPTool) Description() string {
	return t.tool.Description
}

func (t *MCPTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {

	args := map[string]any{}

	if input != "" {
		args["path"] = input
	}

	return t.client.CallTool(
		ctx,
		t.tool.Name,
		args,
	)
}
