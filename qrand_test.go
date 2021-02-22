package qrand

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQrand(t *testing.T) {
	mockResponse := `{
		  "type": "uint16",
		  "length": 16,
		  "data": [
		    19557,
		    1850,
		    3173,
		    37261,
		    23540,
		    52245,
		    26620,
		    1346,
		    17192,
		    30584,
		    58129,
		    8559,
		    57096,
		    31736,
		    60568,
		    7818
		  ],
		  "success": true
		}`
	setupMockClient(mockResponse)

	q := New()
	_, err := q.Int()
	require.NoError(t, err)

	key := [8]byte{}
	n, err := q.Read(key[:])

	assert.Equal(t, n, 8)
	require.NoError(t, err)

	_, err = q.Intn(60000)
	require.NoError(t, err)
}
