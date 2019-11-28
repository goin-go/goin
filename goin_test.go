package goin

import (
	"testing"
)

func TestGoin(t *testing.T) {

	t.Run("新建对象", func(t *testing.T) {
		var g = New(&Options{
			BasePath: "/api",
		})

		if g.basePath != "/api" {
			t.Fatal("basepath error")
		}
	})

	t.Run("添加路由", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {

		})
	})
}
