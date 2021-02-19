package qrand

import (
	"strconv"
)

func Int() (int, error) {
	response, err := Get(1, "hex16", 6)

	if err != nil {
		return 0, err
	}

	h, _ := response.Data[0].(string)

	i, err := strconv.ParseInt(h, 16, 64)

	if err != nil {
		return 0, err
	}

	return int(i), nil
}

func Read(p []byte) (n int, err error) {
	size := len(p)
	length := 1
	if len(p) > 1024 {
		size = 1024
		length = len(p) / 1024
	}

	response, err := Get(length, "hex16", size)

	if err != nil {
		return
	}

	var data string

	for _, value := range response.Data {
		str, _ := value.(string)
		data += str
	}

	for n = 0; n < len(p); n++ {
		val, _ := strconv.ParseInt(data[n*2:(n*2)+2], 16, 64)
		p[n] = byte(val)
	}

	return
}
