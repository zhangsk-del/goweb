package gee

import "testing"

func TestHttpContext(t *testing.T) {
	engine := NewEngine()

	engine.Get("/hello", func(ctx *Context) {

		ctx.String(200, "hello golang")
	})

	engine.Run(":8000")
}
