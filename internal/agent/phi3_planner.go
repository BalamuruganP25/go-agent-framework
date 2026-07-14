package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/BalamuruganP25/go-agent-framework/internal/llm"
	"github.com/BalamuruganP25/go-agent-framework/internal/tools"
)

type PlannerResponse struct {
	ToolName string `json:"tool_name"`
	Input    string `json:"input"`
}

type Phi3Planner struct {
	llm      llm.Client
	registry *tools.Registry
}

func NewPhi3Planner(
	llmClient llm.Client,
	registry *tools.Registry,
) *Phi3Planner {
	return &Phi3Planner{
		llm:      llmClient,
		registry: registry,
	}
}

func (p *Phi3Planner) Plan(
	ctx context.Context,
	message string,
) (*Plan, error) {

	resp, err := p.llm.Chat(
		ctx,
		[]llm.Message{
			{
				Role:    "system",
				Content: p.buildPrompt(),
			},
			{
				Role:    "user",
				Content: message,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	// Phi3 sometimes adds extra text.
	start := strings.Index(resp, "{")
	end := strings.LastIndex(resp, "}")

	if start == -1 || end == -1 || start > end {
		return nil, nil
	}

	jsonStr := resp[start : end+1]

	var plannerResp PlannerResponse

	err = json.Unmarshal(
		[]byte(jsonStr),
		&plannerResp,
	)
	if err != nil {
		return nil, nil
	}

	if plannerResp.ToolName == "" {
		return nil, nil
	}

	log.Printf("Planner response: %s", resp)

	return &Plan{
		ToolName: plannerResp.ToolName,
		Input:    plannerResp.Input,
	}, nil
}

func (p *Phi3Planner) buildPrompt() string {
	var builder strings.Builder

	builder.WriteString(`
		You are a tool planner.

		Choose one tool if it can answer the user's request.

		Return ONLY valid JSON.

		Example:
		{"tool_name":"weather","input":"Chennai"}

		Available tools:

`)

	for _, tool := range p.registry.List() {
		builder.WriteString(
			fmt.Sprintf(
				"- %s\nDescription: %s\n\n",
				tool.Name(),
				tool.Description(),
			),
		)
	}

	return builder.String()
}
