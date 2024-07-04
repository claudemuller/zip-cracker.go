package main

import (
	"flag"
	"os"

	z "github.com/yeka/zip"

	"github.com/claudemuller/zip-cracker/pkg/zip"
)

// Improvements: manually do the decompression and password jazz.
func main() {
	var filename string
	var wordlist string

	flag.StringVar(&filename, "file", "", "the .zip file to crack")
	flag.StringVar(&wordlist, "wordlist", "", "the wordlist to use for dictionary attack")
	flag.Parse()

	if filename == "" || wordlist == "" {
		flag.Usage()
		return
	}

	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	isZip, err := zip.IsZip(r)
	if err != nil {
		panic(err)
	}
	r.Close()

	if !isZip {
		println("not a zip file")
		return
	}

	zr, err := z.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	defer zr.Close()

	wlr, err := os.Open(wordlist)
	if err != nil {
		panic(err)
	}
	defer wlr.Close()

	pass, err := zip.Crack(zr, wlr)
	if err != nil {
		panic(err)
	}

	println("password is: ", pass)
}
