package agent

import (
	"context"
	"fmt"

	"github.com/BalamuruganP25/go-agent-framework/internal/tools"
)

type Executor struct {
	tools *tools.Registry
}

func NewExecutor(
	tools *tools.Registry,
) *Executor {
	return &Executor{
		tools: tools,
	}
}

func (e *Executor) Execute(
	ctx context.Context,
	plan *Plan,
) (string, error) {

	tool, ok := e.tools.Get(
		plan.ToolName,
	)
	if !ok {
		return "",
			fmt.Errorf(
				"tool %s not found",
				plan.ToolName,
			)
	}

	return tool.Execute(
		ctx,
		plan.Input,
	)
}