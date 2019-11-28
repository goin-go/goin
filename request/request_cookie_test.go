package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCookie(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/query?name=test&pwd=123&children=1&children=2&c", nil)
	req := NewRequest(r)
	c := &Cookie{
		r: req,
	}

	c.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "test",
	})

	c.AddCookie(&http.Cookie{
		Name:  "test2",
		Value: "test2",
	})
	t.Run("String", func(t *testing.T) {
		v, has := c.String("test")

		if !has {
			t.Fatal("未获取到Cookie值")
		}

		if v != "test" {
			t.Fatal("获取到Cookie值不正确")
		}

		v, has = c.String("test3")

		if has {
			t.Fatal("不该获取到Cookie值")
		}
	})

	t.Run("DefaultString", func(t *testing.T) {
		v := c.DefaultString("test", "")
		if v != "test" {
			t.Fatal("获取到Cookie值不正确")
		}

		v = c.DefaultString("test3", "123")

		if v != "123" {
			t.Fatal("获取到Cookie值不正确")
		}
	})
}
