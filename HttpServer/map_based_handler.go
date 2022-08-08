package main

import (
	"net/http"
)

type Handler interface {
	http.Handler
	Route(method string, pattern string, handlerFunc func(ctx *Context)) // server 可以把 Route 委托给这边的 Handler
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

// HandlerBasedOnMap 基于 map 的路由
type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	// 判定路由是否已经注册
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
