package eventnotificationtransactionhandler

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/services"
)

type notificationTransaction struct {
	notificationService services.NotificationService
}

func New(notificationService services.NotificationService) bus.EventHandle {
	return &notificationTransaction{
		notificationService: notificationService,
	}
}

func (ent *notificationTransaction) Perform(bus.Event) error {

	return ent.notificationService.SendEmail()
}
