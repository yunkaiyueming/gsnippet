package test_ok

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, playground")
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Routes []Route

//var routes Routes

var routes Routes = Routes{
	Route{
		Name:        "GetRoutes",
		Method:      GET,
		Pattern:     "/routes",
		HandlerFunc: GetRoutes,
	},
}

func GetRoutes(res http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(res).Encode([]Routes{}); err != nil {
		panic(err)
	}
}

//func init() {
//	routes = Routes{
//		Route{
//			Name:        "GetRoutes",
//			Method:      GET,
//			Pattern:     "/routes",
//			HandlerFunc: GetRoutes,
//		},
//	}
//}
