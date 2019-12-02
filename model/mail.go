package model

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/smtp"

	uuid "github.com/satori/go.uuid"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

type Comms struct {
	Token    string
	OTP      string
	Name     string
	Username string
	Password string
}

type Mail struct {
	from    string
	to      string
	subject string
	body    string
}

func NewMail(to string, subject string) *Mail {
	return &Mail{
		to:      to,
		subject: subject,
	}
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (m *Mail) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	m.body = buffer.String()
	return nil
}

func (m *Mail) sendMail() bool {
	body := "To:" + m.to + "\r\nSubject:" + m.subject + "\r\n" + MIME + "\r\n" + m.body

	auth := smtp.PlainAuth("", "shonphand@gmail.com", "kjfsusizidokwusz", "smtp.gmail.com.")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{m.to}
	err := smtp.SendMail("smtp.gmail.com.:25", auth, "shonphand@gmail.com", to, []byte(body))
	if err != nil {
		return false
	}

	return true
}

func (m *Mail) Send(templateName string, items interface{}) {
	err := m.parseTemplate("/home/shon/Documents/Go_practise/LoginApp/resources/template/"+templateName, items)
	if err != nil {
		fmt.Println("Error in parsing file", err.Error())
	}
	if ok := m.sendMail(); ok {
		log.Printf("Email has been sent to %s\n", m.to)
	} else {
		log.Printf("Failed to send the email to %s\n", m.to)
	}
}

func GenerateOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func GenerateToken() string {
	token := uuid.NewV4().String()
	return token
}
