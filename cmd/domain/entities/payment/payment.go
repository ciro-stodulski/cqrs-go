package payment

import entity "cqrs-go/cmd/domain/entities"

type Payment struct {
	ReceiveIdAccount entity.ID
	PaymentIdAccount entity.ID
	Value            int
}

func New(receiveIdAccount entity.ID, paymentIdAccount entity.ID, value int) (*Payment, error) {
	newPay := &Payment{
		ReceiveIdAccount: receiveIdAccount,
		PaymentIdAccount: paymentIdAccount,
		Value:            value,
	}

	return newPay, nil
}
