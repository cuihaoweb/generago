package util

import (
	"reflect"
)

// CopyStruct 将a结构体赋值非b结构体， 前提是两者必须存在相同字段
func CopyStruct(a interface{}, b interface{}) {
	verifyCopyMap(b)
	var t1 = reflect.TypeOf(a)
	var v1 = reflect.ValueOf(a)
	var v2 = reflect.ValueOf(b).Elem()

	for k := 0; k < t1.NumField(); k++ {
		name := t1.Field(k).Name
		value := v1.FieldByName(name)
		if val2 := v2.FieldByName(name); val2.IsValid() == true {
			val2.Set(value)
		}
	}
}
func verifyCopyMap(a interface{}) {
	var t = reflect.TypeOf(a)
	t1 := t.Kind().String()

	if t1 != "ptr" {
		panic("第二个参数必须为地址")
	}
}
