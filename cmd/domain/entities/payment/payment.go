package payment

import (
	entity "cqrs-go/cmd/domain/entities"
	domainexceptions "cqrs-go/cmd/domain/exceptions"
)

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

	err := newPay.validate()

	if err != nil {
		return nil, err
	}

	return newPay, nil
}

func (p *Payment) validate() error {

	err := p.validateAccounts()

	if err != nil {
		return err
	}

	err = p.validateValue()

	return err
}

func (p *Payment) validateAccounts() error {
	if p.ReceiveIdAccount == p.PaymentIdAccount {
		return domainexceptions.PaymentAccountCannotBeSame()
	}

	return nil
}

func (p *Payment) validateValue() error {
	if p.Value <= 0 {
		return domainexceptions.PaymentCannotBeNegative()
	}

	return nil
}
