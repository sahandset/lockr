package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"gopkg.in/yaml.v3"
)

var Patterns = map[string]*regexp.Regexp{
	"AWS Access Key": regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	"JWT":            regexp.MustCompile(`eyJ[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+`),
	"DB_PASSWORD":    regexp.MustCompile(`(?i)password\s*=\s*.+`),
	"DEBUG":          regexp.MustCompile(`(?i)debug\s*=\s*true`),
}

func LoadCustomRules(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read ruleset: %w", err)
	}

	ext := filepath.Ext(path)
	var raw map[string]string

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &raw); err != nil {
			return fmt.Errorf("invalid JSON format: %w", err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &raw); err != nil {
			return fmt.Errorf("invalid YAML format: %w", err)
		}
	default:
		return fmt.Errorf("unsupported ruleset file type: %s", ext)
	}

	for name, pattern := range raw {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("invalid regex in rule %q: %w", name, err)
		}
		Patterns[name] = re
	}
	return nil
}
