# gmailer
Go module that sends email using GMail OAuth

Example usage:
```go
package main

import (
  "fmt"
  "github.com/Presence-Web-Services/gmailer"
)

func main() {
  cfg := gmailer.Config{
    ClientID: "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    AccessToken: "YOUR_ACCESS_TOKEN",
    RefreshToken: "YOUR_REFRESH_TOKEN",
    EmailTo: "email_to@email.com",
    EmailFrom: "email_from@email.com",
    ReplyTo: "reply-to@email.com",
    Subject: "Test subject",
    Body: "This is a test email body... thanks!",
  }

  var srv, err = cfg.Authenticate()
  if err != nil {
    fmt.Println("Error occured during authenticate")
  }

  err = cfg.Send(srv)
  if err != nil {
    fmt.Println("Error occured during send")
  }
}
```
