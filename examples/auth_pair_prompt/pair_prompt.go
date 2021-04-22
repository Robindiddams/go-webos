package main

import (
	"fmt"
	"log"

	webos "github.com/robindiddams/go-webos"
)

func main() {
	tv, err := webos.NewTV("192.168.0.186")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	key, err := tv.AuthorisePrompt()
	if err != nil {
		log.Fatalf("could not authorise using prompt: %v", err)
	}

	// this key can be used for future request to the TV using the AuthoriseClientKey method
	fmt.Println("Client Key:", key)

	tv.Notification("📺👌")
}
