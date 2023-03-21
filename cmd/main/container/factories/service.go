package factories

import (
	"cqrs-go/cmd/domain/services"
	accountservice "cqrs-go/cmd/infra/services/account"
	billpaymentservice "cqrs-go/cmd/infra/services/bill-payment"
	statementservice "cqrs-go/cmd/infra/services/statement"
)

type (
	ServiceCaseContext struct {
		AccountService     services.AccountService
		BillPaymentService services.BillPaymentService
		StatementService   services.StatementService
	}
)

func MakeServiceContext(infra_context InfraContext) ServiceCaseContext {
	return ServiceCaseContext{
		AccountService:     accountservice.New(),
		BillPaymentService: billpaymentservice.New(),
		StatementService:   statementservice.New(),
	}
}
