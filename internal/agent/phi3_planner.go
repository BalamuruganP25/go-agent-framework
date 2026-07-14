package agent

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/BalamuruganP25/go-agent-framework/internal/llm"
)

const plannerPrompt = `
You are an AI tool planner.

Available tools:

1. time
Description:
Get the current date and time.

Examples:
User: What time is it?
Response:
{"tool_name":"time","input":""}

2. calculator
Description:
Perform arithmetic calculations.

Examples:
User: calculate 100 * 25
Response:
{"tool_name":"calculator","input":"100 * 25"}

User: what is 100 multiplied by 25
Response:
{"tool_name":"calculator","input":"100 * 25"}

If no tool is needed, return:

{"tool_name":"","input":""}

IMPORTANT:
- Return ONLY valid JSON.
- Do not explain.
- Do not use markdown.
- Do not wrap JSON in code blocks.
`

type PlannerResponse struct {
	ToolName string `json:"tool_name"`
	Input    string `json:"input"`
}

type Phi3Planner struct {
	llm llm.Client
}

func NewPhi3Planner(
	llmClient llm.Client,
) *Phi3Planner {
	return &Phi3Planner{
		llm: llmClient,
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
				Content: plannerPrompt,
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

	return &Plan{
		ToolName: plannerResp.ToolName,
		Input:    plannerResp.Input,
	}, nil
}