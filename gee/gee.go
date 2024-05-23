package gee

import (
	"net/http"
)

// HandlerFunc http方法
type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

// NewEngine 初始化引擎
func NewEngine() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// Post 提供对外的Post方法
func (e *Engine) Post(addr string, handler HandlerFunc) {
	e.router.addRouter("POST", addr, handler)
}

// Get 提供对外的Get方法
func (e *Engine) Get(addr string, handler HandlerFunc) {
	e.router.addRouter("GET", addr, handler)
}

// Run 启动一个httpServer
func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 	找到则执行对应的方法
	ctx := newContext(w, r)
	e.router.handle(ctx)
}
