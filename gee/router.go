package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc // 请求url和方法的映射
}

type RouterGroup struct {
	prefix string       // 路由前缀
	parent *RouterGroup // 父级
	engine *Engine
	router *router
}

func NewRouterGroup(engine *Engine) *RouterGroup {
	return &RouterGroup{
		engine: engine,
		router: newRouter(),
	}
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
		router: group.router,
	}
	return newGroup
}

// Post 提供对外的Post方法
func (group *RouterGroup) Post(addr string, handler HandlerFunc) {
	pattern := group.prefix + addr
	group.router.addRouter("POST", pattern, handler)
}

// Get 提供对外的Get方法
func (group *RouterGroup) Get(addr string, handler HandlerFunc) {
	pattern := group.prefix + addr
	group.router.addRouter("GET", pattern, handler)
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
