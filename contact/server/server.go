package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
