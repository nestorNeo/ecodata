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
	Lat float64
	Log float64
}

type AudioRecord struct {
	Dispositivo    string
	Desplazamiento []Coordenadas
	Inicio         time.Time
	Fin            time.Time
	Name           string // name of the file with extension
	Content        []byte
}
