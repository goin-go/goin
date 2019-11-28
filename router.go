package goin

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/goin-go/goin/util"
)

type (
	// Router Router
	Router struct {
		goin  *Goin
		nodes map[string]*Node
	}

	// Node Node结构
	Node struct {
		next  map[rune]*Node
		param *Params

		fullPath string
		handlers []HandlerFunc
	}

	// Params 参数结构
	Params struct {
		name          string
		isRegexpAlias bool
		regexpStr     string
		min           string
		max           string
	}
)

// newRouter 新建路由树
func newRouter(g *Goin) *Router {
	return &Router{
		goin:  g,
		nodes: map[string]*Node{},
	}
}

// Insert 添加一个路由
func (tree *Router) Insert(method string, path string, handles ...HandlerFunc) error {

	if len(path) == 0 {
		return errors.New("路由路径不能为空")
	}
	if len(handles) == 0 {
		return errors.New("路由行为不能为空")
	}

	// 获取到请求类型对应的前缀树
	root, ok := tree.nodes[method]

	// 没有获取到就新建
	if !ok {
		root = new(Node)
		root.next = make(map[rune]*Node, 0)
		tree.nodes[method] = root
	}

	handles = append(tree.goin.middlewares, handles...)
	return root.Insert(path, handles)
}

// Find 查找路由
func (tree *Router) Find(method string, path string) (*Node, map[string]string, bool) {
	// 获取到请求类型对应的前缀树
	root, ok := tree.nodes[method]

	if !ok {
		return nil, nil, false
	}

	node, params, has := root.Search(path)
	if !has {
		return nil, nil, false
	}
	return node, params, true
}

// Insert 添加节点
func (root *Node) Insert(path string, handlers []HandlerFunc) error {
	fullPath := path

	params, err := getPathParams(path)
	if err != nil {
		return err
	}

	paramsIndex := 0

	path = removePathParams(path)
	for _, c := range path {
		if root.next[c] == nil {
			node := new(Node)
			node.next = make(map[rune]*Node, 0)

			root.next[c] = node
		}

		if c == ':' {
			oldParam := root.next[c].param
			newParam := params[paramsIndex]
			if oldParam == nil {
				root.next[c].param = newParam
			} else if oldParam.name != newParam.name || oldParam.isRegexpAlias != newParam.isRegexpAlias || oldParam.regexpStr != newParam.regexpStr {
				return fmt.Errorf("路由：%v 与其他路由存在参数 %v ≈ %v 冲突", fullPath, root.next[c].param.name, params[paramsIndex].name)
			}
			paramsIndex++
		}

		root = root.next[c]
	}

	if root.fullPath != "" {
		return fmt.Errorf("存在冲突路由：%v ≈ %v", root.fullPath, fullPath)
	}
	root.fullPath = fullPath
	root.handlers = handlers
	return nil
}

// Search 查找节点
func (root *Node) Search(path string) (*Node, map[string]string, bool) {
	params := make(map[string]string)

	for i := 0; i < len(path); i++ {
		c := path[i]
		if root.next[rune(c)] == nil && len(root.next) == 1 && root.next[':'] != nil {
			c = ':'
			if param := root.next[':'].param; param != nil {

				ni := strings.Index(path[i:], "/")
				if ni == -1 {
					ni = len(path)
				} else {
					ni = i + ni
				}
				paramVal := path[i:ni]

				/*
					参数类型：
					int:纯数字
					string:字符串
					bool:布尔
				*/
				if param.isRegexpAlias {
					switch param.regexpStr {
					case "int":
						intVal, err := strconv.ParseInt(paramVal, 10, 0)
						if err != nil {
							return nil, nil, false
						}
						if param.min != "*" {
							min, _ := strconv.ParseInt(param.min, 0, 64)
							if intVal < min {
								return nil, nil, false
							}
						}
						if param.max != "*" {
							max, _ := strconv.ParseInt(param.max, 0, 64)
							if intVal > max {
								return nil, nil, false
							}
						}

					case "string":
						l := int64(len(paramVal))
						if param.min != "*" {
							min, _ := strconv.ParseInt(param.min, 0, 64)
							if l < min {
								return nil, nil, false
							}
						}
						if param.max != "*" {
							max, _ := strconv.ParseInt(param.max, 0, 64)
							if l > max {
								return nil, nil, false
							}
						}
					}
				} else {
					re := regexp.MustCompile(param.regexpStr)
					if !re.MatchString(paramVal) {
						return nil, nil, false
					}
				}

				params[param.name] = paramVal
				i = ni - 1

			} else {
				return nil, nil, false
			}
		} else if root.next[rune(c)] == nil {
			return nil, nil, false
		}

		root = root.next[rune(c)]
	}
	return root, params, true
}

// runLastHandler 运行最后一个Handler
func (root *Node) runLastHandler(ctx *Context) {
	if len(root.handlers) > 0 {
		root.handlers[len(root.handlers)-1](ctx)
	}
}

// handlerNames 获取所有handler的名称
func (root *Node) handlerNames() (names []string) {
	if len(root.handlers) > 0 {
		lastHandler := root.handlers[len(root.handlers)-1]
		names = append(names, runtime.FuncForPC(reflect.ValueOf(lastHandler).Pointer()).Name())
	}
	return names
}

// getPathParams 获取路径中所有的参数
func getPathParams(path string) ([]*Params, error) {
	params := make([]*Params, 0)
	start := -1
	end := -1
	for i := 0; i < len(path); i++ {
		if path[i] == ':' && start == -1 {
			start = i + 1
			continue
		}

		if start >= 1 {
			if path[i] == '/' {
				end = i
			}
			if i == len(path)-1 {
				end = i + 1
			}
		} else {
			start = -1
			end = -1
		}

		if start > -1 && end > start {
			s := path[start:end]
			rp := "string"

			// 参数包含正则规则
			regStart := strings.Index(s, "{")
			if regStart > 0 && util.LastChar(s) == '}' {
				rp = s[(regStart + 1):(len(s) - 1)]

				s = s[:regStart]
			}
			start = -1
			end = -1
			param := &Params{
				name:      s,
				regexpStr: rp,
				min:       "*",
				max:       "*",
			}

			if strings.HasPrefix(rp, "int") || strings.HasPrefix(rp, "string") {
				param.isRegexpAlias = true

				// 判断是否存在[min:max]格式
				mmStart := strings.Index(rp, "[")
				split := strings.Index(rp, ":")
				if mmStart > 0 && split > mmStart && util.LastChar(rp) == ']' {
					mm := rp[(mmStart + 1):(len(rp) - 1)]
					mmv := strings.Split(mm, ":")
					if len(mmv) == 2 {
						var err error
						if mmv[0] == "" {
							mmv[0] = "*"
						} else {
							_, err = strconv.ParseInt(mmv[0], 10, 64)
						}

						if mmv[1] == "" {
							mmv[1] = "*"
						} else {
							_, err = strconv.ParseInt(mmv[1], 10, 64)
						}
						if err == nil {
							param.regexpStr = rp[:mmStart]
							param.min = mmv[0]
							param.max = mmv[1]
						}

					}

				}

			} else {
				_, err := regexp.Compile(rp)
				if err != nil {
					return nil, errors.New("正则表达式错误:" + rp)
				}
			}

			params = append(params, param)
		}
	}
	return params, nil
}

// removePathParams 移除路径中的参数
// 把路径中的参数名称给移除，只保留:
func removePathParams(path string) string {
	for i := 0; i < len(path); i++ {
		if path[i] == ':' {
			ns := strings.Index(path[i:], "/")
			if ns == -1 {
				ns = len(path) - i
			}
			path = path[0:i+1] + path[ns+i:]
		}
	}
	return path
}
