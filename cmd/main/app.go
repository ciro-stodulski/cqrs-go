package app

import (
	"cqrs-go/cmd/main/bus"
	"cqrs-go/cmd/main/container"
	"cqrs-go/cmd/main/modules"
	"cqrs-go/cmd/main/modules/http"
	"cqrs-go/cmd/shared/env"
)

type App struct {
	modules []modules.Module
}

func (app *App) start() error {
	var err error

	for i := 0; i < len(app.modules); i++ {
		if app.modules[i].RunGo() {
			go app.modules[i].Start()
		} else {
			err = app.modules[i].Start()
		}
	}

	return err
}

func New() error {
	env.Load()

	c := container.New()

	app := &App{
		modules: []modules.Module{
			http.New(bus.New(c)),
		},
	}

	err := app.start()

	return err
}
