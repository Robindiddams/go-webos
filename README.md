# Go WebOS 📺

A small Go library for interaction with webOS enabled TVs. ~Tested on LG webOS TV UH668V (webOS version 05.30.20).~ This is forked from github.com/kaperys/go-webos.

[![Go Report Card](https://goreportcard.com/badge/github.com/kaperys/go-webos)](https://goreportcard.com/report/github.com/kaperys/go-webos)

```go
dialer := websocket.Dialer{
    HandshakeTimeout: 10 * time.Second,
    // the TV uses a self-signed certificate
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    NetDial: (&net.Dialer{Timeout: time.Second * 5}).Dial,
}

tv, err := webos.NewTV(&dialer, "<tv-ipv4-address>")
if err != nil {
    log.Fatalf("could not dial TV: %v", err)
}
defer tv.Close()

// the MessageHandler must be started to read responses from the TV
go tv.MessageHandler()

// AuthorisePrompt shows the authorisation prompt on the TV screen
key, err := tv.AuthorisePrompt()
if err != nil {
    log.Fatalf("could not authorise using prompt: %v", err)
}

// the key returned can be used for future request to the TV using the
// AuthoriseClientKey(<key>) method, instead of AuthorisePrompt()
fmt.Println("Client Key:", key)

// see commands.go for available methods
tv.Notification("📺👌")
```

## V1.0.0 Goals:

- [x] make the examples runnable
- [ ] add `Discover` method to search local network for tv (if feasable without extraneous dependencies)
- [ ] internalize the Dialer and MessageHandler
- [ ] use json.RawMessage for partial decoding of messages instead of github.com/mitchellh/mapstructure
- [ ] maybe switch to easyjson
- [ ] use standard go errors

See [examples](examples/) for usage.
