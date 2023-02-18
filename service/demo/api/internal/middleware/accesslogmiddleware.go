package middleware

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"time"
)

type AccessLogMiddleware struct {
}

func NewAccessLogMiddleware() *AccessLogMiddleware {
	return &AccessLogMiddleware{}
}

func (m *AccessLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Passthrough to next handler if need
		startTime := time.Now().Local()

		// 从 body 中取出数据
		bodyByte, _ := io.ReadAll(r.Body)

		// copy后 还给业务使用
		r.Body = io.NopCloser(bytes.NewBuffer(bodyByte))

		next(w, r)

		endTime := time.Now().Local()
		ms := (endTime.Nanosecond() - startTime.Nanosecond()) / 1000000
		type AccessLog struct {
			Url    string      `json:"url"`
			Method string      `json:"method"`
			Query  string      `json:"query"`
			Body   string      `json:"body"`
			Header http.Header `json:"header"`
			Const  int         `json:"const"`
		}

		l := AccessLog{
			Url:    r.URL.Path,
			Method: r.Method,
			Query:  r.URL.Query().Encode(),
			Body:   string(bodyByte),
			Header: r.Header,
			Const:  ms,
		}
		logx.WithContext(r.Context()).WithFields(logx.Field("request", l)).Info()
	}
}
