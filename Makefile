all:
	@go build -o bin/dfy -v

install: all
	@cp bin/dfy ${GOPATH}/bin

test: 
	@go test ./...