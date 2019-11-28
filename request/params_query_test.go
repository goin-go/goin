package request

import (
	"net/url"
	"reflect"
	"testing"
)

func TestQueryParams_Array(t *testing.T) {
	type fields struct {
		query url.Values
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
		want1  bool
	}{
		{
			name: "query_array",
			fields: fields{
				query: testQueryData,
			},
			args: args{
				key: "children",
			},
			want:  []string{"1", "2", "3"},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.Array(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Array() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryParams_Bool(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.Bool(tt.args.key)
			if got != tt.want {
				t.Errorf("Bool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Bool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryParams_DefaultBool(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			if got := p.DefaultBool(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParams_DefaultFloat(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
			},
			args: args{
				key: "price2",
				def: 10.1,
			},
			want: 10.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &QueryParams{
				query: tt.fields.query,
			}
			if got := p.DefaultFloat(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("DefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParams_DefaultInt(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			if got := p.DefaultInt(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParams_DefaultString(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
			},
			args: args{
				key: "name2",
				def: "admin",
			},
			want: "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &QueryParams{
				query: tt.fields.query,
			}
			if got := p.DefaultString(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParams_DefaultUint(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			if got := p.DefaultUint(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("DefaultUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParams_Float(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.Float(tt.args.key)
			if got != tt.want {
				t.Errorf("Float() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Float() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryParams_Int(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.Int(tt.args.key)
			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Int() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryParams_String(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.String(tt.args.key)
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryParams_Uint(t *testing.T) {
	type fields struct {
		query url.Values
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
			name: "query_array",
			fields: fields{
				query: testQueryData,
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
			p := &QueryParams{
				query: tt.fields.query,
			}
			got, got1 := p.Uint(tt.args.key)
			if got != tt.want {
				t.Errorf("Uint() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Uint() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
