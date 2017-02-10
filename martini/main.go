package main

import (
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		// return 200, "hello"
		res.WriteHeader(200)
		res.Write([]byte("Ss"))
	})

	m.Get("/user/:name", func(params martini.Params) string {
		return "user " + params["name"]
	})

	m.Run()
}
