package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //handler 등록 - 어떤 request가 들어왔을때 어떤 것을 할 것인지. - index 페이지 경로
		fmt.Fprintf(w, "Hello World") // function에 있는 것을 할 것이다. Fprint : Writer에 프린트 하라
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "heeloba")
	})

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil) // 웹서버 구동. reqeust를 기다리는 상태.
}
