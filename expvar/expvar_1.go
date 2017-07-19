package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

var visits = expvar.NewInt("visits")
var tcpMap = expvar.NewMap("tcp")
var requests, requestsFailed expvar.Int
var start = time.Now()

func init() {
	tcpMap.Set("requests", &requests)
	tcpMap.Set("requests_failed", &requestsFailed)
}

func CalcuateTime() interface{} {
	return time.Since(start).String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	expvar.Publish("runTime", expvar.Func(CalcuateTime))
	http.ListenAndServe(":1818", nil)
}
