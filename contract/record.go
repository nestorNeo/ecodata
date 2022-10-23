package contract

import (
	"time"
)

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

type Coordenadas struct {
	Lat float64 `json:"lat"`
	Log float64 `json:"log"`
}

type AudioRecord struct {
	Dispositivo    string        `json:"dispositivo"`
	Desplazamiento []Coordenadas `json:"desplazamiento"`
	Inicio         time.Time     `json:"inicio"`
	Fin            time.Time     `json:"fin"`
	Name           string        // name of the file with extension
	Content        []byte        `json:"content,omitempty"`
}
