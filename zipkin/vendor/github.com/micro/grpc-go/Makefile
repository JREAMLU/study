all: vet test testrace

deps:
	go get -d -v github.com/micro/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/micro/grpc-go/...

testdeps:
	go get -d -v -t github.com/micro/grpc-go/...

testgaedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/micro/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/micro/grpc-go/...

build: deps
	go build github.com/micro/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/micro/grpc-go/...

vet:
	./vet.sh

test: testdeps
	go test -cpu 1,4 -timeout 5m github.com/micro/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/micro/grpc-go/...

testappengine: testgaedeps
	goapp test -cpu 1,4 -timeout 5m github.com/micro/grpc-go/...

clean:
	go clean -i github.com/micro/grpc-go/...

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	testgaedeps \
	updatetestdeps \
	build \
	proto \
	vet \
	test \
	testrace \
	clean
