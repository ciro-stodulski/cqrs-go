package factories

import (
	"cqrs-go/cmd/domain/services"
	accountservice "cqrs-go/cmd/infra/services/account"
	billpaymentservice "cqrs-go/cmd/infra/services/bill-payment"
)

type (
	ServiceCaseContext struct {
		AccountService     services.AccountService
		BillPaymentService services.BillPaymentService
	}
)

func MakeServiceContext(infra_context InfraContext) ServiceCaseContext {
	return ServiceCaseContext{
		AccountService:     accountservice.New(),
		BillPaymentService: billpaymentservice.New(),
	}
}
