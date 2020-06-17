TS=$(shell date -u '+%Y/%m/%d %H:%M:%S')

setup:
	@echo $(TS) Installing dependencies ...
	@sudo apt-get install -y protobuf-compiler
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
	@echo $(TS) Done.
