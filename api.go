package qrand

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// URL points to the Quantum Random Number Generator API
const URL = "https://qrng.anu.edu.au/API/jsonI.php"

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Response describes the response from the qrng API
type Response struct {
	DataType string        `json:"type"`
	Length   int           `json:"length"`
	Size     int           `json:"size"`
	Data     []interface{} `json:"data"`
	Success  bool          `json:"success"`
}

// Get makes a formatted GET request with the parameters supplied
func Get(length int) (jsonResponse Response, err error) {
	URLWithParams := fmt.Sprintf("%s?length=%v&type=%v", URL, length, "uint8")

	request, err := http.NewRequest(http.MethodGet, URLWithParams, nil)
	if err != nil {
		return
	}

	response, err := Client.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	return
}
