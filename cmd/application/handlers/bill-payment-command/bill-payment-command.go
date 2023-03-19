package billpaymenthandlers

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"cqrs-go/cmd/domain/services"
	"fmt"
)

type billPaymentCommandHandlers struct {
	accountService     services.AccountService
	billPaymentService services.BillPaymentService
}

func New(accountService services.AccountService, billPaymentService services.BillPaymentService) bus.Handler {
	return &billPaymentCommandHandlers{
		accountService:     accountService,
		billPaymentService: billPaymentService,
	}
}

func (bpH *billPaymentCommandHandlers) Perform(command bus.Command) error {

	fmt.Println(command.Data().(params.CreateBillPaymentParams).AccountId)

	bpH.accountService.GetAccountByDocument("test")

	return nil
}
