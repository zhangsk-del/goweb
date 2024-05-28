package gee

import (
	"net/http"
)

// HandlerFunc http方法
type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
}

// NewEngine 初始化引擎
func NewEngine() *Engine {
	engine := &Engine{}
	engine.RouterGroup = NewRouterGroup(engine)
	return engine
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
