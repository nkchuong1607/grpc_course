package main

import (
	"log"

	"github.com/nkchuong1607/grpc_course/contact/contactpb"

	"github.com/astaxie/beego/orm"
)

type ContactInfo struct {
	PhoneNumber string `orm:"size(15);pk"`
	Name        string
	Address     string `orm:"type(text)"`
}

func ConvertPbContact2ContactInfo(pbContact *contactpb.Contact) *ContactInfo {
	return &ContactInfo{
		PhoneNumber: pbContact.PhoneNumber,
		Name:        pbContact.Name,
		Address:     pbContact.Address,
	}
}

func (c *ContactInfo) Insert() error {
	o := orm.NewOrm()

	_, err := o.Insert(c)
	if err != nil {
		log.Printf("Insert contact %+v err %v\n", c, err)
		return err
	}

	log.Printf("Insert %+v successfully\n", c)
	return nil
}
