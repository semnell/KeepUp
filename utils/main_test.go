package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	logger := SetupLogger()
	assert.NotNil(t, logger)
}

func TestSetupSugaredLogger(t *testing.T) {
	sugaredLogger := SetupSugaredLogger()
	assert.NotNil(t, sugaredLogger)
}

func TestLoadConfig(t *testing.T) {
	// Create a temporary YAML file for testing
	tempFile, err := os.CreateTemp("", "test_config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write test YAML content to the temporary file
	testYAML := []byte("jobs:\n  - name: testJob\n    interval: 10")
	err = os.WriteFile(tempFile.Name(), testYAML, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Test LoadConfig function
	conf := LoadConfig(tempFile.Name())
	assert.NotNil(t, conf)
	assert.Equal(t, "testJob", conf.Jobs[0].Name)
	assert.Equal(t, 10, conf.Jobs[0].Interval)
}

func TestRegisterWorkers(t *testing.T) {
	// Create a temporary YAML file for testing
	tempFile, err := os.CreateTemp("", "test_config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write test YAML content to the temporary file
	testYAML := []byte("jobs:\n  - name: testJob\n    interval: 10")
	err = os.WriteFile(tempFile.Name(), testYAML, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Test RegisterWorkers function
	conf := LoadConfig(tempFile.Name())
	RegisterWorkers(conf)
}
