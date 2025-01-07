package web_axios

import (
	"github.com/goravel/framework/contracts/http"
)

type WebAxios struct {
	// Dependent services
}

func NewWebAxios() *WebAxios {
	return &WebAxios{
		// Inject services
	}
}

func (r *WebAxios) Index(ctx http.Context) http.Response {
	return nil
}
