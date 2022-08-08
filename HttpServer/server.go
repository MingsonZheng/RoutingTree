package main

import "net/http"

// Server 是 http server 的顶级抽象
type Server interface {

	// Route 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	Route(pattern string, handlerFunc http.HandlerFunc)

	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用的上
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	// 返回一个实际类型是我实现接口的时候，需要取址
	return &sdkHttpServer{
		Name: name,
	}
}
