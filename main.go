package main

import (
	"fmt"
	"goweb/gee"
	"net/http"
)

func main() {

	engine := gee.NewEngine()

	engine.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello golang")
	})

	engine.Run(":8000")

}
