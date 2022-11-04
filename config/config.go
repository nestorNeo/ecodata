package config

import (
	"encoding/json"
	"io/ioutil"
)

type Server struct {
	Address  string
	Security bool
	// TODO implement overwrite for staging files
	StagingArea       string
	PrefixForTempFile string
	DBAccess          DB
}

type DB struct {
	Enable     bool
	Username   string
	Password   string
	Token      string
	Connection string
}

func DefaultConfig() *Server {
	return &Server{
		Address:           ":8080",
		Security:          false,
		StagingArea:       "",
		PrefixForTempFile: "ecoAudioFile",
		DBAccess: DB{
			Enable: false,
		},
	}
}

func GetConfigFromFile(path string) (*Server, error) {
	cfg := &Server{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, cfg)
	return cfg, err
}
