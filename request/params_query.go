package request

import (
	"net/url"
	"strconv"
)

// QueryParams Query参数结构
type QueryParams struct {
	query url.Values
}

// String 获取字符串值
func (p *QueryParams) String(key string) (string, bool) {
	if s, has := p.query[key]; has && len(s) > 0 {
		return s[0], true
	}
	return "", false
}

// Int 获取整形值
func (p *QueryParams) Int(key string) (int64, bool) {
	if s, has := p.query[key]; has && len(s) > 0 {
		i, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return 0, false
		}
		return i, true
	}
	return 0, false
}

// Uint 获取无符号整形值
func (p *QueryParams) Uint(key string) (uint64, bool) {
	if s, has := p.query[key]; has && len(s) > 0 {
		u, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			return 0, false
		}
		return u, true
	}
	return 0, false
}

// Float 获取浮点值
func (p *QueryParams) Float(key string) (float64, bool) {
	if s, has := p.query[key]; has && len(s) > 0 {
		f, err := strconv.ParseFloat(s[0], 64)
		if err != nil {
			return 0, false
		}
		return f, true
	}
	return 0, false
}

// Bool 获取布尔值
func (p *QueryParams) Bool(key string) (bool, bool) {
	if s, has := p.query[key]; has && len(s) > 0 {
		b, err := strconv.ParseBool(s[0])
		if err != nil {
			return false, false
		}
		return b, true
	}
	return false, false
}

// DefaultString 获取字符串值，如果没有值则使用默认值
func (p *QueryParams) DefaultString(key string, def string) string {
	if s, has := p.query[key]; has && len(s) > 0 {
		return s[0]
	}
	return def
}

// DefaultInt 获取整形值，如果没有值则使用默认值
func (p *QueryParams) DefaultInt(key string, def int64) int64 {
	if s, has := p.query[key]; has && len(s) > 0 {
		i, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return def
		}
		return i
	}
	return def
}

// DefaultUint 获取无符号整形值，如果没有值则使用默认值
func (p *QueryParams) DefaultUint(key string, def uint64) uint64 {
	if s, has := p.query[key]; has && len(s) > 0 {
		u, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			return def
		}
		return u
	}
	return def
}

// DefaultFloat 获取浮点值，如果没有值则使用默认值
func (p *QueryParams) DefaultFloat(key string, def float64) float64 {
	if s, has := p.query[key]; has && len(s) > 0 {
		f, err := strconv.ParseFloat(s[0], 64)
		if err != nil {
			return def
		}
		return f
	}
	return def
}

// DefaultBool 获取布尔值，如果没有值则使用默认值
func (p *QueryParams) DefaultBool(key string, def bool) bool {
	if s, has := p.query[key]; has && len(s) > 0 {
		b, err := strconv.ParseBool(s[0])
		if err != nil {
			return def
		}
		return b
	}
	return def
}

// Array 获取数组
func (p *QueryParams) Array(key string) ([]string, bool) {
	s, has := p.query[key]
	if !has {
		return nil, false
	}
	return s, true
}
