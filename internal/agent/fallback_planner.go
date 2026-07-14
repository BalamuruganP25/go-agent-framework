package agent

import "context"

type FallbackPlanner struct {
	planners []Planner
}

func NewFallbackPlanner(
	planners ...Planner,
) *FallbackPlanner {
	return &FallbackPlanner{
		planners: planners,
	}
}

func (p *FallbackPlanner) Plan(
	ctx context.Context,
	message string,
) (*Plan, error) {

	for _, planner := range p.planners {
		plan, err := planner.Plan(
			ctx,
			message,
		)
		if err != nil {
			continue
		}

		if plan != nil {
			return plan, nil
		}
	}

	return nil, nil
}
