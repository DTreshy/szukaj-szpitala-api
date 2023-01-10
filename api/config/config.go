package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DTreshy/go-validate/validate"
)

type Config struct {
	MongoUri string
	GrpcBind string
	Username string
	Password string
}

func NewConfig() (*Config, error) {
	cfg := Config{
		MongoUri: "mongodb://localhost:27017",
		GrpcBind: "localhost:55010",
		Username: "admin",
		Password: "admin",
	}

	jsonFile, err := os.Open("./config.json")
	if err != nil {
		return &cfg, nil
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("reading config file failed: %w", err)
	}

	configFromFile := make(map[string]any, 0)

	if err := json.Unmarshal(byteValue, &configFromFile); err != nil {
		return nil, fmt.Errorf("unmarshalling config file failed: %w", err)
	}

	mongoUri, err := getEndpointFromConfig(configFromFile, "MongoUri")
	if err != nil {
		return nil, err
	}

	if mongoUri != "" {
		cfg.MongoUri = mongoUri
	}

	grpcBind, err := getEndpointFromConfig(configFromFile, "GrpcBind")
	if err != nil {
		return nil, err
	}

	if grpcBind != "" {
		cfg.GrpcBind = grpcBind
	}

	return &cfg, nil
}

func getEndpointFromConfig(m map[string]any, name string) (string, error) {
	val, ok := m[name]
	if !ok {
		return "", nil
	}

	stringval, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("%s is not a string", name)
	}

	if err := validate.Endpoint(stringval); err != nil {
		return "", err
	}

	return stringval, nil
}
