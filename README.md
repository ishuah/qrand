# qrand
[![Go Report Card](https://goreportcard.com/badge/github.com/ishuah/qrand)](https://goreportcard.com/badge/github.com/ishuah/qrand)
[![GoDoc](https://pkg.go.dev/badge/github.com/ishuah/qrand)](https://pkg.go.dev/github.com/ishuah/qrand)

This package presents a client for the [QRNG@ANU JSON API](https://qrng.anu.edu.au/contact/api-documentation/).

### Install
`go get github.com/ishuah/qrand`

### Usage
```go
package main

import (
	"fmt"
	"github.com/ishuah/qrand"
	"log"
)

func main() {

	q := qrand.New()

	i, err := q.Int()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Quantum random int:", i)

	b, err := q.Byte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Quantum random byte:", b)

	n := 1000
	i, err = q.Intn(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Quantum random int in [0,%d): %d\n", n, i)

	n = 10
	p, err := q.Perm(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Quantum random slice with elements in [0,%d): %v\n", n, p)

	key := [16]byte{}
	len, _ := q.Read(key[:])
	fmt.Printf("Wrote %d quantum random bytes, result: %v\n", len, key)

}

```

### Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/ishuah/qrand
BenchmarkRead-8      	      84	  19316313 ns/op
BenchmarkIntn10-8    	    1268	   1280759 ns/op
BenchmarkIntn100-8   	    1270	   1276469 ns/op
BenchmarkPerm10-8    	     121	  13413403 ns/op
BenchmarkPerm100-8   	      13	 124809965 ns/op
PASS
ok  	github.com/ishuah/qrand	65.249s
```
