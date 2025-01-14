package spider

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"goravel/packages/spider/pkg/movie_spider"
	"io"
	"net/http"
	"net/url"
)

type Spider struct {
	baseUrl string
	page    int
	hour    int
	tp      int
	ac      string
	keyword string
	ids     string
}

func NewSpider() *Spider {
	return &Spider{page: 1, ac: "list"}
}
func (s *Spider) SetAcVideoList() *Spider {
	s.ac = "videolist"
	return s
}

func (s *Spider) BaseUrl(base_url string) *Spider {
	s.baseUrl = base_url
	return s
}

func (s *Spider) SetHour(hour int) *Spider {
	s.hour = hour
	return s
}
func (s *Spider) SetType(t int) *Spider {
	s.tp = t
	return s
}

func (s *Spider) get() (movie_spider.MovieResponse, error) {
	param := url.Values{}
	var data movie_spider.MovieResponse

	param.Set("ac", s.ac)
	if s.page > 0 {
		param.Set("pg", cast.ToString(s.page))
	}
	if s.hour > 0 {
		param.Set("h", cast.ToString(s.hour))
	}
	if s.keyword != "" {
		param.Set("wd", s.keyword)
	}
	response, err := http.Get(fmt.Sprintf("%s?%s", s.baseUrl, param.Encode()))
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (s *Spider) GetList(page int) (movie_spider.MovieResponse, error) {
	if page <= 0 {
		page = 1
	}
	s.page = page
	return s.get()
}

func (s *Spider) Search(keyword string) (movie_spider.MovieResponse, error) {
	s.keyword = keyword
	return s.get()
}

func (s *Spider) Detail(ids string) (movie_spider.MovieResponse, error) {
	s.ids = ids
	return s.get()
}

func (s *Spider) GetCateList() ([]movie_spider.ClassList, error) {
	resp, err := s.get()
	if err != nil {
		return []movie_spider.ClassList{}, err
	}
	return resp.Class, nil
}

func (s *Spider) Ping(keyword string) bool {
	return true
}
