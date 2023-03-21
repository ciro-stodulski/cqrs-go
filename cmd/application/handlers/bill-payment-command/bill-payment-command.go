package billpaymentcommandhandlers

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"cqrs-go/cmd/domain/services"
)

type billPaymentCommandHandlers struct {
	accountService     services.AccountService
	billPaymentService services.BillPaymentService
}

func New(accountService services.AccountService, billPaymentService services.BillPaymentService) bus.CommandHandler {
	return &billPaymentCommandHandlers{
		accountService:     accountService,
		billPaymentService: billPaymentService,
	}
}

func (bpH *billPaymentCommandHandlers) Perform(command bus.Command) error {
	params := command.Data().(params.CreateBillPaymentParams)

	bpH.accountService.GetAccountByDocument(params.ReceiveDocumentAccount)

	return nil
}
