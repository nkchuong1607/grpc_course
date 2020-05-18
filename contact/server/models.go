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
func ConvertContactInfo2PbContact(ci *ContactInfo) *contactpb.Contact {
	return &contactpb.Contact{
		PhoneNumber: ci.PhoneNumber,
		Name:        ci.Name,
		Address:     ci.Address,
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

func Read(phoneNumber string) (*ContactInfo, error) {
	o := orm.NewOrm()
	ci := &ContactInfo{
		PhoneNumber: phoneNumber,
	}
	err := o.Read(ci)
	if err != nil {
		log.Printf("Read contact %+v err %v\n", ci, err)
		return nil, err
	}

	return ci, nil
}

func (c *ContactInfo) Update() error {
	o := orm.NewOrm()

	num, err := o.Update(c, "Name")
	if err != nil {
		log.Printf("Update %+v err %v\n", c, err)
		return err
	}
	log.Printf("update contact %+v, affect %d row\n", c, num)
	return nil
}

func (c *ContactInfo) Delete() error {
	o := orm.NewOrm()

	num, err := o.Delete(c)
	if err != nil {
		log.Printf("delete %+v err %v\n", c, err)
		return err
	}
	log.Printf("delete contact %+v, affect %d row\n", c, num)
	return nil
}

func SearchByName(name string) ([]*ContactInfo, error) {
	result := []*ContactInfo{}
	o := orm.NewOrm()

	num, err := o.QueryTable(new(ContactInfo)).Filter("name__icontains", name).All(&result)

	if err == orm.ErrNoRows {
		log.Printf("search %s found no rows\n", name)
		return result, nil
	}

	if err != nil {
		log.Printf("search %s err %v\n", name, err)
		return nil, err
	}

	log.Printf("search %s found %d rows\n", name, num)
	return result, nil
}
