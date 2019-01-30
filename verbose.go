package verbose

import (
	"bytes"
	"os"
)

func Encode(ampN int) error {
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
			return err
		}
		// Write amplified bytes
		os.Stdout.Write(bytes.Repeat(bs, ampN))
	}
	return nil
}
