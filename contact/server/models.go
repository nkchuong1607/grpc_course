package main

type ContactInfo struct {
	Id          int64  `orm:"auto"`
	PhoneNumber string `orm:"size(15)"`
	Name        string
	Address     string `orm:"type(text)"`
}
