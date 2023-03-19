package billpaymentservice

import (
	"cqrs-go/cmd/domain/entities/payment"
	"cqrs-go/cmd/domain/entities/transaction"
)

func (bpService *billPaymentService) RegisterPayment(payment payment.Payment) (*transaction.Transaction, error) {
	return nil, nil
}
