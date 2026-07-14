package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	client := mcp.NewClient(
		&mcp.Implementation{
			Name:    "go-agent-framework",
			Version: "0.1.0",
		},
		nil,
	)

	transport := &mcp.CommandTransport{
		Command: exec.Command(
			"npx",
			"-y",
			"@modelcontextprotocol/server-filesystem",
			".",
		),
	}

	session, err := client.Connect(
		context.Background(),
		transport,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected")

	result, err := session.CallTool(
		context.Background(),
		&mcp.CallToolParams{
			Name: "list_directory",
			Arguments: map[string]any{
				"path": ".",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"Result: %+v",
		result,
	)
}
