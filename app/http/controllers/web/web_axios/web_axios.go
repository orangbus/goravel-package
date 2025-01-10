package web_axios

import (
	"encoding/json"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/resp"
	"goravel/packages/axios/facades"
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
	data, err := facades.Axios().SetHeader(map[string]string{
		"Content-Type": "application/json",
	}).Authorization("1231").VerifyHttps(true).HttpBin().Get(map[string]any{
		"name": "orangbus",
	})
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	var respData map[string]any
	json.Unmarshal(data, &respData)
	return resp.Data(ctx, respData)
}
