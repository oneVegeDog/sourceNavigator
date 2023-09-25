// Code generated by Kitex v0.7.1. DO NOT EDIT.

package sourcenavigatorservice

import (
	server "github.com/cloudwego/kitex/server"
	source_rpc "github.com/oneVegeDog/sourceNavigator/pkg/source_rpc"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler source_rpc.SourceNavigatorService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
