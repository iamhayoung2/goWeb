package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json : "first_name`
	LastName  string    `json : "last_name"`
	Email     string    `json : "email"`
	CreatedAt time.Time `json : create_at`
}
type fooHandler struct{} // 인스턴스를 만들고

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello Foo!")
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // json 형태로 parsing
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad Reqauest : ", err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
} // ServeHttp라는 Interface를 구현

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {

	mux := http.NewServeMux() // mux 라우터임.

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}) // Handfunc - 핸드러를 등록하는 것. 어떤 경로에 해당하는 reqeust가 들어왔을때 어떤 일을 할 것인지 핸들러 등록

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{}) // 핸드럴를 인스턴스 형태로 등록할때 Handle 함수 사용  // 인스턴스의 인터페이스를 핸들러에 등록

	http.ListenAndServe(":3040", mux) // 위에서 라우터를 만들고, 마지막 listenandserve에 mux = 라우터를 등록함
}

// 라우터 - mux? : 경로에 따라 다르게 분배함
