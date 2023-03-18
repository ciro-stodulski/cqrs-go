package payment

import (
	entity "cqrs-go/cmd/domain/entities"
	domainexceptions "cqrs-go/cmd/domain/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayment(t *testing.T) {
	t.Run("Create new payment successfully:", func(t *testing.T) {
		fakeReceiveIdAccount := entity.NewID()
		fakePaymentIdAccount := entity.NewID()
		value := 10

		fakePayment, error := New(fakeReceiveIdAccount, fakePaymentIdAccount, value)

		assert.NotNil(t, fakePayment)
		assert.Nil(t, error)
		assert.Equal(t, fakePayment.ReceiveIdAccount, fakeReceiveIdAccount)
		assert.Equal(t, fakePayment.PaymentIdAccount, fakePaymentIdAccount)
		assert.Equal(t, fakePayment.Value, value)
	})

	t.Run("Return error account cannot be the same", func(t *testing.T) {
		id := entity.NewID()
		value := 10

		fakePayment, err := New(id, id, value)

		assert.NotNil(t, err)
		assert.Nil(t, fakePayment)
		assert.Equal(t, err, domainexceptions.PaymentAccountCannotBeSame())
	})

	t.Run("Return error value cannot be negative", func(t *testing.T) {
		fakeReceiveIdAccount := entity.NewID()
		fakePaymentIdAccount := entity.NewID()
		value := -10

		fakePayment, err := New(fakeReceiveIdAccount, fakePaymentIdAccount, value)

		assert.NotNil(t, err)
		assert.Nil(t, fakePayment)
		assert.Equal(t, err, domainexceptions.PaymentCannotBeNegative())
	})
}
