# 视频采集

## 快速入门

安装
```bash
go get -u github.com/orangbus/goravel-spider@latest
```

注册
打开 `config/app.go`
```go
import "github.com/orangbus/goravel-spider"
```
```go
"providers": []foundation.ServiceProvider{
	...
    &spider.ServiceProvider{},
},

```
使用
```go
import "github.com/orangbus/goravel-spider/facades"
```
```go
func (r *WebSpider) Ping(ctx http.Context) http.Response {
    status := facades.Spider().BaseUrl("https://xxx.com").Ping()
    return resp.Data(ctx, status)
}
```
