package lfunc

import "fmt"

type Option func(*option)

type option struct {
	name string
	sex  int
	age  int
}

func WithSex(sex int) Option {
	return func(o *option) {
		o.sex = sex
	}
}

func WithAge(age int) Option {
	return func(o *option) {
		o.age = age
	}
}

func NewMan(name string, opts ...Option) (*option, error) {
	opt := new(option)
	opt.name = name
	for _, f := range opts {
		f(opt)
	}
	return opt, nil
}

func Test() {
	man, err := NewMan("张三", WithSex(1), WithAge(20))
	if err != nil {
		return
	}
	fmt.Println(man)
}
