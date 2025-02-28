default: run

proto:
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

build: proto	
	@go build -o bin/server server/main.go	
	@go build -o bin/client client/main.go

clean:
	@rm -rf bin
	@rm -f proto/*.pb.go