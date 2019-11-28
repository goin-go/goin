package goin

import (
	"net"
	"net/http"
	"strings"

	"github.com/goin-go/goin/request"
	"github.com/goin-go/goin/response"
)

// Context Context结构
type Context struct {
	*request.Request
	*response.Response

	index    int
	goin     *Goin
	handlers []HandlerFunc
	Error    error
	kv       map[string]interface{}
	fullPath string
}

// reset 重置Context
func (c *Context) reset(w http.ResponseWriter, req *http.Request) {
	c.index = -1
	c.handlers = []HandlerFunc{}
	c.Request = request.NewRequest(req)
	c.Response = response.NewResponse(w)
	c.kv = make(map[string]interface{})
	c.fullPath = ""
	c.Error = nil
}

// setURLParams 设置URL参数
func (c *Context) setURLParams(params map[string]string) {
	c.Request.SetURLParams(params)
}

// setHandlers 设置handlers
func (c *Context) setHandlers(handlers []HandlerFunc) {
	c.handlers = handlers
}

// Next 执行下一个Handler
func (c *Context) Next() {
	c.index = c.index + 1
	l := len(c.handlers)
	if c.index < l && !c.IsEnd {
		c.handlers[c.index](c)
	}
}

// SetKV 设置kv
func (c *Context) SetKV(key string, value interface{}) {
	c.kv[key] = value
}

// KV 获取kv
func (c *Context) KV(key string) interface{} {
	return c.kv[key]
}

// ClientIP 获取客户端IP
func (c *Context) ClientIP() string {
	clientIP := c.Headers().Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(c.Headers().Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// FullPath 获取请求完整路径
func (c *Context) FullPath() string {
	return c.fullPath
}

// Method 获取请求类型
func (c *Context) Method() string {
	return c.Req.Method
}
