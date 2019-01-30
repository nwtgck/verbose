package main

import (
	"bytes"
	"os"
)

func main () {
	// TODO: hard code
	ampN := 10
	// One byte slice
	bs := make([]byte, 1)
	for {
		// Read one byte
		n, err := os.Stdin.Read(bs)
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		// Write amplified bytes
		os.Stdout.Write(bytes.Repeat(bs, ampN))
	}
}
