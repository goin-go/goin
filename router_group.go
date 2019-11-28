package goin

import (
	"net/http"

	"github.com/goin-go/goin/util"
)

// RouterGroup RouterGroup结构
type RouterGroup struct {
	basePath    string
	goin        *Goin
	isRoot      bool
	middlewares []HandlerFunc
}

// GET 新增GET路由
func (rg *RouterGroup) GET(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodGet, path, handles...)
}

// POST 新增POST路由
func (rg *RouterGroup) POST(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodPost, path, handles...)
}

// DELETE 新增DELETE路由
func (rg *RouterGroup) DELETE(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodDelete, path, handles...)
}

// PUT 新增PUT路由
func (rg *RouterGroup) PUT(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodPut, path, handles...)
}

// PATCH 新增PATCH路由
func (rg *RouterGroup) PATCH(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodPatch, path, handles...)
}

// OPTIONS 新增OPTIONS路由
func (rg *RouterGroup) OPTIONS(path string, handles ...HandlerFunc) error {
	return rg.AddRouter(http.MethodOptions, path, handles...)
}

// AddRouter 新增路由
func (rg *RouterGroup) AddRouter(method string, path string, handles ...HandlerFunc) error {
	handles = append(rg.middlewares, handles...)
	return rg.goin.addRouter(method, rg.calculateAbsolutePath(path), handles...)
}

// Group 创建路由组
func (rg *RouterGroup) Group(groupPath string, handles ...HandlerFunc) *RouterGroup {
	basePath := rg.calculateAbsolutePath(groupPath)
	crg := &RouterGroup{
		basePath:    basePath,
		goin:        rg.goin,
		isRoot:      false,
		middlewares: handles,
	}

	return crg
}

// Use 加载中间件
func (rg *RouterGroup) Use(handler ...HandlerFunc) {
	if len(handler) > 0 {
		rg.middlewares = append(rg.middlewares, handler...)
	}
}

// calculateAbsolutePath 生成绝对路径
func (rg *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return util.JoinPaths(rg.basePath, relativePath)
}
