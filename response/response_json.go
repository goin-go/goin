package response

import (
	"encoding/json"

	"github.com/goin-go/goin/util"
)

// WriteJSON 输出JSON
func (res *Response) WriteJSON(j interface{}) {
	b, err := json.Marshal(j)
	if err != nil {
		b = []byte{}
	}
	res.SetContentType(util.MIMEJSON)
	res.Write(b)
}
