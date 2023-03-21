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
		PathRoot: "/v1/bill-payments",
		Method:   "post",
		Path:     "/",
	}
}

func (rc *registerController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {

	command := appcommmands.NewCreateBillPaymentCommand(params.CreateBillPaymentParams{
		AccountId:              entity.NewID(),
		ReceiveDocumentAccount: "0000578974",
		Value:                  200,
	})

	err := rc.bus.SyncDispatchCommand(command)

	if err != nil {
		return nil, err
	}

	return &controllers.HttpResponse{
		Status: 201,
	}, nil
}

func (rc *registerController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
	log.Default().Println(appErr.Code)

	return nil
}
