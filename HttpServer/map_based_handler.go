package main

import (
	"net/http"
	"sync"
)

type handlerFunc func(ctx *Context)

type Routable interface {
	// Route 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	// method POST, GET, PUT
	Route(method string, pattern string, handlerFunc handlerFunc) // server 可以把 Route 委托给这边的 Handler
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

// HandlerBasedOnMap 基于 map 的路由
type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers sync.Map
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	request := c.R
	key := h.key(request.Method, request.URL.Path)
	// 判定路由是否已经注册
	handler, ok := h.handlers.Load(key)
	if !ok {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	handler.(func(c *Context))(c)
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
