package account

import (
	entity "cqrs-go/cmd/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("Create new account successfully:", func(t *testing.T) {
		fakeId := entity.NewID()
		fakeName := "yolo"
		fakeBalance := 100
		fakeDocument := "01574103089"

		fakeAccount := New(fakeId, fakeName, fakeBalance, fakeDocument)

		assert.NotNil(t, fakeAccount)
		assert.Equal(t, fakeAccount.Name, fakeName)
		assert.Equal(t, fakeAccount.Document, fakeDocument)
		assert.Equal(t, fakeAccount.Balance, fakeBalance)
	})
}
