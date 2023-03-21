package http

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/presentation/http/controllers"
	controllerbillpayment "cqrs-go/cmd/presentation/http/controllers/v1/bill-payment"
	controllergetstatement "cqrs-go/cmd/presentation/http/controllers/v1/get-statement"
)

func loadControllers(b bus.Bus) []controllers.Controller {
	return []controllers.Controller{
		controllerbillpayment.New(b),
		controllergetstatement.New(b),
	}
}

func loadMiddlewaresGlobals() []controllers.Middleware {
	return []controllers.Middleware{}
}
