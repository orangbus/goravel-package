package test

import (
	"github.com/orangbus/axios"
	"testing"
)

func TestGet(t *testing.T) {
	res, err := axios.NewAxios().Get("https://httpbin.org/get", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(res))
}
