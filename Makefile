run:
	go run cmd/terminote.go internal

test:
	go test -v ./...

test_coverage:
	go test -cover ./...