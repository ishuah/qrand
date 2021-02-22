package qrand

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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
	setupMockClient(mockResponse)

	// New Qrand instance
	q := New()

	// Test Int
	_, err := q.Int()
	require.NoError(t, err)

	// Test Read
	key := [8]byte{}
	n, err := q.Read(key[:])
	assert.Equal(t, n, 8)
	require.NoError(t, err)

	// Test Intn
	i, err := q.Intn(60000)
	assert.LessOrEqual(t, i, 60000)
	require.NoError(t, err)
	assert.Panics(t, func() { q.Intn(0) }, "Intn did not panic")

	// Test Perm
	p, err := q.Perm(5)
	require.NoError(t, err)
	assert.Equal(t, len(p), 5)
	assert.Panics(t, func() { q.Perm(0) }, "Perm did not panic")

	teardownMockClient()
}

// Benchmarks
func benchmarkIntn(input int, b *testing.B) {
	q := New()
	for n := 0; n < b.N; n++ {
		q.Intn(input)
	}
}

func benchmarkPerm(input int, b *testing.B) {
	q := New()
	for n := 0; n < b.N; n++ {
		q.Perm(input)
	}
}

func BenchmarkIntn10(b *testing.B)  { benchmarkIntn(10, b) }
func BenchmarkIntn100(b *testing.B) { benchmarkIntn(100, b) }

func BenchmarkPerm10(b *testing.B)  { benchmarkPerm(10, b) }
func BenchmarkPerm100(b *testing.B) { benchmarkPerm(100, b) }
