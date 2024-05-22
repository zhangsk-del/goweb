package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc http方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc // 请求url和方法的映射
}

// NewEngine 初始化引擎
func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// addRouter
func (e *Engine) addRouter(method, addr string, handler HandlerFunc) {
	e.router[method+"-"+addr] = handler
}

// Post 提供对外的Post方法
func (e *Engine) Post(addr string, handler HandlerFunc) {
	e.addRouter("POST", addr, handler)
}

// Get 提供对外的Get方法
func (e *Engine) Get(addr string, handler HandlerFunc) {
	e.addRouter("GET", addr, handler)
}

// Run 启动一个httpServer
func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	key := r.Method + "-" + r.URL.Path

	handlerFunc, ok := e.router[key]
	if ok {
		// 	找到则执行对应的方法
		handlerFunc(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
