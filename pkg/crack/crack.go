package crack

import (
	"bytes"
	"errors"
	"io"
)

func IsZip(file io.Reader) (bool, error) {
	zipHeader := []byte{0x50, 0x4b, 0x03, 0x04}

	header := make([]byte, 4)
	n, err := file.Read(header)
	if err != nil {
		return false, err
	}
	if n <= 0 {
		return false, errors.New("failed reading header")
	}

	return bytes.Equal(header, zipHeader), nil
}
