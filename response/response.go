package response

import (
	"net/http"
)

// Response Response结构
type Response struct {
	W           http.ResponseWriter
	status      int
	body        []byte
	contentType string
	IsEnd       bool
}

// NewResponse 新建Response
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{W: w, IsEnd: false}
}

// SetStatus 设置status
func (res *Response) SetStatus(status int) {
	if !res.IsEnd {
		res.status = status
		res.W.WriteHeader(res.status)
	}
}

// SetHeader 设置Header
func (res *Response) SetHeader(key, value string) {
	if !res.IsEnd {
		res.W.Header().Set(key, value)
	}
}

// SetContentType 设置ContentType
func (res *Response) SetContentType(ct string) {
	if !res.IsEnd {
		res.contentType = ct
	}
}

// End 结束
// 结束响应，把结果返回给客户端
// 如果提前使用End，剩余的Handler将不会运行，已运行的Hanlder依次返回
func (res *Response) End() {
	if !res.IsEnd {
		res.W.Header().Set("Content-Type", res.contentType)
		res.W.Write(res.body)
		res.IsEnd = true
	}
}

// Write 输出
func (res *Response) Write(b []byte) {
	if !res.IsEnd {
		res.body = b
	}
}

// AppendWrite 追加输出
func (res *Response) AppendWrite(b ...byte) {
	if !res.IsEnd {
		res.body = append(res.body, b...)
	}
}

// Body 获取Body
func (res *Response) Body() []byte {
	return res.body
}
