package ldefer

import "fmt"

func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func Test() {
	println(f())
	println(f1())
	println(f2())
	println(f3())
}

type Person struct {
	age int
}

func Test2() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

func defer3() int {
	s := 1
	defer func() {
		s++ //2
		fmt.Println(s)
	}()
	defer func(s int) {
		s++ //2
		fmt.Println(s)
	}(s)
	return s
}

func Test3()  {
	fmt.Println(defer3()) //1
}