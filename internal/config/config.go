package config

import (
	"encoding/json"
	"os"

	"github.com/SaisrikarVollala/nebulagate/internal/server"
)

func LoadServers(path string) ([]*server.Server, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var servers []*server.Server

	err = json.Unmarshal(data, &servers)
	if err != nil {
		return nil, err
	}

	for _, s := range servers {
		s.Alive = true
	}

	return servers, nil
}
