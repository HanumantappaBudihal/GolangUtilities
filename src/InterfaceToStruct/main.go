package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type NewCustomEvent struct {
	Name  string
	Phone string
	Email string
}

type INewCustomEvent interface{}

func main() {
	newCustomer := NewCustomEvent{Name: "XYZ", Phone: "12345678", Email: "test@gmail.com"}
	convert(newCustomer)
}

func convert(event INewCustomEvent) {
	c := NewCustomEvent{}
	mapstructure.Decode(event, &c)
	fmt.Printf("Event is : %v", c)
}
