package goin

import "testing"

func TestRouterGroup(t *testing.T) {
	t.Run("简单组", func(t *testing.T) {
		rg := g.Group("/test")

		rg.GET("/name", func(ctx *Context) {
			ctx.WriteString("body")
		})

		_, _, has := rt.Find("GET", "/test/name")

		if !has {
			t.Fatal("未找到路由，判定失败")
		}
	})

	t.Run("组嵌套", func(t *testing.T) {
		rg1 := g.Group("/group1")
		rg11 := rg1.Group("/group11")

		rg1.GET("/name", func(ctx *Context) {
			ctx.WriteString("body")
		})

		rg11.GET("/name", func(ctx *Context) {
			ctx.WriteString("body")
		})

		_, _, has := rt.Find("GET", "/group1/name")

		if !has {
			t.Fatal("未找到路由，判定失败")
		}

		_, _, has = rt.Find("GET", "/group1/group11/name")

		if !has {
			t.Fatal("未找到路由，判定失败")
		}
	})
}
