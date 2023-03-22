package services

type NotificationService interface {
	SendEmail() error
	SendSms() error
}
