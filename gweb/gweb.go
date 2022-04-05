package gweb

import (
	"go-web/context"
	"go-web/middle"
	"go-web/router"
	"log"
	"net/http"
	"strings"
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
	groups []*RouterGroup // store all groups
}

//  New 是获取Engine的构造函数
func New() EngineInterface {
	engine := &Engine{router: router.NewRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//  NewDefault 是获取Engine的默认构造函数
func NewDefault() EngineInterface {
	engine := &Engine{router: router.NewRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.Use(middle.Logger(), middle.Recovery())
	engine.groups = []*RouterGroup{engine.RouterGroup}
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
	// 查找分组中间件 使用前缀来查找
	var middlewares []context.HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(c.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c.Handlers = middlewares
	engine.handler(c)

}

func (engine *Engine) handler(c *context.Context) {
	handler := engine.router.GetRouter(engine.router.CreateKey(c.Method, c.Path))

	if handler != nil {
		c.Handlers = append(c.Handlers, handler)

	} else {
		c.Handlers = append(c.Handlers, func(c *context.Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	// 将排列好的handler最终按顺序执行
	c.Next()
}
