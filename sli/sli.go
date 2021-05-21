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
	fmt.Printf("1 v: %v,p:%p\n", s1, s1)
	foo(&s1)
	fmt.Printf("3 v: %v,p:%p\n", s1, s1)
	s3 := s1[0:cap(s1)]
	fmt.Printf("4 v: %v,p:%p\n", s3, s3)
}

func foo(sp *[]int) {
	s := append(*sp, 10)
	fmt.Printf("2 v: %v,p:%p\n", s, s)
}

func Test() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := s1[2:5]
	s3 := s1[2:4]
	fmt.Println(s1, s2)

	fmt.Printf("t: %T p: %p,len: %d,cap: %d\n", s1, s1, len(s1), cap(s1)) //
	fmt.Printf("t: %T p: %p,len: %d,cap: %d\n", s2, s2, len(s2), cap(s2)) //
	fmt.Printf("t: %T p: %p,len: %d,cap: %d\n", s3, s3, len(s3), cap(s3)) //

}

func Test2() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("v: %v, p: %p,p2: %p,len: %d,cap: %d\n", arr, arr, &arr, len(arr), cap(arr))
	for i, _ := range arr {
		fmt.Printf("i: %d p:%p\n ", i, &arr[i])
	}
	// 切片1
	s1 := arr[2:5] // 从下标2开始取，取三个元素
	s1[1] = 111
	fmt.Printf("v: %v,p: %p,p2: %p,len: %d,cap: %d\n", s1, s1, &s1, len(s1), cap(s1))
	// 切片2
	s2 := s1[2:7] // 从下标2开始取，取三个元素
	s2[1] = 222
	fmt.Printf("v: %v, p: %p,p2: %p,len: %d,cap: %d\n", s2, s2, &s2, len(s2), cap(s2))

	fmt.Printf("v: %v, p: %p,p2: %p,len: %d,cap: %d\n", arr, arr, &arr, len(arr), cap(arr))
	//切片3 扩容
	s3 := append(arr, 10)
	fmt.Printf("v: %v, p: %p,p2: %p,len: %d,cap: %d\n", s3, s3, &s3, len(s3), cap(s3))
	for i, _ := range s3 {
		fmt.Printf("i: %d p:%p\n ", i, &s3[i])
	}
}
