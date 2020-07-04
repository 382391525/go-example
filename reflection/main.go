package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Hello() {
	fmt.Println("hello world ")
}

func main() {
	m := Manager{
		User: User{
			Id:   2,
			Name: "fzf2",
			Age:  123,
		},
		title: "123",
	}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n",t)
	fmt.Printf("%#v\n",t.Field(0))
	fmt.Printf("%#v\n",t.Field(1))
}

type Manager struct {
	User
	title string
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s : %v \n", m.Name, m.Type)

	}
}
