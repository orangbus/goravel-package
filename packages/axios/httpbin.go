package axios

import "fmt"

type HttpBin struct {
	base_url string
	axios    *Axios
}

func NewHttpBin(axios *Axios) *HttpBin {
	return &HttpBin{
		base_url: "https://httpbin.org",
		axios:    axios,
	}
}

func (a *HttpBin) Get(param map[string]any) ([]byte, error) {
	return a.axios.Get(fmt.Sprintf("%s/get", a.base_url), param)
}

func (a *HttpBin) Post(param map[string]any) ([]byte, error) {
	return a.axios.Get(fmt.Sprintf("%s/post", a.base_url), param)
}
