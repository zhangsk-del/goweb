package gee

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func onlyForV2() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Test(t *testing.T) {
	r := NewEngine()
	r.Use(Logger()) // global midlleware
	r.Get("/", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.Get("/hello", func(c *Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello middle")
		})
	}

	r.Run(":8000")
}
