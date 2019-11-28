package util

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMap2Struct(t *testing.T) {

	jsonStr := `
{
	"id":123,
	"name":     "admin",
	"isMember": true,
	"children": {
		"cname": "children-name"
	},
	"list":    [1, 2, 3],
	"strings": ["aa", "bb", "cc"],
	"childrenList": [{
		"cname": "children-name-1"
	}, {
		"cname": "children-name-2"
	}],
	"time":  "2006-01-02 15:04:05",
	"times": ["2006-01-02 15:04:05", "2006-01-03 15:04:05"]
}
`
	mapData := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		fmt.Println(err.Error())
	}
	type Children struct {
		CName string `goin:"cname"`
	}
	type StructData struct {
		ID           int         `goin:"id"`
		Name         string      `goin:"name"`
		IsMember     bool        `goin:"isMember"`
		Has          bool        `goin:"has" default:"true"`
		Children     Children    `goin:"children"`
		List         []int       `goin:"list"`
		Strings      []string    `goin:"strings"`
		ChildrenList []Children  `goin:"childrenList"`
		Time         time.Time   `goin:"time" time_format:"2006-01-02 15:04:05"`
		Times        []time.Time `goin:"times" time_format:"2006-01-02 15:04:05"`
	}

	var data = &StructData{}
	Map2StructWithOption(mapData, data, &Map2StructOption{
		KeyNameTagName: "goin",
	})

	if !data.IsMember {
		t.Error("失败")
	}
}
