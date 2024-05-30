package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc // 请求url和方法的映射
}

type RouterGroup struct {
	prefix      string       // 路由前缀
	parent      *RouterGroup // 父级
	engine      *Engine
	router      *router
	middlewares []HandlerFunc // support middleware
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
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Post 提供对外的Post方法
func (group *RouterGroup) Post(addr string, handler HandlerFunc) {
	pattern := group.prefix + addr
	group.router.addRouter("POST", pattern, handler)
}

// Use 提供添加中间件的方法
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// Get 提供对外的Get方法
func (group *RouterGroup) Get(addr string, handler HandlerFunc) {
	pattern := group.prefix + addr
	group.router.addRouter("GET", pattern, handler)
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// GetRouter  提供获取映射关系
func (r *router) GetRouter(key string) HandlerFunc {
	if handler, ok := r.handlers[key]; ok {
		return handler
	}
	return nil
}

// addRouter
func (r *router) addRouter(method, addr string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, addr)
	r.handlers[method+"-"+addr] = handler
}

// CreateKey 提供创建路由key
func (r *router) CreateKey(method string, router string) string {
	return method + "-" + router
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	handler := r.GetRouter(key)

	if handler != nil {
		c.handlers = append(c.handlers, handler)
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
