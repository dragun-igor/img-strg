generate:
	@cd messenger && rm -rf *
	@protoc *.proto  -I . -I .. --go_out=./api --go-grpc_out=./api