package lwebsocket

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Ws()  {
	router := mux.NewRouter()
	router.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {

	})
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}