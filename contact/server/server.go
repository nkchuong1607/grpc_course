package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nkchuong1607/grpc_course/contact/contactpb"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	connectStr := "root:123456@tcp(127.0.0.1:3306)/contact?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", connectStr, 100, 100)
	if err != nil {
		log.Panicf("register db err %v", err)
	}

	orm.RegisterModel(new(ContactInfo))

	err = orm.RunSyncdb("default", false, false)
	if err != nil {
		log.Panicf("run migrate db err %v", err)
	}

	fmt.Println("connect db sucessfully!")
}

type server struct{}

func (server) Insert(ctx context.Context, req *contactpb.InsertRequest) (*contactpb.InsertResponse, error) {
	log.Printf("calling insert %+v\n", req.Contact)
	ci := ConvertPbContact2ContactInfo(req.Contact)

	err := ci.Insert()

	if err != nil {
		resp := &contactpb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("insert err %v", err),
		}
		return resp, nil
		// return nil, status.Errorf(codes.InvalidArgument, "Insert %+v err %v", ci, err)
	}

	resp := &contactpb.InsertResponse{
		StatusCode: 1,
		Message:    "OK",
	}

	return resp, nil
}

func (server) Read(ctx context.Context, req *contactpb.ReadRequest) (*contactpb.ReadResponse, error) {
	log.Printf("calling read %s\n", req.GetPhoneNumber())
	ci, err := Read(req.GetPhoneNumber())
	if err == orm.ErrNoRows {
		return nil, status.Errorf(codes.InvalidArgument, "Phone %s not exist", req.GetPhoneNumber())
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Read phone %s err %v", req.GetPhoneNumber(), err)
	}

	return &contactpb.ReadResponse{
		Contact: ConvertContactInfo2PbContact(ci),
	}, nil

}

func (server) Update(ctx context.Context, req *contactpb.UpdateRequest) (*contactpb.UpdateResponse, error) {
	if req.GetNewContact() == nil {
		return nil, status.Error(codes.InvalidArgument, "update req with nil contact")
	}
	log.Printf("calling update with data %v\n", req.GetNewContact())
	ci := ConvertPbContact2ContactInfo(req.GetNewContact())
	err := ci.Update()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "update %+v err %v", req.GetNewContact(), err)
	}

	updateContact, err := Read(req.GetNewContact().GetPhoneNumber())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "try to read update contact %+v err %v", req.GetNewContact(), err)
	}

	return &contactpb.UpdateResponse{
		UpdateContact: ConvertContactInfo2PbContact(updateContact),
	}, nil
}

func (server) Delete(ctx context.Context, req *contactpb.DeleteRequest) (*contactpb.DeleteResponse, error) {
	log.Printf("calling delete %s\n", req.GetPhoneNumber())
	if len(req.GetPhoneNumber()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Request delete with empty phone number")
	}

	ci := &ContactInfo{
		PhoneNumber: req.GetPhoneNumber(),
	}
	err := ci.Delete()
	if err != nil {
		return &contactpb.DeleteResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("delete contact %s err %v", req.GetPhoneNumber(), err),
		}, nil
	}

	return &contactpb.DeleteResponse{
		StatusCode: 1,
		Message:    fmt.Sprintf("delete contact %s successfully", req.GetPhoneNumber()),
	}, nil
}

func (server) Search(ctx context.Context, req *contactpb.SearchRequest) (*contactpb.SearchResponse, error) {
	log.Printf("calling search %s\n", req.GetSearchName())
	if len(req.GetSearchName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Request search with empty phone number")
	}

	listCi, err := SearchByName(req.GetSearchName())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Request search %s err %v", req.GetSearchName(), err)
	}

	listPbContact := []*contactpb.Contact{}
	for _, ci := range listCi {
		pbContact := ConvertContactInfo2PbContact(ci)
		listPbContact = append(listPbContact, pbContact)
	}

	return &contactpb.SearchResponse{
		Results: listPbContact,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	contactpb.RegisterContactServiceServer(s, &server{})

	fmt.Println("contact service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
