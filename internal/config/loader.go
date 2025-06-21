package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

// LoadEnvFile loads a .env file into a map
func LoadEnvFile(path string) (map[string]string, error) {
	envMap, err := godotenv.Read(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read .env file: %w", err)
	}
	return envMap, nil
}

// LoadJSONFile parses a JSON file into a flat map
func LoadJSONFile(path string) (map[string]string, error) {
	raw := make(map[string]interface{})
	out := make(map[string]string)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &raw); err != nil {
		return nil, err
	}

	flatten("", raw, out)
	return out, nil
}

// LoadYAMLFile parses a YAML file into a flat map
func LoadYAMLFile(path string) (map[string]string, error) {
	raw := make(map[string]interface{})
	out := make(map[string]string)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(file, &raw); err != nil {
		return nil, err
	}

	flatten("", raw, out)
	return out, nil
}

// flatten turns nested JSON/YAML into flat keys like "auth.token"
func flatten(prefix string, input map[string]interface{}, output map[string]string) {
	for k, v := range input {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case map[string]interface{}:
			flatten(key, val, output)
		case string:
			output[key] = val
		default:
			output[key] = fmt.Sprintf("%v", val)
		}
	}
}
