gen-cal:
	protoc calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-server:
	go run server/server.go
run-client:
	go run client/client.go