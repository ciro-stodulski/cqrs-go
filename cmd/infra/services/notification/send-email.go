package notification

import "log"

func (ns *notificationService) SendEmail() error {
	log.Default().Println("Sending email")
	return nil
}
