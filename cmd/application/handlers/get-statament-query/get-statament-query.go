package getstatementqueryhandlers

import (
	"cqrs-go/cmd/domain/bus"
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/services"
)

type getStatementQueryHandler struct {
	accountService   services.AccountService
	statementService services.StatementService
}

func New(accountService services.AccountService, statementService services.StatementService) bus.QueryHandle {
	return &getStatementQueryHandler{
		accountService:   accountService,
		statementService: statementService,
	}
}

func (gsqh *getStatementQueryHandler) Perform(query any) (any, error) {
	return gsqh.statementService.GetStatementByAccountId(entity.NewID())
}
