package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

type Elastic struct {
	client *elasticsearch.TypedClient
}

func NewClient() (*Elastic, error) {
	hosts := []string{}
	host := cast.ToString(facades.Config().Get("elastic.host"))
	if strings.Contains(host, ",") {
		split := strings.Split(host, ",")
		for _, h := range split {
			hosts = append(hosts, h)
		}
	} else {
		hosts = append(hosts, host)
	}
	c := elasticsearch.Config{
		Addresses: hosts,
		Username:  cast.ToString(facades.Config().GetString("elastic.username")),
		Password:  cast.ToString(facades.Config().GetString("elastic.password")),
		Transport: &http.Transport{
			MaxIdleConns:        100,              // 最大空闲连接数
			MaxIdleConnsPerHost: 2,                // 每个主机的最大空闲连接数
			IdleConnTimeout:     time.Second * 10, // 空闲连接超时时间
		},
	}
	client, err := elasticsearch.NewTypedClient(c)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("elasticsearch 连接错误：%s", err.Error()))
	}

	es := Elastic{client: client}
	return &es, err
}

func (e *Elastic) Version() (*info.Response, error) {
	return e.client.Info().Do(context.Background())
}

func (e *Elastic) Search(indexName string, query map[string]interface{}, page int, limit ...int) ([]interface{}, int64, error) {
	var list []interface{}
	var total int64
	var size int

	if page <= 1 {
		page = 1
	}
	if len(limit) > 0 {
		size = limit[0]
	}
	if size <= 0 || size > 100 {
		size = 20
	}
	query["from"] = (size - 1) * size
	query["size"] = size

	marshal, err := json.Marshal(query)
	if err != nil {
		return list, total, err
	}
	reader := bytes.NewReader(marshal)
	res, err2 := e.client.Search().Index(indexName).Raw(reader).Do(context.Background())
	if err2 != nil {
		return list, total, err
	}
	for _, item := range res.Hits.Hits {
		var data interface{}
		err := json.Unmarshal(item.Source_, &data)
		if err == nil {
			list = append(list, data)
		} else {
			facades.Log().Warning("elasticsearch Unmarshal error: %s", err.Error())
		}
	}
	total = res.Hits.Total.Value
	return list, total, nil
}
