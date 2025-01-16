package web_elastic

import (
	"github.com/goravel/framework/contracts/http"
	elastic "github.com/orangbus/elastic/facades"
	"goravel/app/utils/resp"
	"log"
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
	version, err := elastic.Elastic().Version()
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

	list, total, err := elastic.Elastic().Search(indexName, query, 1)
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.List(ctx, list, total)
}

func (r *WebElastic) Mapping(ctx http.Context) http.Response {
	indexName := ctx.Request().Input("index")
	mapping := map[string]interface{}{}
	param := ctx.Request().Input("mapping")
	log.Printf(param)
	return resp.Success(ctx, param)

	err := elastic.Elastic().Mapping(indexName, mapping)
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Success(ctx, "创建成功")
}

func (r *WebElastic) IndexCreate(ctx http.Context) http.Response {
	index := ctx.Request().Input("index")
	if err := elastic.Elastic().Index().Create(index); err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Success(ctx, "创建成功")
}

func (r *WebElastic) IndexDelete(ctx http.Context) http.Response {
	index := ctx.Request().Input("index")
	if err := elastic.Elastic().IndexDelete(index); err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Success(ctx, "删除成功")
}
