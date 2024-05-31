package gee

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
