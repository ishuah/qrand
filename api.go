package qrand

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// URL points to the Quantum Random Number Generator API
const URL = "https://qrng.anu.edu.au/API/jsonI.php"

// Response describes the response from the qrng API
type Response struct {
	DataType string `json:"type"`
	Length   int    `json:"length"`
	Size     int    `json:"size"`
	Data     Data   `json:"data"`
	Success  bool   `json:"success"`
}

type Data []interface{}

// Get makes a formatted GET request with the parameters supplied
func Get(length int, dataType string, size int) (jsonResponse Response, err error) {
	URLWithParams := fmt.Sprintf("%s?length=%v&type=%v&size=%v", URL, length, dataType, size)
	resp, err := http.Get(URLWithParams)

	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)

	return
}
