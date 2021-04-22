package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"time"

	"github.com/gorilla/websocket"

	webos "github.com/kaperys/go-webos"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("must include an ip address of the tv and client key")
	}
	ipAddr := os.Args[1]
	clientKey := os.Args[2]

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
	}

	tv, err := webos.NewTV(&dialer, ipAddr)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey(clientKey); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	tv.KeyDown()

	tv.KeyEnter()
}
