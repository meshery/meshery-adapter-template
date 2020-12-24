#!make

.PHONY: go-checks
go-checks: go-lint go-fmt go-mod-tidy

.PHONY: go-vet
go-vet:
	go vet ./...

.PHONY: go-lint
go-lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --config .golangci.yml

.PHONY: go-fmt
go-fmt:
	go fmt ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	./scripts/go-mod-tidy.sh

.PHONY: go-test
go-test:
	./scripts/go-test.sh

#.PHONY: go-test-coverage
#go-test-coverage:
#	./scripts/test-coverage.sh

.PHONY: docker-build
docker-build:
	docker build -t layer5/meshery-<adaptor-name> .

.PHONY: docker-run
docker-run:
	(docker rm -f meshery-<adaptor-name>) || true
	docker run --name meshery-<adaptor-name> -d \
	-p 10007:10007 \
	-e DEBUG=true \
	layer5/meshery-<adaptor-name>

.PHONY: run-local
run-local:
	DEBUG=true go run main.go