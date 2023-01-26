GOPATH := $(go env GOPATH) 

default: test-unit

test-unit:
	go test -v -race -count=1 ./internal/...

lint:
	golangci-lint run
	revive -config ./revive.toml
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum

tools: tool-golangci-lint tool-fumpt tool-moq

tool-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -c bash -s -- -b ${GOPATH}/bin v1.50.1

tool-revive:
	go install github.com/mgechev/revive@master

tool-fumpt:
	go install mvdan.cc/gofumpt

tool-moq:
	go install github.com/matryer/moq

