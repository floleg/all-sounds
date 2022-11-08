package config

import (
	"fmt"
	"testing"
)

// TestLoadConfig checks existing and non-existing filesystem config files
func TestLoadConfig(t *testing.T) {
	var tests = []struct {
		name   string
		args   map[string]string
		config Config
		err    bool
	}{
		{
			name: "dev",
			args: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "postgres",
				"DB_PASSWORD": "supersecurepass",
				"DB_NAME":     "all-sounds",
			},
			config: Config{"localhost", "5432", "postgres", "supersecurepass", "all-sounds"},
			err:    false,
		},
		{
			name: "docker",
			args: map[string]string{
				"DB_HOST":     "db",
				"DB_PORT":     "5432",
				"DB_USER":     "postgres",
				"DB_PASSWORD": "supersecurepass2",
				"DB_NAME":     "all-sounds-docker",
			},
			config: Config{"db", "5432", "postgres", "supersecurepass2", "all-sounds-docker"},
			err:    false,
		},
		{
			name:   "preprod",
			config: Config{},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			for key, value := range tt.args {
				t.Setenv(key, value)
			}

			config, _ := LoadConfig()

			if config != tt.config {
				t.Errorf("got %v, want %v", config, tt.config)
			}
		})
	}
}
