package qrand

// Read writes len(p) quantum random bytes into p.
// It returns number of bytes written into p and error(if any occurs).
func Read(p []byte) (n int, err error) {
	response, err := Get(len(p))
	if err != nil {
		return 0, err
	}

	for n = 0; n < len(p); n++ {
		b, _ := response.Data[n].(float64)
		p[n] = byte(b)
	}

	return
}
