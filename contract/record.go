package contract

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
		"CO2":   true,
		"AUDIO": true,
	}
)

type Coordenadas struct {
	Lat float64 `json:"lat"`
	Log float64 `json:"log"`
}

type AudioRecord struct {
	Tipo           string        `json:"tipo"`
	Dispositivo    string        `json:"dispositivo"`
	Desplazamiento []Coordenadas `json:"desplazamiento"`
	// TODO agree in a timeformat to store it in db
	Inicio  string  `json:"inicio"`
	Fin     string  `json:"fin"`
	Name    string  `json:"audiofile"`
	CO2     float64 `json:"co2"`
	Content []byte  `json:"content,omitempty"`
}
