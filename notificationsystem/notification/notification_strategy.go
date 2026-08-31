package notification

import "fmt"

type email struct {
	emailId string
}

func NewEmail(id string) *email {
	return &email{
		emailId: id,
	}
}

func (e *email) SendNotification(content string) {
	fmt.Println("Sending email notification to:", e.emailId, "\n", content)
}

type sms struct {
	mobileNumber string
}

func NewSMS(mobile string) *sms {
	return &sms{mobileNumber: mobile}
}

func (s *sms) SendNotification(content string) {
	fmt.Println("Sending SMS notification to:", s.mobileNumber, "\n", content)
}
