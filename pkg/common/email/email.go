
package email

import (
    "log"
		"bytes"
		"html/template"
		"fmt"
		"gopkg.in/gomail.v2"
		"github.com/spf13/viper"
		"github.com/Horizontal-org/tella-feedback/pkg/common/models"
		"path/filepath"
		"os"
)

func Init() *gomail.Dialer {
		return gomail.NewDialer(
			"email.wearehorizontal.org", 
			25, 
			viper.Get("SMTP_USER").(string), 
			viper.Get("SMTP_PASS").(string),
		)
}

type FeedbackBody struct {
	Text, Platform string
}

func ReportFeedback(dialer *gomail.Dialer, opinion *models.Opinion) {
	
	t := template.New("feedback.html")

	body := FeedbackBody{
		opinion.Text,
		opinion.Platform,
	}

	cwd, _ := os.Getwd()
	var err error
	t, err = t.ParseFiles(filepath.Join( cwd, "./templates/feedback.html" ) )
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, body); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	

	log.Println("Sending")
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@support.tella-app.org")
	m.SetHeader("To", "contact@tella-app.org")
	m.SetHeader("Subject", "New feedback")
	m.SetBody("text/html", result)

	if err := dialer.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Print("Sent")
	}
}