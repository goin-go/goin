package response

import (
	"net/http"
)

// SetCookie 设置Cookie
func (res *Response) SetCookie(cookie *http.Cookie) {
	http.SetCookie(res.W, cookie)
}
