package request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goin-go/goin/util"
)

func TestRequest_SetURLParams(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte("{\"name\":\"test\",\"pwd\":123,\"children\":[1,2]}")))
	r.Header.Set("Content-Type", util.MIMEJSON)
	req := NewRequest(r)

	req.SetURLParams(map[string]string{"name": "admin"})

	if req.URLParams.DefaultString("name", "error") == "error" {
		t.Fatal("获取url参数错误")
	}
}
