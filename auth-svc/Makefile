protoAuth:
	protoc ./pkg/pb/*.proto --go_out=./ --go-grpc_out=./
protoMail:
	protoc ./pkg/pb/*.proto --go_out=./ --go-grpc_out=./

run: 
	go run cmd/main.go
