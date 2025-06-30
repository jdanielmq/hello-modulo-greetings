package main

import (
	"fmt"
	"log"

	"github.com/jdanielmq/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	//log.SetFlags(0). /** es para quitar la fecha y hora del log
	message, err := greetings.Hello("Juan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Juan", "Daniel", "francisco", "adrian", "ignacio", "javier"}
	messages, err1 := greetings.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(messages)

}
