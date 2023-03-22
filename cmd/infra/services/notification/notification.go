package notification

import "cqrs-go/cmd/domain/services"

type notificationService struct {
}

func New() services.NotificationService {
	return &notificationService{}
}
