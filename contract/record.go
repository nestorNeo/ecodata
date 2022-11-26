package contract

import "errors"

/*
{
    "dispositivo": "nestor-samsung",
    "desplazamiento": [
      {
        "lat": 20.675171,
        "log": -103.347328
      },
      {
        "lat": 20.669280,
        "log": -103.388014
      },
      {
        "lat": 20.668312,
        "log": -103.465009
      }
    ],
    "inicio": "2022-10-18 18:54:49.000",
    "fin": "2022-10-18 18:54:49.000"
  }
*/

var (
	VALIDTYPES = map[string]bool{
		"CO2_MEASURE":   true,
		"AUDIO":         true,
		"NOISE_MEASURE": true,
	}
)

type Record struct {
	Tipo        string    `json:"type"`
	Dispositivo string    `json:"dispositivo"`
	Inicio      string    `json:"begin_time"`
	Fin         string    `json:"end_time"`
	Ppm         float64   `json:"ppm,omitempty"`
	Moda        []float64 `json:"moda,omitempty"`
	Decibels    []float64 `json:"dB,omitempty"`
	Std         []float64 `json:"std,omitempty"`
	Interval    int       `json:"interval,omitempty"`
	Lat         []float64 `json:"lat,omitempty"`
	Lon         []float64 `json:"lon,omitempty"`
	Content     []byte    `json:"content,omitempty"`
}

func (r *Record) IsValidCO2() error {

	if r.Ppm < 0 {
		return errors.New("ppm cannot be smaller than 0")
	}
	return nil
}

func (r *Record) IsValidNoise() error {
	valid := len(r.Moda) > 0 && len(r.Std) > 0 && len(r.Decibels) > 0
	if valid {
		return nil
	}
	return errors.New("moda, std and dB are required")
}
