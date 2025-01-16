package contracts

import (
	"github.com/orangbus/goravel-spider"
	"github.com/orangbus/goravel-spider/pkg/movie_spider"
)

type Spider interface {
	BaseUrl(base_url string) *spider.Spider
	SetHour(hour int) *spider.Spider
	SetType(type_id int) *spider.Spider

	GetCateList() ([]movie_spider.ClassList, error)
	GetList(page int, limit ...int) (movie_spider.MovieResponse, error)
	Search(keyword string, page int, limit ...int) (movie_spider.MovieResponse, error)
	Detail(ids string) (movie_spider.MovieResponse, error)
	Ping() bool
}
