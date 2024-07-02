run:
	go run ./cmd/main.go ./data/test.zip

gen-test-file:
	@mkdir -p data
	@LC_CTYPE=C cat /dev/urandom | tr -dc '[:alpha:]' | fold -w 1000 | head -n 1 > ./data/test.txt 
	@zip -e ./data/test.zip ./data/test.txt

test:
	go test ./...
