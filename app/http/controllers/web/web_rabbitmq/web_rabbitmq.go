package web_rabbitmq

import (
	"github.com/goravel/framework/contracts/http"
)

type WebRabbitmq struct {
	// Dependent services
}

func NewWebRabbitmq() *WebRabbitmq {
	return &WebRabbitmq{
		// Inject services
	}
}

func (r *WebRabbitmq) Index(ctx http.Context) http.Response {
	return nil
}
