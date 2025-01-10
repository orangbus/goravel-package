package contracts

type Rabbitmq interface {
	Msg(msg any) error
	Publish(exchangeName, exchangeType string, data interface{}) error
	Routing(exchangeName, key string, data interface{}) error
	Topic(exchangeName, key string, data interface{}) error
	Copsume()
}
