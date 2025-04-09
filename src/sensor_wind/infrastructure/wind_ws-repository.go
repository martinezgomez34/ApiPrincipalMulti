package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api/src/sensor_wind/domain"
)

type WindClient struct {
	wsAPIURL string
	client   *http.Client
}

func NewWindClient(wsAPIURL string) *WindClient {
	return &WindClient{
		wsAPIURL: wsAPIURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (wsc *WindClient) SendSensorData(data domain.SensorWind) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := wsc.client.Post(
		wsc.wsAPIURL+"/api/wind-data",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error enviando datos al WS, status: %d", resp.StatusCode)
	}

	return nil
}