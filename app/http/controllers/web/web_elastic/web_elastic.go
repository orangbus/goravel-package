package web_elastic

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/resp"
	facades2 "goravel/packages/elastic/facades"
)

type WebElastic struct {
	// Dependent services
}

func NewWebElastic() *WebElastic {
	return &WebElastic{
		// Inject services
	}
}

func (r *WebElastic) Index(ctx http.Context) http.Response {
	return nil
}
func (r *WebElastic) Version(ctx http.Context) http.Response {
	version, err := facades2.Elastic().Version()
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Data(ctx, version)
}
func (r *WebElastic) Search(ctx http.Context) http.Response {
	keyword := ctx.Request().Query("keyword")
	indexName := "movies"

	query := map[string]interface{}{}
	query["query"] = map[string]interface{}{
		"match": map[string]interface{}{
			"vod_name": keyword,
		},
	}

	list, total, err := facades2.Elastic().Search(indexName, query, 1)
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.List(ctx, list, total)
}

func (r *WebElastic) IndexList(context http.Context) http.Response {
	return nil
}

func (r *WebElastic) IndexCreate(context http.Context) http.Response {
	return nil
}

func (r *WebElastic) IndexDelete(context http.Context) http.Response {
	return nil
}
