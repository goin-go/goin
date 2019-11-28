package request

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/goin-go/goin/util"
)

// FormParams FormParams结构
type FormParams struct {
	form   map[string]interface{}
	file   map[string][]*multipart.FileHeader
	isJSON bool
}

// initForm 初始化
func (p *FormParams) initForm(req *Request) {
	p.form = make(map[string]interface{}, 0)
	p.isJSON = false

	switch req.ContentType {
	case util.MIMEMultipartPOSTForm:
		err := req.Req.ParseMultipartForm(32 << 20)
		if err != nil {
			return
		}

		for key, vals := range req.Req.MultipartForm.Value {
			p.form[key] = vals
		}
		p.file = req.Req.MultipartForm.File
		break
	case util.MIMEPOSTForm:
		err := req.Req.ParseForm()
		if err != nil {
			return
		}

		for key, vals := range req.Req.PostForm {
			p.form[key] = vals
		}

		break
	case util.MIMEJSON:
		p.isJSON = true
		body, err := ioutil.ReadAll(req.Req.Body)
		if err != nil {
			return
		}
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			return
		}
		p.form = jsonMap
		break
	}
}

// String 获取字符串值
func (p *FormParams) String(key string) (string, bool) {
	if val, has := p.form[key]; has {
		if p.isJSON {
			return util.InterfaceToString(val)
		}
		return val.([]string)[0], true
	}
	return "", false
}

// Int 获取整形值
func (p *FormParams) Int(key string) (int64, bool) {
	if val, has := p.form[key]; has {

		if p.isJSON {
			return util.InterfaceToInt(val, 64)
		}

		i, err := strconv.ParseInt(val.([]string)[0], 10, 64)
		if err != nil {
			return 0, false
		}
		return i, true
	}
	return 0, false
}

// Uint 获取无符号整形值
func (p *FormParams) Uint(key string) (uint64, bool) {
	if val, has := p.form[key]; has {
		if p.isJSON {
			return util.InterfaceToUint(val, 64)
		}
		u, err := strconv.ParseUint(val.([]string)[0], 10, 64)
		if err != nil {
			return 0, false
		}
		return u, true
	}
	return 0, false
}

// Float 获取浮点数值
func (p *FormParams) Float(key string) (float64, bool) {
	if val, has := p.form[key]; has {
		if p.isJSON {
			return util.InterfaceToFloat(val, 64)
		}

		f, err := strconv.ParseFloat(val.([]string)[0], 64)
		if err != nil {
			return 0, false
		}

		return f, true
	}
	return 0, false
}

// Bool 获取布尔值
func (p *FormParams) Bool(key string) (bool, bool) {
	if val, has := p.form[key]; has {
		if p.isJSON {
			return util.InterfaceToBool(val)
		}

		b, err := strconv.ParseBool(val.([]string)[0])
		if err != nil {
			return false, false
		}

		return b, true
	}
	return false, false
}

// DefaultString 获取字符串值，如果没有值则使用默认值
func (p *FormParams) DefaultString(key string, def string) string {
	s, ok := p.String(key)
	if ok {
		return s
	}
	return def
}

// DefaultInt 获取整形值，如果没有值则使用默认值
func (p *FormParams) DefaultInt(key string, def int64) int64 {
	i, ok := p.Int(key)
	if ok {
		return i
	}
	return def
}

// DefaultUint 获取无符号整形值，如果没有值则使用默认值
func (p *FormParams) DefaultUint(key string, def uint64) uint64 {
	u, ok := p.Uint(key)
	if ok {
		return u
	}
	return def
}

// DefaultFloat 获取浮点值，如果没有值则使用默认值
func (p *FormParams) DefaultFloat(key string, def float64) float64 {
	f, ok := p.Float(key)
	if ok {
		return f
	}
	return def
}

// DefaultBool 获取布尔值，如果没有值则使用默认值
func (p *FormParams) DefaultBool(key string, def bool) bool {
	b, ok := p.Bool(key)
	if ok {
		return b
	}
	return def
}

// Array 获取数组
func (p *FormParams) Array(key string) ([]string, bool) {
	if val, has := p.form[key]; has {
		if p.isJSON {
			return util.InterfaceToStrings(val)
		}
		ss, ok := val.([]string)
		return ss, ok
	}
	return nil, false
}

// File File结构
type File struct {
	Filename  string
	Size      int64
	FileBytes []byte
	Type      string
	Endings   []string
}

// File 获取上传的文件
func (p *FormParams) File(key string) (*File, bool) {
	if p.file == nil {
		return nil, false
	}
	headers, ok := p.file[key]
	if !ok || len(headers) == 0 {
		return nil, false
	}
	file, err := headers[0].Open()
	defer file.Close()

	if err != nil {
		return nil, false
	}

	fileBytes, err := ioutil.ReadAll(file)
	fileType := http.DetectContentType(fileBytes)
	fileEndings, err := mime.ExtensionsByType(fileType)

	return &File{
		Filename:  headers[0].Filename,
		Size:      headers[0].Size,
		FileBytes: fileBytes,
		Type:      fileType,
		Endings:   fileEndings,
	}, true
}

// FileHeader 获取FileHeader
func (p *FormParams) FileHeader(key string) (*multipart.FileHeader, bool) {
	if p.file == nil {
		return nil, false
	}
	headers, ok := p.file[key]
	if !ok || len(headers) == 0 {
		return nil, false
	}

	return headers[0], true
}

// SaveFile 保存上传的文件
func (p *FormParams) SaveFile(key string, path string) bool {
	newFile, err := os.Create(path)
	if err != nil {
		return false
	}
	defer newFile.Close()
	file, ok := p.File(key)
	if !ok {
		return false
	}
	if _, err := newFile.Write(file.FileBytes); err != nil {
		return false
	}
	return true
}
