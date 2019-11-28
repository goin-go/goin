package response

import (
	"github.com/goin-go/goin/util"
)

// WriteString 输出字符串
func (res *Response) WriteString(s string) {
	res.SetContentType(util.MIMEPlain)
	res.Write([]byte(s))
}

// AppendWriteString 追加字符串
func (res *Response) AppendWriteString(s string) {
	res.AppendWrite([]byte(s)...)
}
