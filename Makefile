build:
	go build -o ./bin/flagparse ./cmd/flagparse

testcheck: build
	go vet -vettool=./bin/flagparse ./testdata/a || exit 0
	go vet -vettool=./bin/flagparse ./testdata/b || exit 0
	go vet -vettool=./bin/flagparse ./testdata/c || exit 0
	go vet -vettool=./bin/flagparse ./testdata/d || exit 0
