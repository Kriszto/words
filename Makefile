SHELL := /bin/bash
GO_VERSION ?= 1.17
APP_NAME ?= scarmbled-strings
APP_VERSION ?= 1.0.0
DOCKER_IMAGE_NAME ?= scarmbled-strings

DOCKER_GO_FLAGS ?= --rm --tty \
	--env HOME=/tmp \
	--volume "${PWD}:/app" \
	--volume "${GO_PKG_CACHE}:/go/pkg" \
	--workdir "/app"

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

.PHONY: docker-build
#: Build docker image
docker-build:
	docker build . \
		--file Dockerfile \
		--build-arg APP_NAME=${APP_NAME} \
		--build-arg APP_VERSION=${APP_VERSION} \
		--build-arg GO_VERSION=${GO_VERSION} \
		--build-arg RUNTIME_PACKAGES=${RUNTIME_PACKAGES} \
		-t ${DOCKER_IMAGE_NAME}:${APP_VERSION}

.PHONY: docker-lint
#: Lint with golangci-lint
docker-lint	:
	docker run \
		${DOCKER_GO_FLAGS} \
		golangci/golangci-lint:latest \
		golangci-lint run -v --config .golangci.yml

.PHONY: docker-test
#: Run tests with race detection and coverage report
docker-test:
	touch coverage.out coverage.xml.out
	docker run ${DOCKER_GO_FLAGS} \
		golang:${GO_VERSION} \
		sh -c 'LIST_OF_FILES=$$(go list ./... | grep -v /vendor/ | grep -v /src/ | grep -v /proto/); \
				go test -v -count=1 --race -cover $${LIST_OF_FILES} -coverprofile=coverage.out'

.PHONY: docker-run
#: Run tests with race detection and coverage report
docker-run:
	@echo $(LOG_LEVEL)
	@docker run \
		-v $(realpath $(INPUT)):/input.txt \
		-v $(realpath $(DICT)):/dictionary.txt \
		-t ${DOCKER_IMAGE_NAME}:${APP_VERSION} \
		/usr/local/bin/${APP_NAME} -i /input.txt -d /dictionary.txt \
		$(if $(LOG_LEVEL), -l ${LOG_LEVEL},)
