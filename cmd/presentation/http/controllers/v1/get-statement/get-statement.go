package controllergetstatement

import (
	appqueries "cqrs-go/cmd/application/queries"
	"cqrs-go/cmd/domain/bus"
	entity "cqrs-go/cmd/domain/entities"
	domainexceptions "cqrs-go/cmd/domain/exceptions"
	"cqrs-go/cmd/domain/params"
	"cqrs-go/cmd/presentation/http/controllers"
	"log"
)

type (
	getStatementController struct {
		bus bus.Bus
	}
)

func New(bus bus.Bus) controllers.Controller {
	return &getStatementController{bus}
}

func (rc *getStatementController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/statements",
		Method:   "get",
		Path:     "/",
	}
}

func (rc *getStatementController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {
	query := appqueries.NewGetStatementQuery(params.GetStatementQueryParams{
		AccountId: entity.NewID(),
	})

	result, err := rc.bus.DispatchQuery(query)

	if err != nil {
		return nil, err
	}

	return &controllers.HttpResponse{
		Status: 200,
		Data:   result,
	}, nil
}

func (rc *getStatementController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
	log.Default().Println(appErr.Code)

	return nil
}
