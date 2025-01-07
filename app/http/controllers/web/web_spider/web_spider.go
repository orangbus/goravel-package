package web_spider

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/resp"
	"goravel/packages/spider/facades"
)

type WebSpider struct {
	// Dependent services
}

func NewWebSpider() *WebSpider {
	return &WebSpider{
		// Inject services
	}
}

func (r *WebSpider) Index(ctx http.Context) http.Response {
	result, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").GetList(1)
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
