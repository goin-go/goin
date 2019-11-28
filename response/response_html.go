package response

import (
	"html/template"

	"github.com/goin-go/goin/util"
)

// WriteHTML 输出HTML
func (res *Response) WriteHTML(s string) {
	res.SetContentType(util.MIMEHTML)
	res.Write([]byte(s))
}

// Render 渲染模板文件
func (res *Response) Render(tplPath string, data interface{}) {
	t, err := template.ParseFiles(tplPath)

	if err != nil {
		return
	}

	t.Execute(res.W, data)
}
