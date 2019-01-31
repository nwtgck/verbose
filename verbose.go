package verbose

import (
	"bufio"
	"io"
)

func Encode(ampN int, r io.Reader, w io.Writer) error {
	// One byte slice
	bs := make([]byte, 1)
	bw := bufio.NewWriter(w)
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
		b := bs[0]
		for i := 0; i < ampN; i++ {
			bw.WriteByte(b)
		}
	}
	return nil
}

func Decode(ampN int, r io.Reader, w io.Writer) error {
	// One byte slice
	bs := make([]byte, 1)
	br := bufio.NewReader(r)
	for {
		// Read one byte
		n, err := br.Read(bs)
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		// Write top byte
		w.Write(bs)
		// Skip bytes
		br.Discard(ampN - 1)
	}
	return nil
}
