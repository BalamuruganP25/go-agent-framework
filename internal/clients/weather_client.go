package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherClient struct {
	client *http.Client
}

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

type WeatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
	} `json:"current"`
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *WeatherClient) GetCoordinates(
	ctx context.Context,
	city string,
) (float64, float64, error) {

	url := fmt.Sprintf(
		"https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1",
		city,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return 0, 0, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var data GeoResponse

	err = json.NewDecoder(resp.Body).
		Decode(&data)
	if err != nil {
		return 0, 0, err
	}

	if len(data.Results) == 0 {
		return 0, 0,
			fmt.Errorf(
				"city not found",
			)
	}

	return data.Results[0].Latitude,
		data.Results[0].Longitude,
		nil
}

func (c *WeatherClient) GetWeather(
	ctx context.Context,
	lat float64,
	lon float64,
) (*WeatherResponse, error) {

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,wind_speed_10m",
		lat,
		lon,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse

	err = json.NewDecoder(resp.Body).
		Decode(&weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
