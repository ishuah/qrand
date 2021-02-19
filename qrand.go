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
