.PHONY:server client fmt
export GOPATH:=$(shell pwd)

server:
	go build ./src/arrow/main/server.go; mv server ./bin/

client:
	# build client 
	go build ./src/arrow/main/client.go; mv client ./bin/

test:
	go build ./src/arrow/main/test.go; mv test ./bin/
