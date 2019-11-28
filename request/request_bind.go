package request

import (
	"github.com/goin-go/goin/util"
)

// Bind 绑定请求body数据到结构
// 绑定方法使用了反射，性能有影响
func (req *Request) Bind(bindStruct interface{}) error {
	var err error
	var mapData = req.FormParams.form

	err = util.Map2StructWithOption(mapData, bindStruct, &util.Map2StructOption{
		KeyNameTagName:      "goin",
		DefaultTagName:      "default",
		TimeFormatTagName:   "time_format",
		TimeLocationTagName: "time_location",
	})

	if err != nil {
		return err
	}

	return nil
}
