package qrand

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInt(t *testing.T) {
	mockResponse := `{
		  "type": "uint16",
		  "length": 8,
		  "data": [
		    14617,
		    26696,
		    12890,
		    5878,
		    18541,
		    38733,
		    1076,
		    34074
		  ],
		  "success": true
		}`

	setupMockClient(mockResponse)
	i := Int()
	assert.Equal(t, i, 14617)
}

func TestRead(t *testing.T) {
	mockResponse := `{
		"type": "uint16",
		"length": 8,
		"data": [
		14617,
		26696,
		12890,
		5878,
		18541,
		38733,
		1076,
		34074
		],
		"success": true
	}`

	setupMockClient(mockResponse)

	key := [8]byte{}
	n, err := Read(key[:])

	assert.Equal(t, n, 8)
	require.NoError(t, err)
}
