gen-contact:
	protoc contactpb/contact.proto --go_out=plugins=grpc:.
run-server:
	go run server/server.go server/models.go
run-client:
	go run client/client.go