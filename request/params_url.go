package request

import (
	"strconv"
)

// URLParams URLParams结构
type URLParams struct {
	kv map[string]string
}

// String 获取字符串值
func (p *URLParams) String(key string) (string, bool) {
	if s, has := p.kv[key]; has {
		return s, true
	}
	return "", false
}

// Int 获取整形值
func (p *URLParams) Int(key string) (int64, bool) {
	if s, has := p.kv[key]; has {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, false
		}
		return i, true
	}
	return 0, false
}

// Uint 获取无符号整形值
func (p *URLParams) Uint(key string) (uint64, bool) {
	if s, has := p.kv[key]; has {
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return 0, false
		}
		return u, true
	}
	return 0, false
}

// Float 获取浮点值
func (p *URLParams) Float(key string) (float64, bool) {
	if s, has := p.kv[key]; has {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, false
		}
		return f, true
	}
	return 0, false
}

// Bool 获取布尔值
func (p *URLParams) Bool(key string) (bool, bool) {
	if s, has := p.kv[key]; has {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return false, false
		}
		return b, true
	}
	return false, false
}

// DefaultString 获取字符串值，如果没有值则使用默认值
func (p *URLParams) DefaultString(key string, def string) string {
	if s, has := p.kv[key]; has {
		return s
	}
	return def
}

// DefaultInt 获取整形值，如果没有值则使用默认值
func (p *URLParams) DefaultInt(key string, def int64) int64 {
	if s, has := p.kv[key]; has {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return def
		}
		return i
	}
	return def
}

// DefaultUint 获取无符号整形值，如果没有值则使用默认值
func (p *URLParams) DefaultUint(key string, def uint64) uint64 {
	if s, has := p.kv[key]; has {
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return def
		}
		return u
	}
	return def
}

// DefaultFloat 获取浮点值，如果没有值则使用默认值
func (p *URLParams) DefaultFloat(key string, def float64) float64 {
	if s, has := p.kv[key]; has {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return def
		}
		return f
	}
	return def
}

// DefaultBool 获取布尔值，如果没有值则使用默认值
func (p *URLParams) DefaultBool(key string, def bool) bool {
	if s, has := p.kv[key]; has {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return def
		}
		return b
	}
	return def
}
