package services

import (
	"cqrs-go/cmd/domain/entities/payment"
	"cqrs-go/cmd/domain/entities/transaction"
)

type BillPaymentService interface {
	RegisterPayment(payment payment.Payment) (*transaction.Transaction, error)
}
