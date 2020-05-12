package main

import (
	"context"
	"log"

	"github.com/nkchuong1607/grpc_course/contact/contactpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50070", grpc.WithInsecure())

	if err != nil {
		log.Fatalf(" err while dial %v", err)
	}
	defer cc.Close()

	client := contactpb.NewContactServiceClient(cc)

	insertContact(client, "09871", "Contact Name 2", "Address 2")
}

func insertContact(cli contactpb.ContactServiceClient, phone, name, addr string) {
	req := &contactpb.InsertRequest{
		Contact: &contactpb.Contact{
			PhoneNumber: phone,
			Name:        name,
			Address:     addr,
		},
	}
	resp, err := cli.Insert(context.Background(), req)

	if err != nil {
		log.Printf("call insert err %v\n", err)
		return
	}

	log.Printf("insert response %+v", resp)
}
