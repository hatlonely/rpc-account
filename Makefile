binary=account
dockeruser=hatlonely
gituser=hatlonely
repository=go-rpc-account
version=1.0.0
export GOPROXY=https://goproxy.cn

.PHONY: build
build: cmd/main.go internal/*/*.go Makefile vendor
	mkdir -p build/bin
	go build -ldflags "-X 'main.Version=`sh scripts/version.sh`'" cmd/main.go && mv main build/bin/${binary} && cp -r config build/

vendor: go.mod go.sum
	@echo "install golang dependency"
	go mod tidy
	go mod vendor

codegen: api/account.proto
	mkdir -p api/gen/go && mkdir -p api/gen/swagger
	protoc -I.. -I. --gofast_out=plugins=grpc,paths=source_relative:api/gen/go/ $<
	protoc -I.. -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:api/gen/go $<
	protoc -I.. -I. --swagger_out=logtostderr=true:api/gen/swagger $<

.PHONY: dockerenv
dockerenv:
	if [ -z "$(shell docker network ls --filter name=testnet -q)" ]; then \
		docker network create -d bridge testnet; \
	fi
	if [ -z "$(shell docker ps -a --filter name=go-build-env -q)" ]; then \
		docker run --name go-build-env --network testnet -d hatlonely/go-build-env:1.0.0 tail -f /dev/null; \
	fi

.PHONY: cleandockerenv
cleandockerenv:
	if [ ! -z "$(shell docker ps -a --filter name=go-build-env -q)" ]; then \
		docker stop go-build-env  && docker rm go-build-env; \
	fi
	if [ ! -z "$(shell docker network ls --filter name=testnet -q)" ]; then \
		docker network rm testnet; \
	fi

.PHONY: image
image: dockerenv
	rm -rf docker
	docker exec go-build-env rm -rf /data/src/${gituser}/${repository}
	docker exec go-build-env mkdir -p /data/src/${gituser}/${repository}
	docker cp . go-build-env:/data/src/${gituser}/${repository}
	docker exec go-build-env bash -c "cd /data/src/${gituser}/${repository} && make build"
	docker cp go-build-env:/data/src/${gituser}/${repository}/build/ docker/
	docker build --tag=hatlonely/${repository}:${version} .
