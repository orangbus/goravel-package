package web_spider

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/orangbus/spider/facades"
	"goravel/app/utils/resp"
)

type WebSpider struct {
	// Dependent services
}

func NewWebSpider() *WebSpider {
	return &WebSpider{
		// Inject services
	}
}

func (r *WebSpider) Ping(ctx http.Context) http.Response {
	status := facades.Spider().BaseUrl("https://ccc.com").Ping()
	return resp.Data(ctx, status)
}

func (r *WebSpider) Index(ctx http.Context) http.Response {
	result, err := facades.Spider().BaseUrl("https://xx.com").GetList(1)
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.List(ctx, result.List, result.Total)
}

func (r *WebSpider) CateList(ctx http.Context) http.Response {
	result, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").GetCateList()
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Data(ctx, result)
}
