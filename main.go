package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello World")

	var yash = &Person{
		Age:  23,
		Name: "Yash",
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(yash)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newYash := &Person{}
	err = proto.Unmarshal(data, newYash)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newYash` object
	// for good measure
	fmt.Println("Name:", newYash.GetName())
	fmt.Println("Age:", newYash.GetAge())

	fmt.Println("Twitter followers:", newYash.SocialFollowers.GetTwitter())
	fmt.Println("Youtube followers:", newYash.SocialFollowers.GetYoutube())
}
