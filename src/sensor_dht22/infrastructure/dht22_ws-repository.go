package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api/src/sensor_dht22/domain"
)

type DhtClient struct {
    wsAPIURL string
    client   *http.Client
}

func NewDhtClient(wsAPIURL string) *DhtClient {
    return &DhtClient{
        wsAPIURL: wsAPIURL,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (wsc *DhtClient) SendSensorData(data domain.SensorDHT22) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := wsc.client.Post(
        wsc.wsAPIURL+"/api/dht22-data",
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