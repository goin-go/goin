package request

import "net/http"

// Cookie Cookie结构
type Cookie struct {
	r *Request
}

// String 获取字符串值
func (cookie *Cookie) String(key string) (string, bool) {
	c, err := cookie.r.Req.Cookie(key)

	if err != nil {
		return "", false
	}

	return c.Value, true
}

// DefaultString 获取字符串值，如果没有值则使用默认值
func (cookie *Cookie) DefaultString(key string, def string) string {
	c, err := cookie.r.Req.Cookie(key)

	if err != nil {
		return def
	}

	return c.Value
}

// Cookie 获取Cookie
func (cookie *Cookie) Cookie(key string) (*http.Cookie, bool) {
	c, err := cookie.r.Req.Cookie(key)

	if err != nil {
		return nil, false
	}
	return c, true
}

// AddCookie 添加Cookie
func (cookie *Cookie) AddCookie(c *http.Cookie) {
	cookie.r.Req.AddCookie(c)
}
