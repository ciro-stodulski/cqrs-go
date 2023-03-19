package accountservice

import (
	"cqrs-go/cmd/domain/entities/account"
	"fmt"
)

func (aService *accountService) GetAccountByDocument(document string) (*account.Account, error) {

	fmt.Println(document)

	return nil, nil
}
