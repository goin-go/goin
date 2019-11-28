package request

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/goin-go/goin/util"
)

func TestFormParams(t *testing.T) {
	t.Run("普通表单", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte("name=test&pwd=123&children=1&children=2&c")))
		r.Header.Set("Content-Type", util.MIMEPOSTForm)
		req := NewRequest(r)

		f := req.FormParams

		_, ok := f.String("empty")

		if ok {
			t.Fatalf("不应该获取到empty值")
		}

		name, ok := f.String("name")

		if !ok {
			t.Fatalf("未获取到name值")
		}
		if name != "test" {
			t.Fatalf("获取到的name值不正确")
		}

		pwd, ok := f.Int("pwd")

		if !ok {
			t.Fatalf("未获取到pwd值")
		}
		if pwd != 123 {
			t.Fatalf("获取到的pwd值不正确")
		}

		children, ok := f.Array("children")
		if !ok {
			t.Fatalf("未获取到children值")
		}

		if !reflect.DeepEqual(children, []string{"1", "2"}) {
			t.Fatalf("获取到的children值不正确")
		}
	})

	t.Run("带上传文件表单", func(t *testing.T) {
		path := "params_form.go" //要上传文件所在路径
		file, err := os.Open(path)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		// body := bytes.NewBuffer([]byte("name=test&pwd=123&children=1&children=2&c"))
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("my_file", filepath.Base(path))
		writer.WriteField("name", "test")
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		io.Copy(part, file)
		writer.Close()
		r := httptest.NewRequest("POST", "/upload", body)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		req := NewRequest(r)

		f := req.FormParams

		myFile, ok := f.File("my_file")

		if !ok {
			t.Fatalf("未获取到my_file文件")
		}

		if myFile.Filename != "params_form.go" {
			t.Fatalf("获取到my_file的文件名不正确")
		}

		name, ok := f.String("name")

		if !ok {
			t.Fatalf("未获取到name值")
		}
		if name != "test" {
			t.Fatalf("获取到的name值不正确")
		}
	})

	t.Run("JSON表单", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte("{\"name\":\"test\",\"pwd\":123,\"children\":[1,2]}")))
		r.Header.Set("Content-Type", util.MIMEJSON)
		req := NewRequest(r)

		f := req.FormParams

		name, ok := f.String("name")

		if !ok {
			t.Fatalf("未获取到name值")
		}
		if name != "test" {
			t.Fatalf("获取到的name值不正确")
		}

		pwd, ok := f.Int("pwd")

		if !ok {
			t.Fatalf("未获取到pwd值")
		}
		if pwd != 123 {
			t.Fatalf("获取到的pwd值不正确")
		}

		children, ok := f.Array("children")
		if !ok {
			t.Fatalf("未获取到children值")
		}

		if !reflect.DeepEqual(children, []string{"1", "2"}) {
			t.Fatalf("获取到的children值不正确")
		}
	})
}

func TestFormParams_String(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
			},
			want:  "123",
			want1: true,
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "name",
			},
			want:  "test",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got, got1 := p.String(tt.args.key)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FormParams.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormParams_Int(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
		want1  bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
			},
			want:  123,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got, got1 := p.Int(tt.args.key)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FormParams.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormParams_Uint(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
		want1  bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
			},
			want:  123,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got, got1 := p.Uint(tt.args.key)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FormParams.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormParams_Float(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		want1  bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "price",
			},
			want:  12.34,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got, got1 := p.Float(tt.args.key)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FormParams.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormParams_Bool(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "remember",
			},
			want:  true,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got, got1 := p.Bool(tt.args.key)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FormParams.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormParams_DefaultString(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
		def string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
				def: "qq",
			},
			want: "123",
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd2",
				def: "qq",
			},
			want: "qq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got := p.DefaultString(tt.args.key, tt.args.def)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormParams_DefaultInt(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
		def int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
				def: 456,
			},
			want: 123,
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd2",
				def: 456,
			},
			want: 456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got := p.DefaultInt(tt.args.key, tt.args.def)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormParams_DefaultUint(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
		def uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd",
				def: 456,
			},
			want: 123,
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "pwd2",
				def: 456,
			},
			want: 456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got := p.DefaultUint(tt.args.key, tt.args.def)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormParams_DefaultFloat(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
		def float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "price",
				def: 456.3,
			},
			want: 12.34,
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "price2",
				def: 456.3,
			},
			want: 456.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got := p.DefaultFloat(tt.args.key, tt.args.def)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormParams_DefaultBool(t *testing.T) {
	type fields struct {
		form   map[string]interface{}
		file   map[string][]*multipart.FileHeader
		isJSON bool
	}

	type args struct {
		key string
		def bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "普通表单1",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "remember",
				def: false,
			},
			want: true,
		},
		{
			name: "普通表单2",
			fields: fields{
				form:   testFormData,
				isJSON: false,
			},
			args: args{
				key: "remember2",
				def: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FormParams{
				form:   tt.fields.form,
				file:   tt.fields.file,
				isJSON: tt.fields.isJSON,
			}
			got := p.DefaultBool(tt.args.key, tt.args.def)
			if got != tt.want {
				t.Errorf("FormParams.String() got = %v, want %v", got, tt.want)
			}
		})
	}
}
