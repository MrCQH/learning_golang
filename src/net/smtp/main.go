package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	test1()
	test2()
}

// 加密
func test2() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}

// 未加密
func test1() {
	client, err := smtp.Dial("mail.example.com:25")
	checkError(err, "smtp connect")
	// setup the sender
	client.Mail("mrchen")
	// setup the recipient
	client.Rcpt("rezhen")

	dataWriter, err := client.Data()
	checkError(err, "Send Data Body")
	defer dataWriter.Close()
	bufferString := bytes.NewBufferString("This is the email body.")
	_, err = bufferString.WriteTo(dataWriter)
	checkError(err, "Bufferstring WriteTo")
}

func checkError(err error, info string) {
	if err != nil {
		log.Println(info, ": ", err.Error())
	}
}
