package tools

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type CalculatorTool struct{}

func NewCalculatorTool() *CalculatorTool {
	return &CalculatorTool{}
}

func (t *CalculatorTool) Name() string {
	return "calculator"
}

func (t *CalculatorTool) Description() string {
	return "Perform basic arithmetic calculations."
}

func (t *CalculatorTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {

	fields := strings.Fields(input)

	if len(fields) != 3 {
		return "",
			fmt.Errorf(
				"invalid expression",
			)
	}

	left, err := strconv.ParseFloat(
		fields[0],
		64,
	)
	if err != nil {
		return "", err
	}

	operator := fields[1]

	right, err := strconv.ParseFloat(
		fields[2],
		64,
	)
	if err != nil {
		return "", err
	}

	var result float64

	switch operator {
	case "+":
		result = left + right

	case "-":
		result = left - right

	case "*":
		result = left * right

	case "/":
		if right == 0 {
			return "",
				fmt.Errorf(
					"division by zero",
				)
		}

		result = left / right

	default:
		return "",
			fmt.Errorf(
				"unsupported operator",
			)
	}

	return fmt.Sprintf(
		"%v",
		result,
	), nil
}
