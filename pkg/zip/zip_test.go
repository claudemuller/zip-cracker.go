package zip_test

import (
	z "archive/zip"
	"bytes"
	"testing"

	"github.com/claudemuller/zip-cracker/pkg/zip"
)

func TestIsZip(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{"Is a .zip file", genZipData("test.txt", "test text data", true), true},
		{"Is not a .zip file", genZipData("test.txt", "test text data", false), false},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			// Setup...

			// When...
			got, err := zip.IsZip(bytes.NewReader(tt.data))
			if err != nil {
				t.Fail()
			}

			// Then...
			if tt.want != got {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func genZipData(filename, data string, compress bool) []byte {
	if !compress {
		return []byte(data)
	}

	var buf bytes.Buffer

	zw := z.NewWriter(&buf)
	header := z.FileHeader{
		Name:   filename,
		Method: z.Deflate,
	}

	fw, err := zw.CreateHeader(&header)
	if err != nil {
		panic(err)
	}

	n, err := fw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	if n <= 0 {
		panic("didn't write data")
	}

	if err = zw.Close(); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
