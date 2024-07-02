package main

import (
	"flag"
	"os"

	"github.com/claudemuller/zip-cracker/pkg/crack"
)

func main() {
	var filename string

	flag.StringVar(&filename, "file", "", "the .zip file to crack")
	flag.Parse()

	if filename == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	isZip, err := crack.IsZip(file)
	if err != nil {
		panic(err)
	}

	if !isZip {
		println("not a zip file")
		return
	}

	println("zip file")
}
