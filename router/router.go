package router

import (
	"go-web/context"
)

type Router struct {
	// 存储 GET-/router -- HandlerFunc 映射关系
	router map[string]context.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		router: make(map[string]context.HandlerFunc),
	}
}

// AddRoute 提供添加映射关系
func (r *Router) AddRouter(method string, router string, handler context.HandlerFunc) {
	key := r.CreateKey(method, router)
	r.router[key] = handler
}

// GetRouter  提供获取映射关系
func (r *Router) GetRouter(key string) context.HandlerFunc {
	if handler, ok := r.router[key]; ok {
		return handler
	}
	return nil
}

// CreateKey 提供创建路由key
func (r *Router) CreateKey(method string, router string) string {
	return method + "-" + router
}
