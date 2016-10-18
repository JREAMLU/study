package main

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

var list []struct {
	Name    string
	Address string
}

func main() {
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	for _, r := range list {
		m.SetHeader("From", "no-reply@plu.cn")
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", r.Name))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
