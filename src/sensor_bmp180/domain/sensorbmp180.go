package domain

// SensorBMP180 representa los datos del sensor BMP180
type SensorBMP180 struct {
	Status   string  `json:"status"`
	Pressure float64 `json:"pressure"`
}

type SensorBMP180Repository interface {
	Save(data SensorBMP180) error               // Método para guardar datos del sensor
	GetSensorData() ([]SensorBMP180, error)     // Método para obtener datos del sensor
}
