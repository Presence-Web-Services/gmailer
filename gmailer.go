/*
Gmailer allows for OAUTH authentication to GMail and sending of basic email messages
*/
package gmailer

import (
  "context"
  "time"
  "encoding/base64"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Config is type containing OAuth credentials and email message information
type Config struct {
  clientID string
  clientSecret string
  accessToken string
  refreshToken string
  emailTo string
  emailFrom string
  replyTo string
  subject string
  body string
}

// Authenticate authenticates with GMail OAuth, and returns a pointer to gmail.Service when successfully authenticated
func (config *Config) Authenticate() (gmailService *gmail.Service, err error) {
  oauthConfig := oauth2.Config{
    ClientID:     config.clientID,
    ClientSecret: config.clientSecret,
    Endpoint:     google.Endpoint,
    RedirectURL:  "http://localhost",
  }
  oauthToken := oauth2.Token{
    AccessToken:  config.accessToken,
    RefreshToken: config.refreshToken,
    TokenType:    "Bearer",
    Expiry:       time.Now(),
  }

  tokenSource := oauthConfig.TokenSource(context.Background(), &oauthToken)

  gmailService, err = gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
  return
}

// Send takes a pointer to gmail.Service and is called on a Config type to send an email
func (config *Config) Send(gmailService *gmail.Service) (err error) {
  var message gmail.Message

  emailTo := "To: " + config.emailTo + "\r\n"
  emailFrom := "From " + config.emailFrom + "\r\n"
  replyTo := "ReplyTo " + config.replyTo + "\r\n"
  subject := "Subject: " + config.subject + "\n"
  mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
  msg := []byte(emailTo + emailFrom + replyTo + subject + mime + "\n" + config.body)

  message.Raw = base64.URLEncoding.EncodeToString(msg)

  // Send the message
  _, err = gmailService.Users.Messages.Send("me", &message).Do()
  return
}
