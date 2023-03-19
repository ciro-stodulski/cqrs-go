package billpaymentservice

import "cqrs-go/cmd/domain/services"

type billPaymentService struct {
}

func New() services.BillPaymentService {
	return &billPaymentService{}
}
