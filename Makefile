GOPATH := $(go env GOPATH) 

default: test-unit

test-unit:
	go test -v -race -count=1 ./listener/internal/...
	go test -v -race -count=1 ./publisher/internal/...

tools: tool-golangci-lint tool-fumpt tool-moq

tool-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -c bash -s -- -b ${GOPATH}/bin v1.50.1

tool-revive:
	go install github.com/mgechev/revive@master

tool-fumpt:
	go install mvdan.cc/gofumpt

tool-moq:
	go install github.com/matryer/moq

docker-build:
	docker build --no-cache -t listener listener/.

docker-run: docker-build
	docker-compose up -d

docker-logs:
	docker-compose logs -f


