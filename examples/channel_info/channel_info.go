package main

import (
	"fmt"
	"log"

	webos "github.com/robindiddams/go-webos"
)

func main() {
	tv, err := webos.NewTV("192.168.1.67")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey("6c7b2ec679ffd1c2736abd621153eabb"); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	fmt.Println(tv.CurrentChannel())
	fmt.Println(tv.CurrentProgram())
}
