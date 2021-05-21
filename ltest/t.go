package ltest

import "fmt"

type MS struct {
	Name string
	Cgs  int
}

func Test() {
	hashRun := make(map[string]*MS)
	mss := []MS{
		MS{Name: "m", Cgs: 100},
		MS{Name: "m", Cgs: 1},
	}
	for _, ms := range mss {
		if h, ok := hashRun[ms.Name]; ok {
			h.Cgs += ms.Cgs
		} else {
			hashRun[ms.Name] = &ms
		}
	}
	for _, ms := range hashRun {
		fmt.Println(ms.Cgs)
	}
}

func Test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Printf("v: %v,p:%p\n", val, &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func Test3() {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
}
