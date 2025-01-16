package web_rabbitmq

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/orangbus/rabbitmq/facades"
	"goravel/app/utils/resp"
)

type WebRabbitmq struct {
	// Dependent services
}

func NewWebRabbitmq() *WebRabbitmq {
	return &WebRabbitmq{
		// Inject services
	}
}

func (r *WebRabbitmq) Msg(ctx http.Context) http.Response {
	msg := ctx.Request().Input("msg", "默认消息内容")
	err := facades.Rabbitmq().Msg(msg)
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Success(ctx, "success")
}

func (r *WebRabbitmq) Publish(ctx http.Context) http.Response {
	return nil
}
func (r *WebRabbitmq) Routing(ctx http.Context) http.Response {
	return nil
}
func (r *WebRabbitmq) Topic(ctx http.Context) http.Response {
	return nil
}
