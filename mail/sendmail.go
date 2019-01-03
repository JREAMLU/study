package main

import (
	"log"

	"github.com/matcornic/hermes"
	"gopkg.in/gomail.v2"
)

func main() {
	fromAccount := "账号@qq.com"
	password := "密码"
	address := "发送的邮箱"

	d := gomail.NewDialer("smtp.qq.com", 587, fromAccount, password)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	// QQ邮箱 From 必须和发送的邮箱保持一致
	m.SetHeader("From", fromAccount)
	// m.SetHeader("From", "no-reply@xxx.cn")
	m.SetHeader("To", address)
	// m.SetAddressHeader("Cc", "", "dan")
	m.SetHeader("Subject", "Newsletter #1")

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "YaoMing",
			Copyright: "Copyright © 2019. All rights reserved.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Signature: "Yours truly, YaoMing",
			FreeMarkdown: `
> _Hermes_ service will shutdown the **1st August 2017** for maintenance operations. 

Services will be unavailable based on the following schedule:

| Services | Downtime |
| :------:| :-----------: |
| Service A | 2AM to 3AM |
| Service B | 4AM to 5AM |
| Service C | 5AM to 6AM |

---

Feel free to contact us for any question regarding this matter at [support@hermes-example.com](mailto:support@hermes-example.com) or in our [Gitter](https://gitter.im/)

`,
		},
	}

	body, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}

	m.SetBody("text/html", body)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send email to %q: %v", address, err)
	}
	m.Reset()
}
