package goin

import (
	"testing"
)

func TestRouter(t *testing.T) {
	t.Run("简单路由", func(t *testing.T) {
		t.Run("根节点路由", func(t *testing.T) {
			cleanRouterTree()

			err := rt.Insert("GET", "/", func(ctx *Context) {
				ctx.WriteString("body")
			})
			if err != nil {
				t.Fatal("添加路由返回错误，判定失败")
			}

			node, _, has := rt.Find("GET", "/")

			if !has {
				t.Fatal("未找到路由，判定失败")
			}
			ctx := createTestContext("/")
			node.runLastHandler(ctx)

			if string(ctx.Response.Body()) != "body" {
				t.Fatalf("返回内容不对应，期望：%v，实际：%v", "body", string(ctx.Response.Body()))
			}
		})

		t.Run("添加空路由", func(t *testing.T) {

			t.Run("无路径的空路由", func(t *testing.T) {
				err := rt.Insert("GET", "", func(ctx *Context) {
					ctx.WriteString("/body")
				})
				if err == nil {
					t.Fatal("添加路由返回成功，判定失败")
				}
			})

			t.Run("无handler的空路由", func(t *testing.T) {
				err := rt.Insert("GET", "/")
				if err == nil {
					t.Fatal("添加路由返回成功，判定失败")
				}
			})
		})

		t.Run("添加一级路由", func(t *testing.T) {
			err := rt.Insert("GET", "/body", func(ctx *Context) {
				ctx.WriteString("/body")
			})
			if err != nil {
				t.Fatal("添加路由返回错误，判定失败")
			}

			node, _, has := rt.Find("GET", "/body")

			if !has {
				t.Fatal("未找到路由，判定失败")
			}
			ctx := createTestContext("/body")
			node.runLastHandler(ctx)

			if string(ctx.Response.Body()) != "/body" {
				t.Fatalf("返回内容不对应")
			}
		})

		t.Run("重复添加一级路由", func(t *testing.T) {
			err := rt.Insert("GET", "/body", func(ctx *Context) {
				ctx.WriteString("/body")
			})

			if err == nil {
				t.Fatal("添加路由未返回错误，判定失败")
			}
		})

		t.Run("添加多级路由", func(t *testing.T) {
			cleanRouterTree()
			ctx := createTestContext("/")

			t.Run("添加根路由", func(t *testing.T) {
				err := rt.Insert("GET", "/", func(ctx *Context) {
					ctx.WriteString("/")
				})
				if err != nil {
					t.Fatal("添加路由返回错误，判定失败")
				}
			})

			t.Run("添加一级路由", func(t *testing.T) {
				err := rt.Insert("GET", "/base", func(ctx *Context) {
					ctx.WriteString("/base")
				})
				if err != nil {
					t.Fatal("添加路由返回错误，判定失败")
				}
			})

			t.Run("添加二级路由", func(t *testing.T) {
				err := rt.Insert("GET", "/base/test", func(ctx *Context) {
					ctx.WriteString("/base/test")
				})
				if err != nil {
					t.Fatal("添加路由返回错误，判定失败")
				}
			})

			t.Run("添加三级路由", func(t *testing.T) {
				err := rt.Insert("GET", "/base/test/run", func(ctx *Context) {
					ctx.WriteString("/base/test/run")
				})
				if err != nil {
					t.Fatal("添加路由返回错误，判定失败")
				}
			})

			t.Run("获取根路由", func(t *testing.T) {
				node, _, has := rt.Find("GET", "/")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}
				node.runLastHandler(ctx)

				if string(ctx.Response.Body()) != "/" {
					t.Fatalf("根路由返回内容不对应，期望：%v，实际：%v", "/", string(ctx.Response.Body()))
				}
			})

			t.Run("获取一级路由", func(t *testing.T) {
				node, _, has := rt.Find("GET", "/base")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}
				node.runLastHandler(ctx)

				if string(ctx.Response.Body()) != "/base" {
					t.Fatalf("一级路由返回内容不对应，期望：%v，实际：%v", "/base", string(ctx.Response.Body()))
				}
			})

			t.Run("获取二级路由", func(t *testing.T) {
				node, _, has := rt.Find("GET", "/base/test")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}
				node.runLastHandler(ctx)

				if string(ctx.Response.Body()) != "/base/test" {
					t.Fatalf("二级路由返回内容不对应，期望：%v，实际：%v", "/base/test", string(ctx.Response.Body()))
				}
			})

			t.Run("获取三级路由", func(t *testing.T) {
				node, _, has := rt.Find("GET", "/base/test/run")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}
				node.runLastHandler(ctx)

				if string(ctx.Response.Body()) != "/base/test/run" {
					t.Fatalf("三级路由返回内容不对应，期望：%v，实际：%v", "/base/test/run", string(ctx.Response.Body()))
				}
			})
		})
	})

	t.Run("路由参数", func(t *testing.T) {
		cleanRouterTree()

		t.Run("一个参数", func(t *testing.T) {
			rt.Insert("GET", "/:id", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, params, has := rt.Find("GET", "/admin")

			if !has {
				t.Fatal("未找到路由，判定失败")
				return
			}

			if params["id"] != "admin" {
				t.Fatalf("参数内容不对应，实际返回：%v，判定失败", params["id"])
			}
		})

		t.Run("两个参数", func(t *testing.T) {
			rt.Insert("GET", "/test/:id/:name", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, params, has := rt.Find("GET", "/test/123/admin")

			if !has {
				t.Fatal("未找到路由，判定失败")
			}

			if params["id"] != "123" {
				t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
			}
			if params["name"] != "admin" {
				t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
			}
		})

		t.Run("内置正则表达式参数", func(t *testing.T) {
			rt.Insert("GET", "/test2/:id{int}/:name{string}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, params, has := rt.Find("GET", "/test2/123/admin")

			if !has {
				t.Fatal("未找到路由，判定失败")
			}

			if params["id"] != "123" {
				t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
			}
			if params["name"] != "admin" {
				t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
			}

			_, params, has = rt.Find("GET", "/test2/123a/admin")
			if has {
				t.Fatal("找到路由，判定失败")
			}
		})
		t.Run("内置正则表达式长度参数", func(t *testing.T) {
			t.Run("min max 都存在", func(t *testing.T) {
				cleanRouterTree()
				rt.Insert("GET", "/test2/:id{int[100:200]}/:name{string[2:6]}", func(ctx *Context) {
					ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
				})

				_, params, has := rt.Find("GET", "/test2/123/admin")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}

				if params["id"] != "123" {
					t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
				}
				if params["name"] != "admin" {
					t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
				}

				_, params, has = rt.Find("GET", "/test2/99/a")

				if has {
					t.Fatal("找到路由，判定失败")
				}
			})

			t.Run("只存在min", func(t *testing.T) {
				cleanRouterTree()
				rt.Insert("GET", "/test2/:id{int[100:]}/:name{string[2:]}", func(ctx *Context) {
					ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
				})

				_, params, has := rt.Find("GET", "/test2/123/admin")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}

				if params["id"] != "123" {
					t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
				}
				if params["name"] != "admin" {
					t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
				}

				_, params, has = rt.Find("GET", "/test2/99/a")

				if has {
					t.Fatal("找到路由，判定失败")
				}
				_, params, has = rt.Find("GET", "/test2/9999/a")

				if has {
					t.Fatal("找到路由，判定失败")
				}
			})

			t.Run("只存在max", func(t *testing.T) {
				cleanRouterTree()
				rt.Insert("GET", "/test2/:id{int[:200]}/:name{string[:6]}", func(ctx *Context) {
					ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
				})

				_, params, has := rt.Find("GET", "/test2/123/admin")

				if !has {
					t.Fatal("未找到路由，判定失败")
				}

				if params["id"] != "123" {
					t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
				}
				if params["name"] != "admin" {
					t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
				}

				_, params, has = rt.Find("GET", "/test2/201/a")

				if has {
					t.Fatal("找到路由，判定失败")
				}

				_, params, has = rt.Find("GET", "/test2/200/aaaaaaa")

				if has {
					t.Fatal("找到路由，判定失败")
				}
			})

		})
		t.Run("正则表达式参数", func(t *testing.T) {
			rt.Insert("GET", "/test2/:id{^\\d+$}/:name", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, params, has := rt.Find("GET", "/test2/123/admin")

			if !has {
				t.Fatal("未找到路由，判定失败")
			}

			if params["id"] != "123" {
				t.Fatalf("参数1内容不对应，实际返回：%v，判定失败", params["id"])
			}
			if params["name"] != "admin" {
				t.Fatalf("参数2内容不对应，实际返回：%v，判定失败", params["name"])
			}

			rt.Insert("GET", "/test3/:id{^\\d+$}/:name{^A[\\w]+4$}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, _, has = rt.Find("GET", "/test3/a123/admin")

			if has {
				t.Fatal("不应该找到路由，但实际找到路由，判定失败")
			}

			_, _, has = rt.Find("GET", "/test3/123/admin")

			if has {
				t.Fatal("不应该找到路由，但实际找到路由，判定失败")
			}

			_, _, has = rt.Find("GET", "/test3/123/Admin4")

			if !has {
				t.Fatal("应该找到路由，但实际未找到路由，判定失败")
			}
		})
	})

	t.Run("获取所有handler Name", func(t *testing.T) {

		node, _, has := rt.Find("GET", "/test3/123/Admin4")

		if !has {
			t.Fatal("未找到路由，判定失败")
		}

		names := node.handlerNames()

		if len(names) == 0 {
			t.Fatal("应该返回一个长度大于0的字符串数组，判定失败")
		}
	})

	t.Run("模拟错误", func(t *testing.T) {
		t.Run("模拟无路由参数", func(t *testing.T) {
			cleanRouterTree()
			rt.Insert("GET", "/:id", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			rt.nodes["GET"].next[rune('/')].next[rune(':')].param = nil

			_, _, has := rt.Find("GET", "/123")
			if has {
				t.Fatal("找到路由，判定失败")
			}
		})
		t.Run("模拟正则错误", func(t *testing.T) {
			cleanRouterTree()
			rt.Insert("GET", "/:id{\\d[]+}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			_, _, has := rt.Find("GET", "/123")
			if has {
				t.Fatal("找到路由，判定失败")
			}
		})
	})
}
func BenchmarkRouter(b *testing.B) {
	cleanRouterTree()

	b.Run("根节点路由", func(b *testing.B) {
		rt.Insert("GET", "/", func(ctx *Context) {
			ctx.WriteString("body")
		})
		for i := 0; i < b.N; i++ {
			rt.Find("GET", "/")
		}
	})
	b.Run("多级路由", func(b *testing.B) {
		b.Run("一级路由", func(b *testing.B) {
			rt.Insert("GET", "/base", func(ctx *Context) {
				ctx.WriteString("/base")
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/base")
			}
		})
		b.Run("二级路由", func(b *testing.B) {
			rt.Insert("GET", "/base/test", func(ctx *Context) {
				ctx.WriteString("/base/test")
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/base/test")
			}
		})
		b.Run("三级路由", func(b *testing.B) {
			rt.Insert("GET", "/base/test/run", func(ctx *Context) {
				ctx.WriteString("/base/test/run")
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/base/test/run")
			}
		})
		b.Run("四级路由", func(b *testing.B) {
			rt.Insert("GET", "/base/test/run/info", func(ctx *Context) {
				ctx.WriteString("/base/test/run/info")
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/base/test/run/info")
			}
		})
	})

	b.Run("路由参数", func(b *testing.B) {
		cleanRouterTree()
		b.Run("一个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/test")
			}
		})
		b.Run("两个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id/:name", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/test/admin")
			}
		})
		b.Run("三个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id/:name/:cate", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/test/admin/news")
			}
		})
		b.Run("四个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id/:name/:cate/:in", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/test/admin/news/off")
			}
		})
	})
	b.Run("路由参数使用内置正则表达式", func(b *testing.B) {

		b.Run("一个参数", func(b *testing.B) {
			cleanRouterTree()
			rt.Insert("GET", "/:id{int[100:200]}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123")
			}
		})
		b.Run("两个参数", func(b *testing.B) {
			cleanRouterTree()
			rt.Insert("GET", "/:id{int}/:name{string[2:10]}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123/admin")
			}
		})
		b.Run("三个参数", func(b *testing.B) {
			cleanRouterTree()
			rt.Insert("GET", "/:id{int}/:name{string}/:cate{string}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123/admin/news")
			}
		})
	})
	b.Run("路由参数使用正则表达式", func(b *testing.B) {
		cleanRouterTree()
		b.Run("一个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id{\\d+}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})

			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123")
			}
		})
		b.Run("两个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id{\\d+}/:name{[\\w]+}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123/admin")
			}
		})
		b.Run("三个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id{\\d+}/:name{[\\w]+}/:cate{[\\w]+}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123/admin/news")
			}
		})
		b.Run("四个参数", func(b *testing.B) {
			rt.Insert("GET", "/:id{\\d+}/:name{[\\w]+}/:cate{[\\w]+}/:in{[true|false]}", func(ctx *Context) {
				ctx.WriteString(ctx.URLParams.DefaultString("id", ""))
			})
			for i := 0; i < b.N; i++ {
				rt.Find("GET", "/123/admin/news/false")
			}
		})
	})
}
