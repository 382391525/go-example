package sli

import (
	"fmt"
	"reflect"
)

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

func Add() {

	s1 := make([]int, 4, 10)
	fmt.Printf("1  v: %v,p:%p\n", s1, s1)
	foo(&s1)
	fmt.Printf("3 v: %v,p:%p\n", s1, s1)
	s3 := s1[0:cap(s1)]
	fmt.Printf("3 v: %v,p:%p\n", s3, s3)
}

func foo(sp *[]int) {
	s := append(*sp, 10)
	fmt.Printf("2 v: %v,p:%p\n", s, s)
}
