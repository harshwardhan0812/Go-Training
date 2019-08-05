package main

import (
	"errors"
	"fmt"
	"github.com/go-bongo/bongo"
	"time"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName string
	LastName string
	DateOfBirth time.Time
	Address Address
}

type Address struct {
	Address1 string
	Address2 string
	City string
	State string
	Country string
}

func (s *User) BeforeSave(*bongo.Collection) error {
	fmt.Println("I'm called before save", s)
	return errors.New("Error aa gya")
}

func CreateUser(user *User) string {
	//user := &User{
	//	FirstName:"Harshwardhan112",
	//	LastName:"Chauhan",
	//	DateOfBirth:time.Now(),
	//	Address: Address{
	//		Address1: "asasd",
	//		Address2: "dasdad",
	//		City:     "New Delhi",
	//		State:    "Delhi",
	//		Country:  "INDIA",
	//	},
	//}
	//user.DateOfBirth = date(user.DateOfBirth)
	//user.DateOfBirth, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err := Connection.Collection("user").Save(user); err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr)
		}
		fmt.Println("Got a real error:", err)

		return "Error aa gaya"
	}
	return "Ho gaya Add"
}