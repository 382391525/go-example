package lreflect

import (
	"fmt"
	"reflect"
)

func TypeOf() {
	var i interface{}
	i = 1
	typeOf := reflect.TypeOf(i)
	fmt.Println(typeOf.Name())
}

func ValueOf() {
	var i interface{}
	i = 1
	valueOf := reflect.ValueOf(i)
	fmt.Println(valueOf.Type())
}

func ReflectText(i interface{}) {

	reType := reflect.TypeOf(i)
	fmt.Println("reflect.Type=", reType)
	reVal := reflect.ValueOf(i)
	reValKind := reVal.Kind()
	fmt.Println("reflect.ValueOf=", reVal)
	fmt.Println("reflect.Kind=", reValKind)
	fmt.Printf("reflect.Value= %v, type= %T\n", reVal, reVal)
	iVal := reVal.Interface()
	num := iVal.(int)
	fmt.Printf("num = %v, type = %T\n", num, num)
}

type Student struct {
	Id   int
	Name string
	Age  int
}

func ReflectStruct(i interface{}) {
	// 获取变量的 reflect.Type
	reType := reflect.TypeOf(i)
	fmt.Println("reflect.Type=", reType)

	// 获取变量的 reflect.Value
	reVal := reflect.ValueOf(i)
	fmt.Println("reflect.Value=", reVal)

	// 将 reVal 转成 interface
	iVal := reVal.Interface()
	fmt.Printf("iVal= %v, type= %T\n", iVal, iVal)
	// iVal.Name 会报错Unresolved reference 'Name'
	// fmt.Printf("iVal= %v, type= %T, name= %v\n", iVal, iVal, iVal.Name)

	// 将 interface 通过类型断言 转回成 Student
	// stu:= iVal.(Student)
	if stu, ok := iVal.(Student); ok {
		fmt.Printf("stu= %v, type= %T, name=%v\n", stu, stu, stu.Name)
	}
}
