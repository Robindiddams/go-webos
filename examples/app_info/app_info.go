package main

import (
	"fmt"
	"log"
	"os"

	webos "github.com/robindiddams/go-webos"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("must include an ip address of the tv and client key")
	}
	ipAddr := os.Args[1]
	clientKey := os.Args[2]

	tv, err := webos.NewTV(ipAddr)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey(clientKey); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	fmt.Println(tv.CurrentApp())
}
