package commandbillpaymenthandler

import (
	appsevents "cqrs-go/cmd/application/events"
	"cqrs-go/cmd/domain/bus"
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/payment"
	"cqrs-go/cmd/domain/entities/transaction"
	"cqrs-go/cmd/domain/enums"
	"cqrs-go/cmd/domain/params"
	"cqrs-go/cmd/domain/services"
	"time"
)

type billPaymentCommandHandlers struct {
	accountService     services.AccountService
	billPaymentService services.BillPaymentService
	eventBus           bus.EventBus
}

func New(accountService services.AccountService, billPaymentService services.BillPaymentService, eventBus bus.EventBus) bus.CommandHandler {
	return &billPaymentCommandHandlers{
		accountService:     accountService,
		billPaymentService: billPaymentService,
		eventBus:           eventBus,
	}
}

func (bpH *billPaymentCommandHandlers) Perform(command bus.Command) error {
	paramscbpp := command.Data().(params.CreateBillPaymentParams)

	bpH.accountService.GetAccountByDocument(paramscbpp.ReceiveDocumentAccount)

	bpH.billPaymentService.RegisterPayment(payment.Payment{
		ReceiveIdAccount: paramscbpp.AccountId,
		PaymentIdAccount: paramscbpp.AccountId,
		Value:            paramscbpp.Value,
	})

	transactionEvent := appsevents.NewTransactionEvent(params.NotificationTransactionParams{
		Date: time.Now(),
		Trasaction: transaction.Transaction{
			AccountId: entity.NewID(),
			Value:     1000,
			Kind:      enums.Bill,
			Status:    enums.Successful,
		},
	})

	bpH.eventBus.DispatchEvent(transactionEvent)

	return nil
}
