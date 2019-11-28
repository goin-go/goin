package request

var testFormData = map[string]interface{}{
	"name":     []string{"test"},
	"pwd":      []string{"123"},
	"remember": []string{"true"},
	"price":    []string{"12.34"},
}

var testQueryData = map[string][]string{
	"name":     []string{"test"},
	"pwd":      []string{"123"},
	"remember": []string{"true"},
	"price":    []string{"12.34"},
	"children": []string{"1", "2", "3"},
}
