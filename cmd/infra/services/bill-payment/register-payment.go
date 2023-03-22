package billpaymentservice

import (
	"cqrs-go/cmd/domain/entities/payment"
	"cqrs-go/cmd/domain/entities/transaction"
	"log"
)

func (bpService *billPaymentService) RegisterPayment(payment payment.Payment) (*transaction.Transaction, error) {
	log.Default().Println("Starting register bill-payment")

	return nil, nil
}
