package agent

import "context"

type Planner interface {
	Plan(ctx context.Context, message string) (*Plan, error)
}
