.golangci.yml.PHONY: build
build:
	go build -o ./scrambled-strings main.go

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run -v --config .golangci.yml

.PHONY: test
test:
	go test  -count=1 -v ./...

.PHONY: coverage
coverage:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out
