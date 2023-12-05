package main

import (
	"example.com/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	message, err := greetings.Hello("Mario")
	if err != nil {
		log.Fatal("error calling Hello: ", err.Error())
	}
	println(message)
}
