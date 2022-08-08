package main

import (
	"fmt"
	"net/http"
)

// Server 是 http server 的顶级抽象
type Server interface {

	// Route 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	// method POST, GET, PUT
	Route(method string, pattern string, handlerFunc func(ctx *Context))

	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用的上
	Name    string
	handler *HandlerBasedOnMap // 基于 map 的路由
}

func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {
	handler := s.handler
	http.Handle("/", handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	// 返回一个实际类型是我实现接口的时候，需要取址
	return &sdkHttpServer{
		Name: name,
	}
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	// 返回一个虚拟的 user id 表示注册成功了
	fmt.Fprintf(ctx.W, "%d", 123)

	// 返回 json 对象
	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
