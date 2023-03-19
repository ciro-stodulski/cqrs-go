package http

import (
	domainexceptions "cqrs-go/cmd/domain/exceptions"
	controllers_http "cqrs-go/cmd/presentation/http/controllers"

	"github.com/gin-gonic/gin"
)

func loadRoutes(controllers []controllers_http.Controller, api gin.RouterGroup) {
	for _, ctr := range controllers {
		route := ctr
		route_config := ctr.LoadRoute()

		if route_config.PathRoot == "" {
			return
		}

		api_group := api.Group(route_config.PathRoot)

		loadMiddlewares(route_config, api_group)

		function := func(gin_context *gin.Context) {
			params := loadParams(gin_context)

			if route_config.Dto != nil {
				if err := gin_context.BindJSON(&route_config.Dto); err != nil {
					return
				}
			}

			result, err := route.Handle(controllers_http.HttpRequest{
				Body:    route_config.Dto,
				Params:  params,
				Query:   gin_context.Request.URL.Query(),
				Headers: gin_context.Request.Header,
			})

			if err != nil {

				var result_error *controllers_http.HttpResponseError

				if appErr, ok := err.(*domainexceptions.ApplicationException); ok {
					result_error = route.HandleError(appErr, nil)
				} else {
					result_error = route.HandleError(nil, appErr)
				}

				data := &controllers_http.HttpResponseError{
					Data:   controllers_http.HttpError{Code: "INTERNAL_SERVER_ERROR", Message: "internal server error"},
					Status: 500,
				}

				if result_error != nil {
					data = result_error
				}

				gin_context.JSON(data.Status, data.Data)
			} else {
				if result.Headers != nil {
					for _, header := range result.Headers {
						gin_context.Header(header.Key, header.Value)
					}
				}

				status := 200

				if result.Status != 0 {
					status = result.Status
				}

				if result.Data != nil {
					gin_context.JSON(status, result.Data)
					return
				}

				gin_context.Status(status)
			}
		}

		switch route_config.Method {
		case "get":
			api_group.GET(route_config.Path, function)
		case "post":
			api_group.POST(route_config.Path, function)
		case "put":
			api_group.PUT(route_config.Path, function)
		case "patch":
			api_group.PATCH(route_config.Path, function)
		case "delete":
			api_group.DELETE(route_config.Path, function)
		default:
		}
	}
}

func loadParams(context *gin.Context) controllers_http.Params {
	var params controllers_http.Params

	if context.Params != nil {
		for _, param := range context.Params {
			param := controllers_http.Param{
				Key:   param.Key,
				Value: param.Value,
			}

			params = append(params, param)
		}
	}

	return params
}

func loadMiddlewares(route controllers_http.CreateRoute, api_group *gin.RouterGroup) {
	if len(route.Middlewares) > 0 {

		for _, mds := range route.Middlewares {
			middleware := func(gin_context *gin.Context) {
				params := loadParams(gin_context)

				if route.Dto != nil {
					if err := gin_context.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				mds(controllers_http.HttpRequest{
					Params:  params,
					Query:   gin_context.Request.URL.Query(),
					Headers: gin_context.Request.Header,
					Next:    gin_context.Next,
					Body:    route.Dto,
				})
			}

			api_group.Use(middleware)
		}
	}
}
