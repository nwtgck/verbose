package verbose

import (
	"bytes"
	"io"
)

func Encode(ampN int, r io.Reader, w io.Writer) error {
	// One byte slice
	bs := make([]byte, 1)
	for {
		// Read one byte
		n, err := r.Read(bs)
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		// Write amplified bytes
		w.Write(bytes.Repeat(bs, ampN))
	}
	return nil
}
