package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// func RenderTemplate()  {

// }

func SendEmail(name string, useremail string, token string)  {
	fmt.Println(name, useremail,token)
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port =  465
	server.Username="oluwajuwonfalore"
	server.Password = "winners2020"
	server.KeepAlive=false
	server.ConnectTimeout=30*time.Second
	server.SendTimeout=30*time.Second
	//server.Encryption = mail.EncryptionSTARTTLS
	server.Encryption = mail.EncryptionSSL

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo(useremail)
	email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	t,_:= template.ParseFiles("./email-templates/verify.html")
	var body bytes.Buffer
	 
	t.Execute(&body, struct{
		Name string
		Email string
		Token string
	}{
		Name:name,
		Email:useremail,
		Token:token,
	})

	email.SetBody(mail.TextHTML,  body.String())
	 

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
	
}

func SendEmailToken(name string, useremail string, token string)  {
	fmt.Println(name, useremail,token)
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port =  1025
	server.KeepAlive=false
	server.ConnectTimeout=10*time.Second
	server.SendTimeout=10*time.Second
	//server.Encryption = mail.EncryptionSSL

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo(useremail)
	email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	t,_:= template.ParseFiles("./email-templates/passwordtoken.html")
	var body bytes.Buffer
	 
	t.Execute(&body, struct{
		Name string
		Email string
		Token string
	}{
		Name:name,
		Email:useremail,
		Token:token,
	})

	email.SetBody(mail.TextHTML,  body.String())
	 

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
	
}