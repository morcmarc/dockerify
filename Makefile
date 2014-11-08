test: 
	@go test ./...

all:
	@go build -o bin/dockerify -v

install: all
	@cp bin/dockerify ${GOPATH}/bin