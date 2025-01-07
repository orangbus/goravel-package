package rabbitmq

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "rabbitmq"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewClient(), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {

}
