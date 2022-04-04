package gweb

import (
	"go-web/context"
	"log"
)

type RouterInterface interface {
	Group(prefix string) *RouterGroup
	Use(middlewares ...context.HandlerFunc)
}

type RouterGroup struct {
	prefix      string // 分组叠加
	engine      *Engine
	middlewares []context.HandlerFunc // support middleware
}

// Group 提供路由分组控制
func (group RouterGroup) Group(prefix string) *RouterGroup {

	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
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

// Use 提供添加中间件方法
func (group *RouterGroup) Use(middlewares ...context.HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
