package gee

import (
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

// Run 启动一个httpServer
func (e *Engine) Run(addr string) {
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

	ctx := newContext(w, r)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}
