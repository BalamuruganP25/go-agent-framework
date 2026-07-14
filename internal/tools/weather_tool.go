package tools

import (
	"context"
	"fmt"

	"github.com/BalamuruganP25/go-agent-framework/internal/clients"
)

type WeatherTool struct {
	client *clients.WeatherClient
}

func NewWeatherTool(
	client *clients.WeatherClient,
) *WeatherTool {
	return &WeatherTool{
		client: client,
	}
}

func (t *WeatherTool) Name() string {
	return "weather"
}

func (t *WeatherTool) Description() string {
	return "Get current weather information."
}

func (t *WeatherTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {

	lat, lon, err :=
		t.client.GetCoordinates(
			ctx,
			input,
		)
	if err != nil {
		return "", err
	}

	weather, err :=
		t.client.GetWeather(
			ctx,
			lat,
			lon,
		)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Temperature %.1f°C, Wind %.1f km/h",
		weather.Current.Temperature,
		weather.Current.WindSpeed,
	), nil
}
