package qrand

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func setupMockClient(mockResponse string) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(mockResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
}

func TestAPICallSuccess(t *testing.T) {
	mockResponse := `{
			  "type": "string",
			  "length": 1,
			  "size": 10,
			  "data": [
			    "0195a96a618e47f02bbf"
			  ],
			  "success": true
			}`

	setupMockClient(mockResponse)
	response, err := Get(1, "hex16", 10)

	require.NoError(t, err)
	assert.Equal(t, response.Success, true)
}
