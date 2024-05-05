package config

import (
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"net/http"
	"os"
)

func LoadDatabaseConfig(path string) (*pg.Options, error) {
	dbConfig, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pgOptions := &pg.Options{}
	err = json.Unmarshal(dbConfig, pgOptions)
	if err != nil {
		return nil, err
	}

	return pgOptions, nil
}

func LoadServerConfig(path string) (*http.Server, error) {
	serverConfig, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	serverOptions := &http.Server{}
	err = json.Unmarshal(serverConfig, serverOptions)
	if err != nil {
		return nil, err
	}

	return serverOptions, nil
}
