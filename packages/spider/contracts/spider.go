package contracts

import (
	"goravel/packages/spider"
	"goravel/packages/spider/pkg/movie_spider"
)

type Spider interface {
	BaseUrl(base_url string) *spider.Spider
	GetCateList() ([]movie_spider.ClassList, error)
	GetList(page int) (movie_spider.MovieResponse, error)
	Search(keyword string) (movie_spider.MovieResponse, error)
	Detail(ids string) (movie_spider.MovieResponse, error)
}
