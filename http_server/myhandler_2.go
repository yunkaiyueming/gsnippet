package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func sayHello3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 3.")
}

func sayBye3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye, this is version 3.")
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("统一处理<br>"))
	if h, ok := mux[req.URL.String()]; ok {

		h(res, req)
		return
	}
	io.WriteString(res, "My server: "+req.URL.String())
}

func main() {
	server := http.Server{
		Addr:        ":12345",     // TCP address to listen on, ":http" if empty
		Handler:     &myHandler{}, // handler to invoke, http.DefaultServeMux if nil
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello3
	mux["/bye"] = sayBye3

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
