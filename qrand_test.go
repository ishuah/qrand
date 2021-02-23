package qrand

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQrand(t *testing.T) {
	mockResponse := `{
	  "type": "uint16",
	  "length": 32,
	  "data": [
	    4411,
	    3276,
	    28247,
	    35756,
	    65068,
	    41619,
	    10689,
	    62740,
	    56618,
	    55394,
	    597,
	    50127,
	    25730,
	    9534,
	    24830,
	    52791,
	    12857,
	    26368,
	    23174,
	    53896,
	    44735,
	    23506,
	    36228,
	    48110,
	    220,
	    31654,
	    17868,
	    29436,
	    44061,
	    43284,
	    53124,
	    45306
	  ],
	  "success": true
	}`
	setupMockClient(mockResponse, 200, nil)

	// Test Read
	key := [8]byte{}
	n, err := Read(key[:])
	assert.Equal(t, n, 8)
	require.NoError(t, err)

	mockResponse = `{
		"success": false
	}`
	setupMockClient(mockResponse, 404, errors.New("Unsuccessful call"))

	key = [8]byte{}
	_, err = Read(key[:])
	require.Error(t, err)

	teardownMockClient()
}

// Benchmarks

func BenchmarkRead(b *testing.B) {
	p := [16]byte{}
	for n := 0; n < b.N; n++ {
		Read(p[:])
	}
}
