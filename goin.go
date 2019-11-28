package goin

import (
	"fmt"
	"net/http"
	"sync"
)

type (
	// HandlerFunc HandlerFunc
	HandlerFunc func(ctx *Context)

	// Goin Goin结构
	Goin struct {
		rTree       *Router
		middlewares []HandlerFunc
		router404   HandlerFunc
		router500   HandlerFunc
		pool        sync.Pool
		*RouterGroup
	}
)

var (
	// defaultRouter404Handler 默认404
	defaultRouter404Handler HandlerFunc = func(ctx *Context) {
		ctx.SetStatus(404)
		ctx.WriteString("404")
		ctx.End()
	}
	// defaultRouter500Handler 默认500
	defaultRouter500Handler HandlerFunc = func(ctx *Context) {
		ctx.SetStatus(500)
		ctx.WriteString(ctx.Error.Error())
		ctx.End()
	}
)

// Options Goin对象创建时的选项
type Options struct {
	BasePath  string
	Router404 HandlerFunc
	Router500 HandlerFunc
}

// New 创建Goin对象
func New(options *Options) *Goin {
	goin := &Goin{}
	goin.rTree = newRouter(goin)

	if options.BasePath == "" {
		options.BasePath = "/"
	}
	goin.RouterGroup = &RouterGroup{
		basePath:    options.BasePath,
		goin:        goin,
		isRoot:      true,
		middlewares: make([]HandlerFunc, 0),
	}

	// 404
	if options.Router404 != nil {
		goin.router404 = func(ctx *Context) {
			options.Router404(ctx)
			ctx.End()
		}
	} else {
		goin.router404 = defaultRouter404Handler
	}

	// 500
	if options.Router500 != nil {
		goin.router500 = func(ctx *Context) {
			options.Router500(ctx)
			ctx.End()
		}
	} else {
		goin.router500 = defaultRouter500Handler
	}

	goin.pool.New = func() interface{} {
		return &Context{goin: goin}
	}

	return goin
}

// addRouter 添加路由
func (g *Goin) addRouter(method string, path string, handles ...HandlerFunc) error {
	return g.rTree.Insert(method, path, handles...)
}

// Group 创建路由组
func (g *Goin) Group(groupPath string, handles ...HandlerFunc) *RouterGroup {
	return g.RouterGroup.Group(groupPath, handles...)
}

// ServeHTTP
func (g *Goin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := g.pool.Get().(*Context)
	ctx.reset(w, req)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case string:
				ctx.Error = fmt.Errorf(err.(string))
			case error:
				ctx.Error = err.(error)
			}

			g.router500(ctx)
		}
	}()

	node, params, ok := g.rTree.Find(req.Method, req.URL.Path)

	if !ok || len(node.handlers) == 0 {
		g.router404(ctx)
	} else {
		ctx.setURLParams(params)
		ctx.setHandlers(node.handlers)
		ctx.fullPath = node.fullPath
	}
	ctx.Next()

	// 最终输出
	ctx.End()
}

// Use 加载中间件
func (g *Goin) Use(handler ...HandlerFunc) {
	if len(handler) > 0 {
		g.middlewares = append(g.middlewares, handler...)
	}
}
