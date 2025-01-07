package facades

import (
	"log"

	"goravel/packages/rabbitmq"
	"goravel/packages/rabbitmq/contracts"
)

func Rabbitmq() contracts.Rabbitmq {
	instance, err := rabbitmq.App.Make(rabbitmq.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Rabbitmq)
}
