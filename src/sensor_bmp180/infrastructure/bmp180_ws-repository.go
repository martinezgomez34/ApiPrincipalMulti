package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api/src/sensor_bmp180/domain"
)

type BmpClient struct {
    wsAPIURL string
    client   *http.Client
}

func NewBmpClient(wsAPIURL string) *BmpClient {
    return &BmpClient{
        wsAPIURL: wsAPIURL,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (wsc *BmpClient) SendSensorData(data domain.SensorBMP180) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := wsc.client.Post(
        wsc.wsAPIURL+"/api/bmp180-data",
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