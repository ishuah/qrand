package qrand

import (
	"fmt"
	"time"
)

var (
	qrand Qrand
)

func init() {
	qrand = Qrand{stream: make(chan float64, 8)}
	go qrand.Generator()
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
			fmt.Println("The struggle is real")
			continue
		}

		q.err = nil
		go q.Generator()
		return
	}
}

func (q *Qrand) Generator() {
	for {
		response, err := Get(1024, "uint16", 1)

		if err != nil {
			q.err = err
			fmt.Println("calling q.retry")
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

func (q *Qrand) Byte() (byte, error) {
	if q.err != nil {
		return byte(0), q.err
	}

	b := <-q.stream
	return byte(b), nil
}

func Int() (int, error) {
	i, err := qrand.Int()
	return int(i), err
}

func Read(p []byte) (n int, err error) {

	for n = 0; n < len(p); n++ {
		b, err := qrand.Byte()

		if err != nil {
			return 0, err
		}

		p[n] = b
	}

	return
}
