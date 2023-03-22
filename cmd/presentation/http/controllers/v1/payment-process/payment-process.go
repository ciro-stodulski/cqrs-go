package controllerbillpayment

import (
	appcommmands "cqrs-go/cmd/application/commands"
	"cqrs-go/cmd/domain/bus"
	entity "cqrs-go/cmd/domain/entities"
	domainexceptions "cqrs-go/cmd/domain/exceptions"
	"cqrs-go/cmd/domain/params"
	"cqrs-go/cmd/presentation/http/controllers"

	"log"
)

type (
	registerController struct {
		bus bus.Bus
	}
)

func New(bus bus.Bus) controllers.Controller {
	return &registerController{bus}
}

func (rc *registerController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/payment-process",
		Method:   "post",
		Path:     "/",
	}
}

func (rc *registerController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {

	command := appcommmands.NewCreateBillPaymentCommand(params.CreateBillPaymentParams{
		AccountId:              entity.NewID(),
		ReceiveDocumentAccount: "01401507464",
		Value:                  500,
	})

	rc.bus.AsyncDispatchCommand(command)

	return &controllers.HttpResponse{
		Status: 202,
	}, nil
}

func (rc *registerController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
	log.Default().Println(appErr.Code)

	return nil
}
