/*
Gmailer allows for OAUTH authentication to GMail and sending of basic email messages
*/
package gmailer

import (
	"context"
	"encoding/base64"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Config is type containing OAuth credentials and email message information
type Config struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
	EmailTo      string
	EmailFrom    string
	ReplyTo      string
	Subject      string
	Body         string
  Service      *gmail.Service
}

// Authenticate authenticates with GMail OAuth, and returns a pointer to gmail.Service when successfully authenticated
func (config *Config) Authenticate() (err error) {
	oauthConfig := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}
	oauthToken := oauth2.Token{
		AccessToken:  config.AccessToken,
		RefreshToken: config.RefreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	tokenSource := oauthConfig.TokenSource(context.Background(), &oauthToken)

	config.Service, err = gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
  return
}

// Send takes a pointer to gmail.Service and is called on a Config type to send an email
func (config *Config) Send() (err error) {
	var message gmail.Message

	emailTo := "To: " + config.EmailTo + "\r\n"
	emailFrom := "From: " + config.ReplyTo + "<" + config.EmailFrom + ">" + "\r\n"
	replyTo := "Reply-To: " + config.ReplyTo + "\r\n"
	subject := "Subject: " + config.Subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + emailFrom + replyTo + subject + mime + "\n" + config.Body)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = config.Service.Users.Messages.Send("me", &message).Do()
  return
}
