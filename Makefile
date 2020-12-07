.PHONY: pb run

run:
	go run *.go

pb:
	protoc -I/usr/local/include -I.\
       -I$(GOPATH)/src/ \
       --go_out=plugins=grpc,source_relative=:. welcome.proto
