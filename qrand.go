package qrand

import (
	"time"
)

func New() *Qrand {
	qrand := Qrand{stream: make(chan float64, 8)}
	go qrand.generator()
	return &qrand
}

type Qrand struct {
	stream chan float64
	err    error
}

func (q *Qrand) retry() {
	for {
		time.Sleep(time.Second)
		_, err := Get(1, "uint16", 1)
		if err != nil {
			continue
		}

		q.err = nil
		go q.generator()
		return
	}
}

func (q *Qrand) generator() {
	for {
		response, err := Get(1024, "uint16", 1)
		if err != nil {
			q.err = err
			go q.retry()
			return
		}

		for _, value := range response.Data {
			i, _ := value.(float64)
			q.stream <- i
		}
	}
	defer close(q.stream)
}

func (q *Qrand) Int() (int, error) {
	if q.err != nil {
		return 0, q.err
	}

	i := <-q.stream
	return int(i), nil
}

func (q *Qrand) Intn(n int) (int, error) {
	if n <= 0 {
		panic("Invalid argument to Intn")
	}
	if q.err != nil {
		return 0, q.err
	}

	var i float64

	if n&(n-1) == 0 {
		i = <-q.stream
		return int(i) & (n - 1), nil
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	i = <-q.stream
	for int32(i) > max {
		i = <-q.stream
	}

	return int(i) % n, nil
}

func (q *Qrand) Perm(n int) ([]int, error) {
	if n <= 0 {
		panic("Invalid argument to Perm")
	}

	m := make([]int, n)
	for j := 0; j < n; {
		i, err := q.Intn(n)
		if err != nil {
			return nil, err
		}
		m[j] = int(i)
		j++
	}
	return m, nil
}

func (q *Qrand) Byte() (byte, error) {
	if q.err != nil {
		return byte(0), q.err
	}

	b := <-q.stream
	return byte(b), nil
}

func (q *Qrand) Read(p []byte) (n int, err error) {

	for n = 0; n < len(p); n++ {
		b, err := q.Byte()

		if err != nil {
			return 0, err
		}

		p[n] = b
	}

	return
}
