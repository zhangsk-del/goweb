package gweb

import (
	"fmt"
	"go-web/context"
	"net/http"
)

// EngineInterface 提供访问Engine的方法
type EngineInterface interface {
	Get(router string, handler HandlerFunc)
	Post(router string, handler HandlerFunc)
	Run(addr string) error
}

// HandlerFunc 提供处理函数模板
type HandlerFunc func(*context.Context)

// Engine 实现 Handler ServerHttp
type Engine struct {
	// 存储 GET-/router -- HandlerFunc 映射关系
	router map[string]HandlerFunc
}

//  New 是获取Engine的构造函数
func New() EngineInterface {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// Get 提供定义Get请求的方法
func (engine *Engine) Get(router string, handler HandlerFunc) {
	engine.addRoute("GET", router, handler)
}

// Post 提供定义post请求的方法
func (engine *Engine) Post(router string, handler HandlerFunc) {
	engine.addRoute("POST", router, handler)
}

func (engine *Engine) addRoute(method string, router string, handler HandlerFunc) {
	key := method + "-" + router
	engine.router[key] = handler
}

// Run 提供启动http服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 核心处理逻辑
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		c := context.NewContext(resp, req)
		handler(c)
	} else {
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}
