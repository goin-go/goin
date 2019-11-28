package request

import (
	"net/http"
)

// Request Request结构
type Request struct {
	Req         *http.Request
	URLParams   *URLParams
	QueryParams *QueryParams
	FormParams  *FormParams
	ContentType string
	Cookie      *Cookie
}

// NewRequest 创建新Request对象
func NewRequest(req *http.Request) *Request {
	request := &Request{Req: req}
	request.initAll()
	return request
}

// initAll 初始化
func (req *Request) initAll() {
	// 请求的Content-Type
	req.ContentType = filterFlags(req.Req.Header.Get("Content-Type"))

	// 初始化
	req.URLParams = &URLParams{kv: make(map[string]string)}
	req.QueryParams = &QueryParams{query: req.Req.URL.Query()}
	req.FormParams = &FormParams{}
	req.Cookie = &Cookie{r: req}

	req.FormParams.initForm(req)
}

// SetURLParams 设置url参数
func (req *Request) SetURLParams(params map[string]string) {
	req.URLParams.kv = params
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

// Headers 获取请求Headers
func (req *Request) Headers() http.Header {
	return req.Req.Header
}
