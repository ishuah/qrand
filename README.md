# qrand
[![Go Report Card](https://goreportcard.com/badge/github.com/ishuah/qrand)](https://goreportcard.com/badge/github.com/ishuah/qrand)
[![GoDoc](https://pkg.go.dev/badge/github.com/ishuah/qrand)](https://pkg.go.dev/github.com/ishuah/qrand)

This package presents a client for the [QRNG@ANU JSON API](https://qrng.anu.edu.au/contact/api-documentation/).
Qrand exposes one function, Read(p []byte) that writes len(p) bytes into p. You can use this function to seed mature random number generator packages in Go.

### Install
`go get github.com/ishuah/qrand`

### Usage

The following example demonstrates how to use qrand.Read(p []byte) to seed math/rand.

```go
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"

	"github.com/ishuah/qrand"
)

func main() {

	key := [8]byte{}
	len, err := qrand.Read(key[:]) // Read 8 random bytes into key
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d quantum random bytes, result: %v\n", len, key)

	seed := binary.BigEndian.Uint64(key[:]) // Convert []byte to uint64
	fmt.Printf("Seed: %d\n", seed)
	rand.Seed(int64(seed)) // Seed math/rand
	fmt.Printf("rand.Int() -> %d\n", rand.Int())
}

```
