package request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goin-go/goin/util"
)

func TestRequest_GET(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/query?name=test&pwd=123&children=1&children=2&c", nil)
	req := NewRequest(r)

	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "alias_default_get",
			args: args{
				key: "name",
			},
			want:  "test",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := req.GET(tt.args.key)
			if got != tt.want {
				t.Errorf("Request.GET() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Request.GET() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRequest_DefaultGET(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/query?name=test&pwd=123&children=1&children=2&c", nil)
	req := NewRequest(r)

	type args struct {
		key string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alias_get",
			args: args{
				key: "name2",
				def: "admin",
			},
			want: "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := req.DefaultGET(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("Request.DefaultGET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_POST(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte("name=test&pwd=123&children=1&children=2&c")))
	r.Header.Set("Content-Type", util.MIMEPOSTForm)
	req := NewRequest(r)

	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "alias_post",
			args: args{
				key: "name",
			},
			want:  "test",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := req.POST(tt.args.key)
			if got != tt.want {
				t.Errorf("Request.POST() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Request.POST() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRequest_DefaultPOST(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer([]byte("name=test&pwd=123&children=1&children=2&c")))
	r.Header.Set("Content-Type", util.MIMEPOSTForm)
	req := NewRequest(r)

	type args struct {
		key string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alias_default_post",
			args: args{
				key: "name2",
				def: "admin",
			},
			want: "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := req.DefaultPOST(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("Request.DefaultPOST() = %v, want %v", got, tt.want)
			}
		})
	}
}
