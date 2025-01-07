package resp

import "github.com/goravel/framework/contracts/http"

func Error(ctx http.Context, msg string, code ...int) http.Response {
	errorCode := 202
	if len(code) > 0 {
		errorCode = code[0]
	}
	return ctx.Response().Success().Json(map[string]any{
		"code": errorCode,
		"msg":  msg,
	})
}

func Success(ctx http.Context, data any, msg ...string) http.Response {
	msgTip := "success"
	if len(msg) > 0 {
		msgTip = msg[0]
	}
	return ctx.Response().Success().Json(map[string]any{
		"code": 200,
		"data": data,
		"msg":  msgTip,
	})
}

func Data(ctx http.Context, data any, msg ...string) http.Response {
	msgTip := "success"
	if len(msg) > 0 {
		msgTip = msg[0]
	}
	return ctx.Response().Success().Json(map[string]any{
		"code": 200,
		"data": data,
		"msg":  msgTip,
	})
}

func List(ctx http.Context, list any, total int64) http.Response {
	return ctx.Response().Success().Json(map[string]any{
		"code":  200,
		"data":  list,
		"total": total,
		"msg":   "success",
	})
}
