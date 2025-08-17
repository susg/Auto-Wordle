package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	// Test that config is not empty
	assert.NotEmpty(t, config)

	// Test that required fields are present
	assert.NotEmpty(t, config.BaseWordsPath)
	assert.NotNil(t, config.Colours)
	assert.NotEmpty(t, config.MandatoryFile)
	assert.NotNil(t, config.SupportedWordLengths)
	assert.Greater(t, config.FileChunkSize, 0)
	assert.Greater(t, config.WordsBatchSize, 0)
}
