package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	"goravel/app/http/controllers/web/web_axios"
	"goravel/app/http/controllers/web/web_elastic"
	"goravel/app/http/controllers/web/web_spider"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	elasticController := web_elastic.NewWebElastic()
	facades.Route().Prefix("elastic").Group(func(router route.Router) {
		router.Get("version", elasticController.Version)
		router.Get("search", elasticController.Search)

		router.Prefix("index").Group(func(router route.Router) {
			router.Post("mapping", elasticController.Mapping)
			router.Get("create", elasticController.IndexCreate)
			router.Get("delete", elasticController.IndexDelete)
		})
	})

	axiosController := web_axios.NewWebAxios()
	facades.Route().Prefix("axios").Group(func(router route.Router) {
		router.Get("get", axiosController.Index)
		router.Get("post", axiosController.Index)
	})

	spiderController := web_spider.NewWebSpider()
	facades.Route().Prefix("spider").Group(func(router route.Router) {
		router.Get("list", spiderController.Index)
		router.Get("cate", spiderController.CateList)
	})
	//
	//facades.Route().Prefix("rabbitmq").Group(func(router route.Router) {
	//
	//})
}
