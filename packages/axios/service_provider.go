package axios

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "axios"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewAxios(), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/orangbus/goravel-axios", map[string]string{
		"config/axios.go": app.ConfigPath("axios.go"),
	})
}
