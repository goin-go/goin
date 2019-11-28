package request

import "testing"

var p = &URLParams{
	kv: make(map[string]string, 0),
}

func TestUrlParamsNormal(t *testing.T) {

	t.Run("Add", func(t *testing.T) {
		p.kv["name"] = "admin"

		if p.kv["name"] != "admin" {
			t.Error("添加数据失败")
			t.FailNow()
		}
	})
	t.Run("GetString", func(t *testing.T) {
		if s, ok := p.String("name"); !ok || s != "admin" {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("GetInt", func(t *testing.T) {
		p.kv["age"] = "12"
		if age, ok := p.Int("age"); !ok || age != 12 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("DefaultGetInt", func(t *testing.T) {
		if age := p.DefaultInt("age2", 12); age != 12 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("GetUint", func(t *testing.T) {
		p.kv["age"] = "12"
		if age, ok := p.Uint("age"); !ok || age != 12 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("DefaultGetUint", func(t *testing.T) {
		if age := p.DefaultUint("age2", 12); age != 12 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("GetFloat", func(t *testing.T) {
		p.kv["num"] = "12.2"
		if num, ok := p.Float("num"); !ok || num != 12.2 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("DefaultGetFloat", func(t *testing.T) {
		if num := p.DefaultFloat("num2", 12.2); num != 12.2 {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("GetBool", func(t *testing.T) {
		p.kv["isChild"] = "true"
		if age, ok := p.Bool("isChild"); !ok || age != true {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
	t.Run("DefaultGetBool", func(t *testing.T) {
		if age := p.DefaultBool("isChild2", false); age != false {
			t.Error("获取数据失败")
			t.FailNow()
		}
	})
}
