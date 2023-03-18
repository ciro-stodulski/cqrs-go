package transaction

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/enums"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	t.Run("Create new transaction successfully:", func(t *testing.T) {

		fakeAccount := entity.NewID()
		value := 100

		fakeTransaction, err := New(fakeAccount, value, enums.Pix)

		assert.NotNil(t, fakeTransaction)
		assert.Nil(t, err)
		assert.Equal(t, fakeTransaction.AccountId, fakeAccount)
		assert.Equal(t, fakeTransaction.Kind, enums.Pix)
		assert.Equal(t, fakeTransaction.Value, value)
	})
}
