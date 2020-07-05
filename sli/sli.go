package sli

import "reflect"

//判断是否是切片类型
func IsSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

//切片逆序
func Reverse(arg interface{}) interface{} {
	slice, ok := IsSlice(arg)
	if !ok {
		return arg
	}
	l := slice.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = slice.Index(l - i - 1).Interface()
	}
	return ret
}
