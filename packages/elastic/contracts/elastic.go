package contracts

import "github.com/elastic/go-elasticsearch/v8/typedapi/core/info"

type Elastic interface {
	Version() *info.Response
	Search(indexName string, query map[string]interface{}, page int, limit ...int) ([]interface{}, int64, error)
}
