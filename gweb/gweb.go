package gweb

import (
	"go-web/context"
	"go-web/router"
	"log"
	"net/http"
)

// EngineInterface 提供访问Engine的方法
type EngineInterface interface {
	RouterInterface
	Get(router string, handler context.HandlerFunc)
	Post(router string, handler context.HandlerFunc)
	Run(addr string) error
}

// Engine 实现 Handler ServerHttp
type Engine struct {
	*RouterGroup
	router *router.Router
}

//  New 是获取Engine的构造函数
func New() EngineInterface {
	engine := &Engine{router: router.NewRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	return engine
}

// Run 提供启动http服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// Get 提供定义Get请求的方法
func (engine *Engine) Get(router string, handler context.HandlerFunc) {
	engine.router.AddRouter("GET", router, handler)
}

// Post 提供定义post请求的方法
func (engine *Engine) Post(router string, handler context.HandlerFunc) {
	engine.router.AddRouter("POST", router, handler)
}

// ServeHTTP 核心处理逻辑
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("router:", req.URL, "method:", req.Method)
	c := context.NewContext(resp, req)
	handler := engine.router.GetRouter(engine.router.CreateKey(c.Method, c.Path))
	if handler != nil {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
