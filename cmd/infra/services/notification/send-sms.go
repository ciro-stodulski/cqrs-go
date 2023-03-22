package notification

import "log"

func (ns *notificationService) SendSms() error {
	log.Default().Println("Sending sms")

	return nil
}
