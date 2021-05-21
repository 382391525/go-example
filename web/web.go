package web

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "http 服务器")
	})
	http.ListenAndServe(":8080", nil)
}
