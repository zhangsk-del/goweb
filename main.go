package main

import (
	"fmt"
	"goweb/gee"
)

func main() {

	engine := gee.NewEngine()

	engine.Get("/hello", func(ctx *gee.Context) {
		fmt.Fprintf(ctx.Writer, "hello golang")
	})

	engine.Run(":8000")

}
