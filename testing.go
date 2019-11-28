package goin

import (
	"net/http"
	"net/http/httptest"
)

var g = New(&Options{})

var rt = g.rTree

func cleanRouterTree() {
	rt.nodes = map[string]*Node{}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	g.ServeHTTP(rr, req)

	return rr
}

func createTestContext(url string) *Context {
	req, _ := http.NewRequest("GET", url, nil)
	response := executeRequest(req)

	ctx := &Context{goin: g}
	ctx.reset(response, req)
	return ctx
}

func testRequest(h http.Handler, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	return w
}
