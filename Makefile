all: deps test install

compile:
	@echo "--> Cross-compiling"
	@goxc

bintray: compile
	@echo "--> Publishing"
	@goxc bintray

install:
	@echo "--> Installing"
	@go install

deps:
	@echo "--> Resolving dependencies"
	@godeps restore

test:
	@echo "--> Running tests"
	@go test -cover ./...