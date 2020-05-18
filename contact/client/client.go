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

	// insertContact(client, "09873", "Contact Name 3", "Address 3")
	// readContact(client, "09873")
	// updateContact(client, &contactpb.Contact{
	// 	PhoneNumber: "09873",
	// 	Name:        "Contact Name Update 3",
	// 	Address:     "Address Update 3",
	// })
	// deleteContact(client, "09873")
	searchContact(client, "na")
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

	log.Printf("insert response %+v\n", resp)
}

func readContact(cli contactpb.ContactServiceClient, phone string) {
	req := &contactpb.ReadRequest{
		PhoneNumber: phone,
	}

	resp, err := cli.Read(context.Background(), req)
	if err != nil {
		log.Printf("call read err %v\n", err)
		return
	}

	log.Printf("read response %+v\n", resp.GetContact())
}

func updateContact(cli contactpb.ContactServiceClient, updateContact *contactpb.Contact) {
	req := &contactpb.UpdateRequest{
		NewContact: updateContact,
	}

	resp, err := cli.Update(context.Background(), req)
	if err != nil {
		log.Printf("call update err %v\n", err)
		return
	}

	log.Printf("update response %+v\n", resp.GetUpdateContact())
}

func deleteContact(cli contactpb.ContactServiceClient, phone string) {
	req := &contactpb.DeleteRequest{
		PhoneNumber: phone,
	}
	resp, err := cli.Delete(context.Background(), req)
	if err != nil {
		log.Printf("call delete err %v\n", err)
		return
	}

	log.Printf("delete response %+v\n", resp)
}

func searchContact(cli contactpb.ContactServiceClient, searchString string) {
	req := &contactpb.SearchRequest{
		SearchName: searchString,
	}

	resp, err := cli.Search(context.Background(), req)
	if err != nil {
		log.Printf("call search err %v\n", err)
		return
	}

	log.Printf("search response %+v\n", resp)
}
