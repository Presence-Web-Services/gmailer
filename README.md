# gmailer
Go module that sends email using GMail OAuth.

Email envelope header shows email sent from name as the `ReplyTo` address, while the from address is `EmailFrom`:

`From: reply-to@email.com <email-from@email.com>`

This is useful in combination with `EmailFrom` set to a no-reply address.

Based on code borrowed from [BinodKafle/gomail](https://github.com/BinodKafle/gomail)

Example usage (main.go):
```go
package main

import (
	"fmt"
	"github.com/presence-web-services/gmailer/v2"
)

func main() {
	cfg := gmailer.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		AccessToken:  "YOUR_ACCESS_TOKEN",
		RefreshToken: "YOUR_REFRESH_TOKEN",
		EmailTo:      "email-to@email.com",
		EmailFrom:    "email-from@email.com",
		ReplyTo:      "reply-to@email.com",
		Subject:      "Test Subject",
		Body:         "This is a test email body... thanks!",
		HP:           false,
	}

	var err = cfg.Authenticate()
	if err != nil {
		fmt.Println("Error occurred during authenticate")
	}

	err = cfg.Send()
	if err != nil {
		fmt.Println("Error occurred during send")
	}
}
```

How to run:
```
go get github.com/presence-web-services/gmailer/v2
go run main.go
```
