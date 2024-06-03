package main

import (
	"fmt"
	"goweb/gee"
)

func main() {

	engine := gee.NewEngine()
	engine.Use(gee.Logger(), gee.Recovery())

	v1 := engine.Group("/v1")
	{
		v1.Get("/hello", func(ctx *gee.Context) {
			name := ctx.Query("name")
			path := ctx.Path
			req := ctx.Req
			fmt.Println(path)
			fmt.Println(req)
			ctx.String(200, "hello golang v1 test1,query %s", name)
		})

		v1.Get("/test", func(ctx *gee.Context) {
			ctx.JSON(200, "hello golang v1 test2")
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.Get("/hello", func(ctx *gee.Context) {
			fmt.Fprintf(ctx.Writer, "hello golang v2 test1")

		})
		v2.Post("/test", func(ctx *gee.Context) {
			ctx.Data(200, []byte("hello golang v2 test2"))
		})
	}

	engine.Run(":8000")

}
