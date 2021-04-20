package verbose

import (
	"io"
	"io/ioutil"
)

func Encode(ampN int, r io.Reader, w io.Writer) error {
	var bs [1]byte
	for {
		// Read one byte
		n, err := r.Read(bs[:])
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		// Write amplified bytes
		for i := 0; i < ampN; i++ {
			if _, err := w.Write(bs[:]); err != nil {
				return err
			}
		}
	}
	return nil
}

func Decode(ampN int, r io.Reader, w io.Writer) error {
	var bs [1]byte
	for {
		// Read one byte
		n, err := r.Read(bs[:])
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		// Write top byte
		if _, err := w.Write(bs[:]); err != nil {
			return err
		}
		// Skip bytes
		if _, err := io.CopyN(ioutil.Discard, r, int64(ampN-1)); err != nil {
			return err
		}
	}
	return nil
}
