package config

import (
	"fmt"
	"testing"
)

// TestLoadConfig checks existing and non-existing filesystem config files
func TestLoadConfig(t *testing.T) {
	var tests = []struct {
		name   string
		config Config
		panic  bool
	}{
		{
			name:   "dev",
			config: Config{"localhost", "5432", "postgres", "-NQI2tIM?|G>B@A2", "all-sounds"},
			panic:  false,
		},
		{
			name:   "docker",
			config: Config{"db", "5432", "postgres", "-NQI2tIM?|G>B@A2", "all-sounds"},
			panic:  false,
		},
		{
			name:   "preprod",
			config: Config{},
			panic:  true,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			// Here we expect LoadConfig to panic because the config file should not exist
			// USe the https://go.dev/doc/effective_go#recover trick to regain routine control
			if tt.panic {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in f", r)
					}
				}()

				LoadConfig(tt.name, "../../configs")
				t.Errorf("Panic was expected")
			} else {
				// Check that the loaded confiuration is conform as exepected if the file exists on filesytem
				config, _ := LoadConfig(tt.name, "../../configs")

				if config != tt.config {
					t.Errorf("got %v, want %v", config, tt.config)
				}
			}
		})
	}
}
