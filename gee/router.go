package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc // 请求url和方法的映射
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRouter
func (r *router) addRouter(method, addr string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, addr)
	r.handlers[method+"-"+addr] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
