package lmap

import "fmt"

func Test() {
	m1 := make(map[string]string, 10)
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "bb"
	fmt.Printf("1  v: %v,p:%p\n", m1, m1)
	EditMap(m1)
	fmt.Printf("3  v: %v,p:%p\n", m1, m1)
}

func EditMap(m map[string]string) {
	m["edit"] = "test"
	fmt.Printf("2  v: %v,p:%p\n", m, m)
}
