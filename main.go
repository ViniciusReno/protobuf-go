package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"vinicius.reno/protobuf/person"
)

func main() {

	// Create a new person message
	p := &person.Person{
		Name:   "Alice",
		Age:    30,
		Emails: []string{"alice@example.com", "alice@work.com"},
	}

	// Serialize the person message
	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}

	// Deserialize the person message
	p2 := &person.Person{}
	err = proto.Unmarshal(data, p2)
	if err != nil {
		panic(err)
	}

	// Print the person message
	fmt.Println(p2.GetName())
	fmt.Println(p2.GetAge())
	fmt.Println(p2.GetEmails())
}
