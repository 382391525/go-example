package lpprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

var data []string

func Test() {
	go func() {
		for {
			l := Add("fzf")
			fmt.Println("len:", l)
			time.Sleep(10 * time.Second)
		}
	}()
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		return
	}
}

func Add(str string) int {
	d := []byte(str)
	data = append(data, string(d))
	return len(data)
}

func Test2() {
	http.HandleFunc("/lookup/heap", func(w http.ResponseWriter, r *http.Request) {
		lookup := pprof.Lookup("heap")
		err := lookup.WriteTo(os.Stdout, 2)
		if err != nil {
			return 
		}
	})
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
