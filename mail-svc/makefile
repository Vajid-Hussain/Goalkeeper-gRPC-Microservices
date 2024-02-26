protoReminder:
	# protoc ./pkg/pb/mail.proto --go_out=./ --go-grpc_out= ./
	protoc ./pkg/pb/reminder.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./

protoMail:
	proto ./pkg/pb/mail.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./

run:
	go run cmd/main.go