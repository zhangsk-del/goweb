package gee

import (
	"net/http"
	"testing"
)

func TestRouterGroup(t *testing.T) {

	r := NewEngine()
	r.Get("/index", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/", func(c *Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.Get("/hello", func(c *Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
		v2.Post("/login", func(c *Context) {
			c.JSON(http.StatusOK, H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":8000")
}
