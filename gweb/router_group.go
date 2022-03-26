package gweb

import (
	"go-web/context"
	"log"
)

type RouterInterface interface {
	Group(prefix string) *RouterGroup
}

type RouterGroup struct {
	prefix string // 分组叠加
	engine *Engine
}

// Group 提供路由分组控制
func (group RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
	}
}

// Get 提供定义Get请求的方法
func (group *RouterGroup) Get(router string, handler context.HandlerFunc) {
	pattern := group.prefix + router
	log.Println("Get", "router:", pattern)
	group.engine.Get(pattern, handler)
}

// Post 提供定义post请求的方法
func (group *RouterGroup) Post(router string, handler context.HandlerFunc) {
	pattern := group.prefix + router
	log.Println("Post", "router:", pattern)
	group.engine.Post(pattern, handler)
}
