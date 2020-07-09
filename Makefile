GO=go
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))
GOFLAGS= 
BIN=bin/swerve
VERSION := $(shell git rev-parse HEAD)
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
IMAGE="axelspringer/swerve"

all: build/local

help:
	@echo 'Available commands:'
	@echo
	@echo 'Usage:'
	@echo '    make deps     		Install go deps.'
	@echo '    make build/local    	Compile the project.'
	@echo '    make test/local    	Run ginkgo test suites.'
	@echo '    make build/docker    Create docker container'
	@echo '    make clean    		Clean the directory tree.'
	@echo

deps:
	$(GO) get github.com/onsi/ginkgo/ginkgo
	$(GO) get github.com/onsi/gomega/...
	$(GO) get github.com/julienschmidt/httprouter
	$(GO) get github.com/aws/aws-sdk-go
	$(GO) get github.com/sirupsen/logrus

test/local:
	ginkgo --race --cover --coverprofile "$(ROOT_DIR)/swerve.coverprofile" ./...
	go tool cover -html=swerve.coverprofile -o swerve_test_coverage.html

build/local: test/local
	$(GO) build -ldflags "-X main.Version=$(VERSION)" -o $(BIN) $(GOFLAGS) $(RACE) main.go

deploy/local: build/linux 
	docker restart `docker ps | grep "/swerve" | awk '{printf $$1}'` 

build/linux: test/local
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-X main.Version=$(VERSION)" -o "$(BIN)_linux" $(GOFLAGS)  main.go

build/docker:
	docker build -t $(IMAGE):$(TRAVIS_BRANCH)-$(TRAVIS_BUILD_NUMBER) .

compose/up: build/linux
	docker-compose -f docker-compose.yml up -d

compose/down:
	docker-compose -f docker-compose.yml down

push/docker:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push $(IMAGE):$(TRAVIS_BRANCH)-$(TRAVIS_BUILD_NUMBER)

restore:

