package web_axios

import (
	"encoding/json"
	"github.com/goravel/framework/contracts/http"
	"github.com/orangbus/axios/facades"
	"goravel/app/utils/resp"
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

func (r *WebAxios) Post(ctx http.Context) http.Response {
	data, err := facades.Axios().SetHeader(map[string]string{
		"Content-Type": "application/json",
	}).Authorization("1231").VerifyHttps(true).HttpBin().Post(map[string]any{
		"name": "orangbus",
	})
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	var respData map[string]any
	json.Unmarshal(data, &respData)
	return resp.Data(ctx, respData)
}

func (r *WebAxios) PostForm(ctx http.Context) http.Response {
	data, err := facades.Axios().Authorization("1231").VerifyHttps(true).HttpBin().Post(map[string]any{"name": "orangbus"})
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	var respData map[string]any
	json.Unmarshal(data, &respData)
	return resp.Data(ctx, respData)
}

func (r *WebAxios) HttpTest(ctx http.Context) http.Response {
	data, err := facades.Axios().Get("https://api.guangsuapi.com/api.php/provide/vod", map[string]any{"pg": "1"})
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	var respData map[string]any
	json.Unmarshal(data, &respData)
	return resp.Data(ctx, respData)
}
