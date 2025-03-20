package main

import (
	"flag"
	"fmt"
	log "log/slog"

	"gopkg.in/mail.v2"
)

func main() {
	smtpAddr := flag.String("addr", "", "smtp address")
	p := flag.Int("p", 0, "smtp port")
	u := flag.String("u", "", "user")
	pass := flag.String("pass", "", "password")

	from := flag.String("from", "", "from")
	to := flag.String("to", "", "to")

	flag.Parse()

	log.Info(fmt.Sprintf("%s:%s %s:%d", *u, *pass, *smtpAddr, *p))
	log.Info(fmt.Sprintf("from:\t%s\nto:\t%s", *from, *to))

	m := mail.NewMessage()
	m.SetHeader("From", *from)
	m.SetHeader("To", *to)
	m.SetHeader("Subject", "test mail")

	m.SetBody("text/plain", "test email")
	m.AddAlternative("text/html", `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Тестовое письмо</title>
</head>
<body>
    <div class="email-container">
        <div class="email-content">
            <div class="email-body">
                <p>Это тестовое письмо для проверки</p>
            </div>
        </div>
    </div>
</body>
</html>`)

	d := mail.NewDialer(*smtpAddr, *p, *u, *pass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		log.Error(err.Error())
	} else {
		log.Info("sent")
	}
}
