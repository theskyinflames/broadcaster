GOPATH := $(go env GOPATH) 

default: test-unit

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
	docker build --no-cache -t publisher publisher/.

docker-hub-push:
	docker push jaumearus/broadcaster-listener:latest
	docker push jaumearus/broadcaster-publisher:latest

docker-run: docker-build
	docker-compose up -d

docker-logs:
	docker-compose logs -f

k8s-deploy:
	devops/k8s/deploy.sh

k8s-undeploy:
	devops/k8s/undeploy.sh

test-unit:
	cd ./listener;go test -v -race -count=1 ./internal/...
	cd ./publisher;go test -v -race -count=1 ./internal/...


