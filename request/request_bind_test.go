package request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/goin-go/goin/util"
)

func TestRequest_Bind(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte(`{
		"id":2009020012,
		"strings":["strings-1","strings-2","strings-3","strings-4"],
		"object":{
			"name":"objectName",
			"group":[1,3,5,7,9],
			"created_at":"2019-03-24 13:27:01",
			"children":[{
				"field":"children-1-field"
			},{
				"field":"children-2-field"
			},{
				"field":"children-3-field"
			},{
				"field":"children-4-field"
			}]
		}
	}`)))
	r.Header.Set("Content-Type", util.MIMEJSON)
	req := NewRequest(r)

	type child struct {
		Field string `goin:"field"`
	}
	type object struct {
		Name      string    `goin:"name"`
		Group     []int     `goin:"group"`
		Children  []child   `goin:"children"`
		CreatedAt time.Time `goin:"created_at" time_format:"2006-01-02 15:04:05"`
	}
	type bind struct {
		ID      int64    `goin:"id" default:"0"`
		Name    string   `goin:"name" default:"admin"`
		Strings []string `goin:"strings"`
		Object  object   `goin:"object"`
	}

	data := &bind{}

	err := req.Bind(data)

	if err != nil {
		t.Fatalf("bind error")
	}

	if data.ID != 2009020012 {
		t.Fatalf("bind error")
	}

}
