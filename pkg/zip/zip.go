package zip

import (
	"bufio"
	"bytes"
	"errors"
	"io"

	z "github.com/yeka/zip"
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

func Crack(in *z.ReadCloser, wordlist io.Reader) (string, error) {
	scanner := bufio.NewScanner(wordlist)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()

		err := dec(str, in)
		if err != nil {
			continue
		}

		return str, nil
	}

	return "", errors.New("password not found")
}

func dec(pass string, file *z.ReadCloser) error {
	for _, f := range file.File {
		if f.IsEncrypted() {
			f.SetPassword(pass)
		}

		r, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.ReadAll(r); err != nil {
			return err
		}
		defer r.Close()
	}

	return nil
}
