package accountservice

import (
	"cqrs-go/cmd/domain/entities/account"
	"log"
)

func (aService *accountService) GetAccountByDocument(document string) (*account.Account, error) {
	log.Default().Println("Getting account by document: " + document)

	return nil, nil
}
