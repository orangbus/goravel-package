package contracts

import "github.com/orangbus/axios"

type Axios interface {
	HttpBin() *axios.HttpBin
	SetHeader(headers map[string]string) *axios.Axios
	SetProxy(proxy string) *axios.Axios
	VerifyHttps(verify bool) *axios.Axios
	Authorization(auth string) *axios.Axios
	Dd() map[string]interface{} // 打印请求参数

	Get(base_url string, param map[string]any) ([]byte, error)     // get 请求
	Post(base_url string, body map[string]any) ([]byte, error)     // post 请求
	PostForm(base_url string, body map[string]any) ([]byte, error) // post 请求

	//Post(api_url string, body []map[string]interface{})   // post 请求
	//Upload(api_url string, body []map[string]interface{}) // 上传文件
}
