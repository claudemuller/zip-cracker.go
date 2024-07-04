# Zip Cracker

[![go](https://github.com/claudemuller/zip-cracker.go/actions/workflows/go.yml/badge.svg)](https://github.com/claudemuller/zip-cracker.go/actions/workflows/go.yml)

![cracker](image.jpg)

A .zip password cracker written in Go.

## Generate Test Data

```bash
make gen-test-file
```

## Run

### Dictionary Attack

```bash
make run ARGS="-file <the_zip_file.zip> -wordlist <the_wordlist.txt>"
```

## Run Tests

```bash
make test
```
