package qrand

// import (
// 	"strconv"
// )

// Generator!

var (
	qrand Qrand
)

func init() {
	qrand = Qrand{stream: make(chan float64, 8)}
	go qrand.Generator()
}

type Qrand struct {
	stream chan float64
}

func (q *Qrand) Generator() {
	for {
		response, _ := Get(1024, "uint16", 1)
		for _, value := range response.Data {
			i, _ := value.(float64)
			q.stream <- i
		}
	}
	close(q.stream)
}

func (q *Qrand) Int() int {
	i := <-q.stream
	return int(i)
}

func (q *Qrand) Byte() byte {
	b := <-q.stream
	return byte(b)
}

func Int() int {
	i := qrand.Int()
	return int(i)
}

func Read(p []byte) (n int, err error) {

	for n = 0; n < len(p); n++ {
		p[n] = qrand.Byte()
	}

	return
}
