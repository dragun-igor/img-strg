generate:
	@cd api && rm -rf *
	@protoc *.proto  -I . -I .. --go_out=./api --go-grpc_out=./api