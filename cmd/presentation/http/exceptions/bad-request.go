package httpexceptions

import "cqrs-go/cmd/presentation/http/controllers"

func BadRequest(data controllers.HttpError) *controllers.HttpResponseError {
	return &controllers.HttpResponseError{
		Data:   data,
		Status: 400,
	}
}
