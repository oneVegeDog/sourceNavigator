// Code generated by Kitex v0.7.1. DO NOT EDIT.

package userservice

import (
			"context"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
				baseresp "github.com/oneVegeDog/sourceNavigator/pkg/rpc/baseresp"
				userservice "github.com/oneVegeDog/sourceNavigator/pkg/rpc/userService"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
 }

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userservice.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"register":
			kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"findUserByName":
			kitex.NewMethodInfo(findUserByNameHandler, newUserServiceFindUserByNameArgs, newUserServiceFindUserByNameResult, false),
		"checkedUser":
			kitex.NewMethodInfo(checkedUserHandler, newUserServiceCheckedUserArgs, newUserServiceCheckedUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "userservice",
		"ServiceFilePath": "idl\rpc\user.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName: 	 serviceName,
		HandlerType: 	 handlerType,
		Methods:     	 methods,
		PayloadCodec:  	 kitex.Thrift,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}



func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*userservice.UserServiceRegisterArgs)
	realResult := result.(*userservice.UserServiceRegisterResult)
	success, err := handler.(userservice.UserService).Register(ctx, realArg.Req)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newUserServiceRegisterArgs() interface{} {
	return userservice.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return userservice.NewUserServiceRegisterResult()
}


func findUserByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*userservice.UserServiceFindUserByNameArgs)
	realResult := result.(*userservice.UserServiceFindUserByNameResult)
	success, err := handler.(userservice.UserService).FindUserByName(ctx, realArg.Req)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newUserServiceFindUserByNameArgs() interface{} {
	return userservice.NewUserServiceFindUserByNameArgs()
}

func newUserServiceFindUserByNameResult() interface{} {
	return userservice.NewUserServiceFindUserByNameResult()
}


func checkedUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*userservice.UserServiceCheckedUserArgs)
	realResult := result.(*userservice.UserServiceCheckedUserResult)
	success, err := handler.(userservice.UserService).CheckedUser(ctx, realArg.Req)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newUserServiceCheckedUserArgs() interface{} {
	return userservice.NewUserServiceCheckedUserArgs()
}

func newUserServiceCheckedUserResult() interface{} {
	return userservice.NewUserServiceCheckedUserResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) Register(ctx context.Context , req *userservice.CreateUserReq) (r *baseresp.Baseresp, err error) {
	var _args userservice.UserServiceRegisterArgs
	_args.Req = req
	var _result userservice.UserServiceRegisterResult
	if err = p.c.Call(ctx, "register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindUserByName(ctx context.Context , req *userservice.FindUserReq) (r *userservice.FindUserResp, err error) {
	var _args userservice.UserServiceFindUserByNameArgs
	_args.Req = req
	var _result userservice.UserServiceFindUserByNameResult
	if err = p.c.Call(ctx, "findUserByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckedUser(ctx context.Context , req *userservice.CheckUserReq) (r *userservice.CheckUserResp, err error) {
	var _args userservice.UserServiceCheckedUserArgs
	_args.Req = req
	var _result userservice.UserServiceCheckedUserResult
	if err = p.c.Call(ctx, "checkedUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

