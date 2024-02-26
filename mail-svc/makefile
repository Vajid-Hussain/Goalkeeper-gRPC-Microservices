protoReminder:
	# protoc ./pkg/pb/mail.proto --go_out=./ --go-grpc_out= ./
	protoc ./pkg/pb/reminder.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./

protoMail:
	protoc ./pkg/pb/mail.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./
	# protoc ./pkg/pb/mail.proto --go_out=./ --go-grpc_out= ./

run:
	go run cmd/main.go