package axios

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Axios struct {
	base_url    string
	body        map[string]any
	headers     map[string]string
	proxy       string
	httpsVerify bool
	timeout     int
}

func NewAxios() *Axios {
	axios := Axios{
		timeout: 30,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return &axios
}

func (a *Axios) HttpBin() *HttpBin {
	return NewHttpBin(a)
}

func (a *Axios) SetHeader(headers map[string]string) *Axios {
	a.headers = headers
	return a
}

func (a *Axios) SetProxy(proxy string) *Axios {
	a.proxy = proxy
	return a
}

func (a *Axios) VerifyHttps(verify bool) *Axios {
	a.httpsVerify = verify
	return a
}

func (a *Axios) Authorization(token string) *Axios {
	a.headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return a
}
func (a *Axios) Get(base_url string, param map[string]any) ([]byte, error) {
	a.base_url = base_url
	a.body = param
	return a.builder("GET")
}
func (a *Axios) Post(base_url string, param map[string]any) ([]byte, error) {
	a.base_url = base_url
	a.body = param
	return a.builder("POST")
}

func (a *Axios) PostForm(base_url string, param map[string]any) ([]byte, error) {
	a.base_url = base_url
	a.body = param
	a.headers["Content-Type"] = "application/x-www-form-urlencoded"
	return a.builder("POST")
}

func (a *Axios) Dd() map[string]interface{} {
	body := map[string]interface{}{}
	body["header"] = a.headers
	body["body"] = a.body
	body["proxy"] = a.proxy
	body["httpsVerify"] = a.httpsVerify
	body["timeout"] = a.timeout
	return body
}

func (a *Axios) builder(method string) ([]byte, error) {
	tr := &http.Transport{}
	if a.httpsVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	if a.proxy != "" {
		parse, err := url.Parse(a.proxy)
		if err != nil {
			return nil, err
		}
		tr.Proxy = http.ProxyURL(parse)
	}

	req_url := a.base_url
	if len(a.body) > 0 {
		params := url.Values{}
		for k, v := range a.body {
			params.Add(k, cast.ToString(v))
		}
		req_url = fmt.Sprintf("%s?%s", a.base_url, params.Encode())
	}

	req, err := http.NewRequest(method, req_url, nil)
	if err != nil {
		return nil, err
	}

	if len(a.headers) > 0 {
		for k, v := range a.headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Transport: tr, Timeout: time.Duration(a.timeout) * time.Second}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("请求错误，错误状态码:%d", response.StatusCode))
	}
	return io.ReadAll(response.Body)
}
