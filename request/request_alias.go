package request

// GET 请求query参数的别名快捷方法
func (req *Request) GET(key string) (string, bool) {
	val, has := req.QueryParams.String(key)
	if has == false {
		return "", false
	}

	return val, true
}

// DefaultGET 请求query参数的别名快捷方法，带默认值
func (req *Request) DefaultGET(key string, def string) string {
	return req.QueryParams.DefaultString(key, def)
}

// POST 请求body参数的别名快捷方法
func (req *Request) POST(key string) (string, bool) {
	return req.FormParams.String(key)
}

// DefaultPOST 请求body参数的别名快捷方法，带默认值
func (req *Request) DefaultPOST(key string, def string) string {
	return req.FormParams.DefaultString(key, def)
}
