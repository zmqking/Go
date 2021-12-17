package server

import "net/http"

type Server interface {
	Rout(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

//实现接口
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	panic("implement me")
}

func (s *sdkHttpServer) Start(address string) error {
	panic("implement me")
}

type Header map[string][]string
