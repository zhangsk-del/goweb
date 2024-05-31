package gee

import (
	"log"
)

type router struct {
	handlers map[string]HandlerFunc // 请求url和方法的映射
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// GetRouter  提供获取映射关系
func (r *router) getRouter(key string) HandlerFunc {
	if handler, ok := r.handlers[key]; ok {
		return handler
	}
	return nil
}

// addRouter
func (r *router) addRouter(method, addr string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, addr)
	r.handlers[r.getKey(method, addr)] = handler
}

// getKey 提供创建路由key
func (r *router) getKey(method string, router string) string {
	return method + "-" + router
}
