package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc http方法
type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
	groups []*RouterGroup // store all groups
}

// NewEngine 初始化引擎
func NewEngine() *Engine {
	engine := &Engine{}
	engine.RouterGroup = NewRouterGroup(engine)
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// NewEngineDefault 初始化一个默认的引擎
func NewEngineDefault() *Engine {
	engine := NewEngine()
	engine.Use(Logger(), Recovery())
	return engine
}

// Run 启动一个httpServer
func (e *Engine) Run(addr string) {
	log.Printf("Listen And Serve%s", addr)
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 	找到则执行对应的方法
	var middlewares []HandlerFunc

	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, r)
	c.handlers = middlewares

	handler := e.router.getRouter(e.router.getKey(c.Method, c.Path))

	if handler != nil {
		c.handlers = append(c.handlers, handler)
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
