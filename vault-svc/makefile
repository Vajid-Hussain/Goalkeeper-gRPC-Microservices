run:
	go run cmd/main.go

protoc-vault:
	protoc ./pkg/pb/vault.proto --go_out=./ --go-grpc_out=./

protoc-mail:
	protoc ./pkg/pb/mail.proto --go_out=./ --go-grpc_out=./