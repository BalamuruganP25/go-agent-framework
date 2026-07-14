package mcp

import (
	"context"
	"encoding/json"
	"os/exec"

	sdkmcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

type SDKClient struct {
	session *sdkmcp.ClientSession
}

var _ Client = (*SDKClient)(nil)

func NewFilesystemClient(
	ctx context.Context,
	path string,
) (*SDKClient, error) {

	client := sdkmcp.NewClient(
		&sdkmcp.Implementation{
			Name:    "go-agent-framework",
			Version: "0.1.0",
		},
		nil,
	)

	transport := &sdkmcp.CommandTransport{
		Command: exec.Command(
			"npx",
			"-y",
			"@modelcontextprotocol/server-filesystem",
			path,
		),
	}

	session, err := client.Connect(
		ctx,
		transport,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &SDKClient{
		session: session,
	}, nil
}

func (c *SDKClient) ListTools(
	ctx context.Context,
) ([]Tool, error) {

	result, err := c.session.ListTools(
		ctx,
		nil,
	)
	if err != nil {
		return nil, err
	}

	tools := make([]Tool, 0, len(result.Tools))

	for _, t := range result.Tools {
		tools = append(
			tools,
			Tool{
				Name:        t.Name,
				Description: t.Description,
			},
		)
	}

	return tools, nil
}

func (c *SDKClient) CallTool(
	ctx context.Context,
	name string,
	args map[string]any,
) (string, error) {

	result, err := c.session.CallTool(
		ctx,
		&sdkmcp.CallToolParams{
			Name:      name,
			Arguments: args,
		},
	)
	if err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(
		result.StructuredContent,
		"",
		"  ",
	)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *SDKClient) Close() error {
	return nil
}
