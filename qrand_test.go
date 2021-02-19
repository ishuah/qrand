package qrand

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInt(t *testing.T) {
	mockResponse := `{
		  "type": "string",
		  "length": 1,
		  "size": 6,
		  "data": [
		    "e984b882190b"
		  ],
		  "success": true
		}`

	setupMockClient(mockResponse)
	_, err := Int()
	require.NoError(t, err)
}

func TestRead(t *testing.T) {
	mockResponse := `{
		  "type": "string",
		  "length": 1,
		  "size": 32,
		  "data": [
		    "218662be8b139f4a38eba4bfe61435913a759bb8af9d3f017574b3ba2f685a34"
		  ],
		  "success": true
	}`
	setupMockClient(mockResponse)

	key := [32]byte{}
	n, err := Read(key[:])

	assert.Equal(t, n, 32)
	require.NoError(t, err)
}
