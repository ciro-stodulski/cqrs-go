package statementservice

import "cqrs-go/cmd/domain/services"

type statementService struct {
}

func New() services.StatementService {
	return &statementService{}
}
